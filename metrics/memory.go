package metrics

import (
	"context"

	"github.com/shirou/gopsutil/v3/mem"
)

func GetMemoryUsed(ctx context.Context) (*mem.VirtualMemoryStat, error) {
	percentages, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return nil, err
	}
	return percentages, nil
}
