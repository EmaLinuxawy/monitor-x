package uiupdate

import (
	"context"
	"fmt"

	"github.com/emaLinuxawy/monitor-x/metrics"
	"github.com/emaLinuxawy/monitor-x/view"
)

func UpdateDiskUsage(ctx context.Context, v *view.View) {
	usage, err := metrics.GetDiskUsage(ctx)
	if err != nil {
		fmt.Println("Error getting disk usage:", err)
		v.DiskList.Title = "Error getting disk usage"
		return
	}
	header := fmt.Sprintf("%-7s | %-6s | %-7s", "Total(GB)", "Used(GB)", "Free(GB)")
	dataRow := fmt.Sprintf("%-9.2f | %-8.2f | %-8.2f",
		float64(usage.Total)/(1024*1024*1024),
		float64(usage.Used)/(1024*1024*1024),
		float64(usage.Free)/(1024*1024*1024))

	v.DiskList.Rows = []string{header, dataRow}
}
