package main

import (
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

	updateFuncs := []func(*view.View){
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
			var wg sync.WaitGroup
			for _, updateFunc := range updateFuncs {
				wg.Add(1)
				go func(f func(*view.View)) {
					defer wg.Done()
					f(v)
				}(updateFunc)
			}
			wg.Wait()

			v.Render()
		}
	}
}
