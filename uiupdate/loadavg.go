package uiupdate

import (
	"context"
	"fmt"
	"time"

	"github.com/emaLinuxawy/monitor-x/metrics"
	"github.com/emaLinuxawy/monitor-x/view"
)

func UpdateLoadAverage(v *view.View) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	l, _ := metrics.GetLoadAverage(ctx)
	header := fmt.Sprintf("%-4s | %-4s | %-4s", "1 Min", "5 Min", "15 Min")
	dataRow := fmt.Sprintf("%-5.2f | %-5.2f | %-5.2f", l.Load1, l.Load5, l.Load15)
	v.LoadAvg.Rows = []string{header, dataRow}
}
