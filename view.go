package main

import "github.com/gizak/termui"

// View can be rendered
type View interface {
	View() termui.GridBufferer
	OnEnter()
	OnUp()
	OnDown()
	OnLeft()
	OnRight()
	OnHighlight()
	OnSelect()
	OnUnHighlight()
}
