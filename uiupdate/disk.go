package uiupdate

import (
	"context"
	"fmt"
	"time"

	"github.com/emaLinuxawy/monitor-x/metrics"
	"github.com/emaLinuxawy/monitor-x/view"
)

func UpdateDiskUsage(v *view.View) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
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
