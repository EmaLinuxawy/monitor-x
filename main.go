package main

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
)

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

func printCPUstats(cpuStats []float64) {
	fmt.Println("-----------")
	fmt.Println("Consumtion Per CPU")
	fmt.Println("-----------")
	for i, percent := range cpuStats {
		fmt.Printf("CPU %d: %f%%\n", i, percent)
	}
	fmt.Println("-----------")
}

func getMemoryUsed(ctx context.Context) (*mem.VirtualMemoryStat, error) {
	percentages, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return nil, err
	}
	return percentages, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cpuStats, err := getCPUPercentages(ctx)
	if err != nil {
		fmt.Println("Error fetching CPU Statistics: ", err)
		return
	}
	printCPUstats(cpuStats)

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
}
