package main

// Page contains different views to render
type Page struct {
	views           []View
	highlightedView int
	viewSelected    bool
}

// NewPage creates a page to render
func NewPage(views []View) *Page {
	if views == nil {
		panic("Can't have nil views for a page")
	}
	r := Page{
		views:           views,
		highlightedView: 0,
		viewSelected:    false,
	}
	r.updateHighlightsAndSelected()
	return &r
}

func (page Page) updateHighlightsAndSelected() {
	for index, view := range page.views {
		if index == page.highlightedView {
			view.OnHighlight()
		} else {
			view.OnUnHighlight()
		}
	}
}

// OnDown moves the cursor down
func (page *Page) OnDown() {
	if page.viewSelected {
		page.views[page.highlightedView].OnDown()
	} else if page.highlightedView > 0 {
		page.highlightedView--
		page.updateHighlightsAndSelected()
	}
}

// OnUp moves the cursor up
func (page *Page) OnUp() {
	if page.viewSelected {
		page.views[page.highlightedView].OnUp()
	} else if page.highlightedView < len(page.views)-1 {
		page.highlightedView++
		page.updateHighlightsAndSelected()
	}
}

func (page Page) OnLeft() {
}

func (page Page) OnRight() {
}

// OnEnter
func (page *Page) OnEnter() {
	if page.viewSelected {
		page.views[page.highlightedView].OnEnter()
	} else {
		page.viewSelected = true
		page.views[page.highlightedView].OnSelect()
	}
}

func (page *Page) OnEscape() {
	page.viewSelected = false
	page.updateHighlightsAndSelected()
}
