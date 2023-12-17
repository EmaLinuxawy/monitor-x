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

			if cpu > 75 {
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

	columns := []Column{
		{"PID", 0},
		{"APP", 0},
		{"CPU %", 0},
		{"MEM (MB)", 0},
	}

	var rows [][]string
	for _, p := range processes {
		row := []string{
			fmt.Sprintf("%d", p.PID),
			p.Name,
			fmt.Sprintf("%.2f", p.CPU),
			fmt.Sprintf("%.2f", p.Memory),
		}
		rows = append(rows, row)
	}

	CalculateMaxWidth(rows, columns)

	var formattedRows []string
	formattedRows = append(formattedRows, FormatRow(columns, []string{"PID", "APP", "CPU %", "MEM (MB)"}))
	for _, row := range rows {
		formattedRows = append(formattedRows, FormatRow(columns, row))
	}

	v.ProcessList.Rows = formattedRows
}
