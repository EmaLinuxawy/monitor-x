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

	columns := []Column{
		{"Total(GB)", 0},
		{"Used(GB)", 0},
		{"Free(GB)", 0},
	}

	var rows [][]string
	row := []string{
		fmt.Sprintf("%.2f GB", float64(usage.Total)/(1024*1024*1024)),
		fmt.Sprintf("%.2f GB", float64(usage.Used)/(1024*1024*1024)),
		fmt.Sprintf("%.2f GB", float64(usage.Free)/(1024*1024*1024)),
	}
	rows = append(rows, row)

	CalculateMaxWidth(rows, columns)

	var formattedRows []string
	formattedRows = append(formattedRows, FormatRow(columns, []string{"Total", "Used", "Free"}))
	for _, row := range rows {
		formattedRows = append(formattedRows, FormatRow(columns, row))
	}
	v.DiskList.Rows = formattedRows
}
