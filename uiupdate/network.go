package uiupdate

import (
	"context"
	"fmt"

	"github.com/emaLinuxawy/monitor-x/metrics"
	"github.com/emaLinuxawy/monitor-x/view"
)

func UpdateNetworkStatistics(ctx context.Context, v *view.View) {
	networkStats, err := metrics.GetNetworkStatistics(ctx)
	if err != nil {
		fmt.Println("Error getting network statistics:", err)
		v.NetworkStats.Title = "Error getting network statistics"
		return
	}
	columns := []Column{
		{"Interface Name", 0},
		{"Bytes Sent", 0},
		{"Bytes Received", 0},
		{"Packets Sent", 0},
		{"Packets Received", 0},
	}

	var rows [][]string
	for _, stat := range networkStats {
		row := []string{stat.InterfaceName,
			fmt.Sprintf("%.2f Gbits", float64(stat.BytesSent)*8/1e9),
			fmt.Sprintf("%.2f Gbits", float64(stat.BytesRecv)*8/1e9),
			fmt.Sprintf("%d", stat.PacketsSent),
			fmt.Sprintf("%d", stat.PacketsRecv),
		}
		rows = append(rows, row)
	}

	CalculateMaxWidth(rows, columns)
	var formattedRows []string
	formattedRows = append(formattedRows, FormatRow(columns, []string{"Interface Name", "Bytes Sent", "Bytes Received", "Packets Sent", "Packets Received"}))
	for _, row := range rows {
		formattedRows = append(formattedRows, FormatRow(columns, row))
	}
	v.NetworkStats.Rows = formattedRows
}
