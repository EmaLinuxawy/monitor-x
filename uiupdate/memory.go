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

	columns := []Column{
		{"Total", 0},
		{"Used", 0},
		{"Free", 0},
	}

	var rows [][]string
	row := []string{
		fmt.Sprintf("%.2f GB", float64(usage.Total)/(1024*1024*1024)),
		fmt.Sprintf("%.2f GB", float64(usage.Used)/(1024*1024*1024)),
		fmt.Sprintf("%.2f GB", float64(usage.Total)/(1024*1024*1024)-float64(usage.Used)/(1024*1024*1024)),
	}
	rows = append(rows, row)

	CalculateMaxWidth(rows, columns)

	var formattedRows []string
	formattedRows = append(formattedRows, FormatRow(columns, []string{"Total", "Used", "Free"}))
	for _, row := range rows {
		formattedRows = append(formattedRows, FormatRow(columns, row))
	}

	v.MemChart.Rows = formattedRows
}
