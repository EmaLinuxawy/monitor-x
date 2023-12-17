package uiupdate

import (
	"context"
	"fmt"

	"github.com/emaLinuxawy/monitor-x/metrics"
	"github.com/emaLinuxawy/monitor-x/view"
)

func UpdateLoadAverage(ctx context.Context, v *view.View) {
	l, _ := metrics.GetLoadAverage(ctx)
	columns := []Column{
		{"1 Min", 0},
		{"5 Min", 0},
		{"15 Min", 0},
	}
	var rows [][]string
	row := []string{
		fmt.Sprintf("%.2f", l.Load1),
		fmt.Sprintf("%.2f", l.Load5),
		fmt.Sprintf("%.2f", l.Load15),
	}

	rows = append(rows, row)
	CalculateMaxWidth(rows, columns)

	var formattedRows []string
	formattedRows = append(formattedRows, FormatRow(columns, []string{"1 Min", "5 Min", "15 Min"}))
	for _, row := range rows {
		formattedRows = append(formattedRows, FormatRow(columns, row))
	}

	v.LoadAvg.Rows = formattedRows
}
