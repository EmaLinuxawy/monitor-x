package uiupdate

import (
	"context"
	"fmt"

	"github.com/emaLinuxawy/monitor-x/metrics"
	"github.com/emaLinuxawy/monitor-x/view"
)

func UpdateLoadAverage(ctx context.Context, v *view.View) {
	l, _ := metrics.GetLoadAverage(ctx)
	header := fmt.Sprintf("%-4s | %-4s | %-4s", "1 Min", "5 Min", "15 Min")
	dataRow := fmt.Sprintf("%-5.2f | %-5.2f | %-5.2f", l.Load1, l.Load5, l.Load15)
	v.LoadAvg.Rows = []string{header, dataRow}
}
