package uiupdate

import "fmt"

type Column struct {
	Header   string
	MaxWidth int
}

func CalculateMaxWidth(rows [][]string, columns []Column) {
	for _, row := range rows {
		for i, colData := range row {
			columns[i].MaxWidth = max(columns[i].MaxWidth, len(colData))
		}
	}
	for i := range columns {
		columns[i].MaxWidth = max(columns[i].MaxWidth, len(columns[i].Header))
	}
}

func FormatRow(columns []Column, row []string) string {
	var formattedRow string
	for i, colData := range row {
		format := fmt.Sprintf("%%-%ds", columns[i].MaxWidth)
		formattedRow += fmt.Sprintf(format, colData)
		if i < len(row)-1 {
			formattedRow += " | "
		}
	}
	return formattedRow
}
