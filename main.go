package main

import (
	"context"
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
)

type ProcessInfo struct {
	Proc *process.Process
	CPU  float64
}

type NetworkStats struct {
	InterfaceName string
	BytesSent     uint64
	BytesRecv     uint64
	PacketsSent   uint64
	PacketsRecv   uint64
}

func getLoadAverage(ctx context.Context) (*load.AvgStat, error) {
	avg, err := load.AvgWithContext(ctx)
	if err != nil {
		return nil, err
	}
	return avg, nil
}

func getTotalCPUUsage(ctx context.Context) (float64, error) {
	percentages, err := cpu.PercentWithContext(ctx, 0, false)
	if err != nil {
		return 0, err
	}
	if len(percentages) > 0 {
		return percentages[0], nil
	}
	return 0, fmt.Errorf("could not get CPU usage")
}

func getCPUPercentages(ctx context.Context) ([]float64, error) {
	cpuStats, err := cpu.PercentWithContext(ctx, 0, true)
	if err != nil {
		return nil, err
	}
	return cpuStats, nil
}

func getMemoryUsed(ctx context.Context) (*mem.VirtualMemoryStat, error) {
	percentages, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return nil, err
	}
	return percentages, nil
}

func getDiskUsage(ctx context.Context) (*disk.UsageStat, error) {
	usage, err := disk.UsageWithContext(ctx, "/")
	if err != nil {
		return nil, err
	}
	return usage, nil
}

func getTopProcesses(ctx context.Context) ([]ProcessInfo, error) {
	proces, err := process.Processes()
	if err != nil {
		return nil, err
	}

	var topProcesses []ProcessInfo

	for _, p := range proces {
		cpuPercent, _ := p.CPUPercent()
		topProcesses = append(topProcesses, ProcessInfo{
			Proc: p,
			CPU:  cpuPercent,
		})
	}

	sort.Slice(topProcesses, func(i, j int) bool {
		return topProcesses[i].CPU > topProcesses[j].CPU
	})

	if len(topProcesses) > 10 {
		topProcesses = topProcesses[:10]
	}
	return topProcesses, nil
}

func getNetworkStatistics(ctx context.Context) ([]NetworkStats, error) {
	ioCounters, err := net.IOCounters(true)
	if err != nil {
		return nil, err
	}

	var stats []NetworkStats
	for _, counter := range ioCounters {
		if counter.BytesSent > 0 && counter.BytesRecv > 0 && counter.PacketsSent > 0 && counter.PacketsRecv > 0 {
			stats = append(stats, NetworkStats{
				InterfaceName: counter.Name,
				BytesSent:     counter.BytesSent,
				BytesRecv:     counter.BytesRecv,
				PacketsSent:   counter.PacketsSent,
				PacketsRecv:   counter.PacketsRecv,
			})
		}
	}
	return stats, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cpuStats, err := getCPUPercentages(ctx)
	if err != nil {
		fmt.Println("Error fetching CPU Statistics: ", err)
		return
	}
	cpuWriter := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(cpuWriter, "CPU Core\tPercentage")
	for i, percent := range cpuStats {
		fmt.Fprintf(cpuWriter, " %d:\t%.2f%%\n", i, percent)
	}

	cpuWriter.Flush()

	totalCPUUsage, err := getTotalCPUUsage(ctx)
	if err != nil {
		fmt.Println("Error getting total CPU usage:", err)
		return
	}
	fmt.Printf("Total CPU usage: %f%%\n", totalCPUUsage)
	fmt.Println("-----------")

	loadAvg, err := getLoadAverage(ctx)
	if err != nil {
		fmt.Println("Error getting load average:", err)
		return
	}
	fmt.Printf("Load average: 1 min: %.2f, 5 min: %.2f, 15 min: %.2f\n", loadAvg.Load1, loadAvg.Load5, loadAvg.Load15)

	fmt.Println("-----------")
	memoryPercentages, err := getMemoryUsed(ctx)
	if err != nil {
		fmt.Println("Error getting memory statistics:", err)
		return
	}
	fmt.Printf("Memory Usage: %.2f%%\n", memoryPercentages.UsedPercent)
	fmt.Println("-----------")

	diskUsage, err := getDiskUsage(ctx)
	if err != nil {
		fmt.Println("Error getting disk usage:", err)
		return
	}
	diskWriter := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(diskWriter, "DiskSpace\tFree Space\tTotal Usage\t")
	fmt.Fprintf(diskWriter, "%.2f GB\t%.2f GB\t%.2f GB\t\n",
		float64(diskUsage.Total)/float64(1<<30),
		float64(diskUsage.Free)/float64(1<<30),
		float64(diskUsage.Used)/float64(1<<30))
	diskWriter.Flush()

	//fmt.Printf("Disk Usage: \nTotal=%.2f GB \nFree=%.2f GB \nUsed=%.2f GB \nInodesFree=%d \nInodesUsed=%d\n", float64(diskUsage.Total)/float64(1<<30), float64(diskUsage.Free)/float64(1<<30), float64(diskUsage.Used)/float64(1<<30), diskUsage.InodesFree, diskUsage.InodesUsed)

	topProcesses, err := getTopProcesses(ctx)
	if err != nil {
		fmt.Println("Error getting top processes:", err)
		return
	}
	fmt.Println("-----------")
	fmt.Println("Top 10 Processes by CPU usage:")
	topCPUW := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintln(topCPUW, "PID\tCPU(%)\tMemory(MB)\tName")

	for _, p := range topProcesses {
		memInfo, err := p.Proc.MemoryInfo()
		if err != nil {
			fmt.Println("Error getting memory info:", err)
			continue
		}
		name, err := p.Proc.Name()
		if err != nil {
			fmt.Println("Error getting process name:", err)
			continue
		}

		// Use Fprintf to write formatted output to the tab writer
		fmt.Fprintf(topCPUW, "%d\t%.2f\t%.2f\t%s\n",
			p.Proc.Pid, p.CPU, float32(memInfo.RSS)/(1<<20), name)
	}

	topCPUW.Flush()

	fmt.Println("-----------")
	fmt.Println("Net Info")

	netStats, err := getNetworkStatistics(ctx)
	if err != nil {
		fmt.Println("Error getting network stats:", err)
		return
	}
	fmt.Println("-----------")
	netWriter := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(netWriter, "Interface Name\tBytes Sent (Gbits)\tBytes Received (Gbits)\tPackets Sent\tPackets Received\t")
	for _, netStat := range netStats {
		fmt.Fprintf(netWriter, "%s\t%.2f\t%.2f\t%d\t%d\t\n",
			netStat.InterfaceName,
			float64(netStat.BytesSent)*8/(1<<30),
			float64(netStat.BytesRecv)*8/(1<<30),
			netStat.PacketsSent,
			netStat.PacketsRecv)
	}
	netWriter.Flush()
}
