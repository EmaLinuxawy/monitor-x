package view

import (
	"fmt"
	"runtime"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type View struct {
	Grid         *ui.Grid
	Help         *widgets.Paragraph
	CpuChart     []*widgets.Paragraph
	MemChart     *widgets.List
	ProcessList  *widgets.List
	DiskList     *widgets.List
	NetworkStats *widgets.List
	LoadAvg      *widgets.List
	CPUUsage     *widgets.List
}

func NewView() *View {
	v := &View{}

	numCPUs := runtime.NumCPU()
	v.CpuChart = make([]*widgets.Paragraph, numCPUs)
	for i := range v.CpuChart {
		v.CpuChart[i] = widgets.NewParagraph()
		v.CpuChart[i].Title = fmt.Sprintf("CPU %d", i)
	}

	v.MemChart = widgets.NewList()
	v.MemChart.Title = "Memory Usage"
	v.MemChart.TextStyle = ui.NewStyle(ui.ColorClear)
	v.MemChart.SelectedRowStyle = ui.NewStyle(ui.ColorGreen)

	v.LoadAvg = widgets.NewList()
	v.LoadAvg.Title = "Load Average"
	v.LoadAvg.TextStyle = ui.NewStyle(ui.ColorClear)
	v.LoadAvg.SelectedRowStyle = ui.NewStyle(ui.ColorGreen)

	v.CPUUsage = widgets.NewList()
	v.CPUUsage.Title = "CPU Usage"
	v.CPUUsage.TextStyle = ui.NewStyle(ui.ColorClear)
	v.CPUUsage.SelectedRowStyle = ui.NewStyle(ui.ColorGreen)

	v.ProcessList = widgets.NewList()
	v.ProcessList.Title = "Top Processes"
	v.ProcessList.TextStyle = ui.NewStyle(ui.ColorClear)
	v.ProcessList.SelectedRowStyle = ui.NewStyle(ui.ColorGreen)

	v.DiskList = widgets.NewList()
	v.DiskList.Title = "Disks Usage"
	v.DiskList.TextStyle = ui.NewStyle(ui.ColorClear)
	v.DiskList.SelectedRowStyle = ui.NewStyle(ui.ColorGreen)

	v.NetworkStats = widgets.NewList()
	v.NetworkStats.Title = "Network Statistics"
	v.NetworkStats.SelectedRowStyle = ui.NewStyle(ui.ColorGreen)
	v.NetworkStats.TextStyle = ui.NewStyle(ui.ColorClear)

	v.Help = widgets.NewParagraph()
	v.Help.Text = "PRESS q or c TO QUIT"
	v.Help.SetRect(0, 0, 50, 5)
	v.Help.TextStyle.Fg = ui.ColorRed
	v.Help.BorderStyle.Fg = ui.ColorCyan

	return v
}

func (v *View) SetLayout() {
	v.Grid = ui.NewGrid()

	cpuRow := make([]interface{}, len(v.CpuChart))
	for i, chart := range v.CpuChart {
		cpuRow[i] = ui.NewCol(1.0/float64(len(v.CpuChart)), chart)
	}
	v.Grid.Set(
		ui.NewRow(0.5/6, ui.NewCol(0.9/3.5, v.LoadAvg), ui.NewCol(0.6/5, v.CPUUsage), ui.NewCol(0.9/3.5, v.Help)),
		ui.NewRow(0.5/6, cpuRow...),
		ui.NewRow(0.5/5, ui.NewCol(0.9/2.6, v.DiskList), ui.NewCol(0.9/2.6, v.MemChart)),
		ui.NewRow(0.5/5, v.NetworkStats),
		ui.NewRow(1.0/4, v.ProcessList),
	)
}

func (v *View) ResetSize() {
	width, height := ui.TerminalDimensions()
	v.Grid.SetRect(0, 0, width, height)
}

func (v *View) Render() {
	ui.Render(v.Grid)
}
