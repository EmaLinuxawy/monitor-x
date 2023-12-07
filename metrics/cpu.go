package metrics

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/process"
)

type ProcessInfo struct {
	PID    int
	Name   string
	CPU    float64
	Memory float64
}

func GetTotalCPUUsage(ctx context.Context) (float64, error) {
	percentages, err := cpu.PercentWithContext(ctx, 0, false)
	if err != nil {
		return 0, err
	}
	if len(percentages) > 0 {
		return percentages[0], nil
	}
	return 0, fmt.Errorf("could not get CPU usage")
}

func GetTopProcesses(ctx context.Context) ([]ProcessInfo, error) {
	proces, err := process.Processes()
	if err != nil {
		return nil, err
	}

	var topProcesses []ProcessInfo

	for _, p := range proces {
		name, err := p.Name()
		if err != nil {
			continue
		}

		memInfo, err := p.MemoryInfo()
		if err != nil {
			continue
		}
		cpuPercent, err := p.CPUPercent()
		if err != nil {
			continue
		}
		topProcesses = append(topProcesses, ProcessInfo{
			PID:    int(p.Pid),
			Name:   name,
			CPU:    cpuPercent,
			Memory: float64(memInfo.RSS / (1 << 20)),
		})
	}

	sort.Slice(topProcesses, func(i, j int) bool {
		return topProcesses[i].CPU > topProcesses[j].CPU
	})

	if len(topProcesses) > 20 {
		topProcesses = topProcesses[:20]
	}
	return topProcesses, nil
}

func FormatCPUUsage(usage float64) string {
	barLength := 10
	filledLength := int(usage / 100 * float64(barLength))
	bar := strings.Repeat("|", filledLength) + strings.Repeat(" ", barLength-filledLength)

	return bar
}
