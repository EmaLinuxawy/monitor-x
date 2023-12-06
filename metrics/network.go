package metrics

import (
	"context"

	"github.com/shirou/gopsutil/v3/net"
)

type NetworkStats struct {
	InterfaceName string
	BytesSent     uint64
	BytesRecv     uint64
	PacketsSent   uint64
	PacketsRecv   uint64
}

func GetNetworkStatistics(ctx context.Context) ([]NetworkStats, error) {
	ioCounters, err := net.IOCounters(true)
	if err != nil {
		return nil, err
	}

	var stats []NetworkStats
	const threshold uint64 = 0.1 * 1e9 / 8
	for _, counter := range ioCounters {
		if counter.BytesSent > threshold && counter.BytesRecv > threshold && counter.PacketsSent > 0 && counter.PacketsRecv > 0 {
			stats = append(stats, NetworkStats{
				InterfaceName: counter.Name,
				BytesSent:     counter.BytesSent,
				BytesRecv:     counter.BytesRecv,
				PacketsSent:   counter.PacketsSent,
				PacketsRecv:   counter.PacketsRecv,
			})
		}
	}
	return stats, nil
}
