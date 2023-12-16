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

	var rows []string
	header := fmt.Sprintf("%-15s | %-15s | %-15s | %-15s | %-15s", "Interface Name", "Bytes Sent(Gbits)", "Bytes Received(Gbits)", "Packets Sent", "Packets Received")
	rows = append(rows, header)
	for _, stat := range networkStats {
		row := fmt.Sprintf("%-15s | %-17.3f | %-21.3f | %-15d | %-15d", stat.InterfaceName, float64(stat.BytesSent)*8/1e9, float64(stat.BytesRecv)*8/1e9, stat.PacketsSent, stat.PacketsRecv)
		rows = append(rows, row)
	}
	v.NetworkStats.Rows = rows
}
