package main

import (
	"github.com/gizak/termui"
)

// WorkWeekSummaryView displays issues on our git repo
type WorkWeekSummaryView struct {
	ls *termui.List
}

func (view *WorkWeekSummaryView) reRender() {
	view.ls.Items = []string{"A test"}
}

func (view *WorkWeekSummaryView) OnHighlight() {
	view.ls.BorderLabelFg = termui.ColorBlue
}

func (view *WorkWeekSummaryView) OnUnHighlight() {
	view.ls.BorderLabelFg = termui.ColorWhite
}

func (view *WorkWeekSummaryView) OnSelect() {
	view.ls.BorderLabelFg = termui.ColorGreen
}

// OnDown moves the cursor down
func (view *WorkWeekSummaryView) OnDown() {

}

// OnUp moves the cursor up
func (view *WorkWeekSummaryView) OnUp() {
}

func (view WorkWeekSummaryView) OnLeft() {
}

func (view WorkWeekSummaryView) OnRight() {
}

// OnEnter selects an issue
func (view *WorkWeekSummaryView) OnEnter() {
}

// View implements View interface
func (view WorkWeekSummaryView) View() termui.GridBufferer {
	return view.ls
}

// NewWorkWeekSummaryView initializes a new instance of NewWorkWeekSummaryView
func NewWorkWeekSummaryView() *WorkWeekSummaryView {
	r := WorkWeekSummaryView{
		ls: termui.NewList(),
	}
	r.ls.Height = termui.TermHeight()
	r.ls.BorderLabel = "Work Week Summary"

	r.reRender()

	return &r
}
