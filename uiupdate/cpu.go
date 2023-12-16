package uiupdate

import (
	"context"
	"fmt"

	"github.com/emaLinuxawy/monitor-x/metrics"
	"github.com/emaLinuxawy/monitor-x/view"
	ui "github.com/gizak/termui/v3"
	"github.com/shirou/gopsutil/v3/cpu"
)

func UpdateCPUData(ctx context.Context, v *view.View) {
	cpuStats, err := cpu.PercentWithContext(ctx, 0, true)
	if err != nil {
		fmt.Println("Error getting CPU stats:", err)
		return
	}
	for i, cpu := range cpuStats {
		if i < len(v.CpuChart) {
			bar := metrics.FormatCPUUsage(cpu)
			v.CpuChart[i].Text = fmt.Sprintf("%s %.2f%%\n", bar, cpu)

			if cpu > 80 {
				v.CpuChart[i].TextStyle.Fg = ui.ColorRed
			} else {
				v.CpuChart[i].TextStyle.Fg = ui.ColorGreen
			}
		}
	}
}
func UpdateTotalCPUUsage(ctx context.Context, v *view.View) {
	cpuUsage, err := metrics.GetTotalCPUUsage(ctx)
	if err != nil {
		fmt.Println("Error getting total CPU usage:", err)
		return
	}
	v.CPUUsage.Rows = []string{fmt.Sprintf("%.3f%%\n", cpuUsage)}
}

func UpdateTopProcesses(ctx context.Context, v *view.View) {
	processes, err := metrics.GetTopProcesses(ctx)
	if err != nil {
		fmt.Println("Error fetching top processes:", err)
		return
	}

	var rows []string
	header := fmt.Sprintf("%-7s %-35s %-15s %-15s", "PID", "APP", "CPU %", "MEM (MB)")
	rows = append(rows, header)
	for _, p := range processes {
		row := fmt.Sprintf("%-7d %-35s %-15.2f %-15.2f", p.PID, p.Name, p.CPU, p.Memory)
		rows = append(rows, row)
	}
	v.ProcessList.Rows = rows
}
