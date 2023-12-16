package uiupdate

import (
	"context"
	"fmt"

	"github.com/emaLinuxawy/monitor-x/metrics"
	"github.com/emaLinuxawy/monitor-x/view"
)

func UpdateMemUsage(ctx context.Context, v *view.View) {
	usage, err := metrics.GetMemoryUsed(ctx)
	if err != nil {
		fmt.Println("Error getting memory usage:", err)
		v.MemChart.Title = "Error getting memory usage"
		return
	}
	total := float64(usage.Total) / (1024 * 1024 * 1024)
	used := float64(usage.Used) / (1024 * 1024 * 1024)
	free := total - used
	header := fmt.Sprintf("%-5s | %-5s | %-5s", "Total(GB)", "Used(GB)", "Free(GB)")
	dataRow := fmt.Sprintf("%-9.2f | %-8.2f | %-6.2f", total, used, free)
	v.MemChart.Rows = []string{header, dataRow}
}
