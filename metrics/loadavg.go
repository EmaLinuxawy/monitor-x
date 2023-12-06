package metrics

import (
	"context"

	"github.com/shirou/gopsutil/v3/load"
)

func GetLoadAverage(ctx context.Context) (*load.AvgStat, error) {
	avg, err := load.AvgWithContext(ctx)
	if err != nil {
		return nil, err
	}
	return avg, nil
}
