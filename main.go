package main

import (
	"log"
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
			uiupdate.UpdateCPUData(v)
			uiupdate.UpdateTopProcesses(v)
			uiupdate.UpdateMemUsage(v)
			uiupdate.UpdateDiskUsage(v)
			uiupdate.UpdateNetworkStatistics(v)
			uiupdate.UpdateLoadAverage(v)
			uiupdate.UpdateTotalCPUUsage(v)
			v.Render()
		}
	}
}
