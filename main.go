package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/emaLinuxawy/monitor-x/uiupdate"
	"github.com/emaLinuxawy/monitor-x/view"
	ui "github.com/gizak/termui/v3"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to initialize termui: %v", err)
	}
	defer ui.Close()

	v := view.NewView()
	v.SetLayout()
	v.ResetSize()

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C

	updateFuncs := []func(context.Context, *view.View){
		uiupdate.UpdateCPUData,
		uiupdate.UpdateTopProcesses,
		uiupdate.UpdateMemUsage,
		uiupdate.UpdateDiskUsage,
		uiupdate.UpdateNetworkStatistics,
		uiupdate.UpdateLoadAverage,
		uiupdate.UpdateTotalCPUUsage,
	}

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "c":
				return
			case "<Resize>":
				v.ResetSize()
			}
		case <-ticker:
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			var wg sync.WaitGroup
			for _, updateFunc := range updateFuncs {
				wg.Add(1)
				go func(f func(context.Context, *view.View)) {
					defer wg.Done()
					f(ctx, v)
				}(updateFunc)
			}
			wg.Wait()
			cancel()

			v.Render()
		}
	}
}
