package metrics

import (
	"context"

	"github.com/shirou/gopsutil/v3/disk"
)

func GetDiskUsage(ctx context.Context) (*disk.UsageStat, error) {
	usage, err := disk.UsageWithContext(ctx, "/")
	if err != nil {
		return nil, err
	}
	return usage, nil
}
