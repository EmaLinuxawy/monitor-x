package uiupdate

import (
	"context"
	"fmt"
	"time"

	"github.com/emaLinuxawy/monitor-x/metrics"
	"github.com/emaLinuxawy/monitor-x/view"
)

func UpdateNetworkStatistics(v *view.View) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
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
