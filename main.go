package main

import (
	"os"

	"github.com/gizak/termui"
	"github.com/joho/godotenv"
	"github.com/xanzy/go-gitlab"
)

func main() {
	if err := termui.Init(); err != nil {
		panic(err)
	}
	defer termui.Close()

	godotenv.Load()

	git := gitlab.NewClient(nil, os.Getenv("GITLAB_KEY"))

	issueView := NewGitIssuesView(git)
	workView := NewWorkWeekSummaryView()

	currentPage := NewPage([]View{issueView, workView})

	termui.Body.AddRows(
		termui.NewRow(
			termui.NewCol(6, 0, issueView.View()),
			termui.NewCol(6, 0, workView.View()),
		))

	termui.Body.Align()
	termui.Render(termui.Body)

	termui.Handle("/sys/kbd/C-x", func(termui.Event) {
		termui.StopLoop()
	})

	// termui.Handle("/sys/kbd", func(e termui.Event) {
	// 	// log.Println(e.Data)
	// 	issueTitles = append(issueTitles, e.Type)
	// 	ls.Items = issueTitles
	// 	termui.Render(ls)
	// })

	termui.Handle("/sys/kbd/<up>", func(termui.Event) {
		currentPage.OnUp()
		termui.Clear()
		termui.Render(termui.Body)
	})

	termui.Handle("/sys/kbd/<down>", func(termui.Event) {
		currentPage.OnDown()
		termui.Clear()
		termui.Render(termui.Body)
	})

	termui.Handle("/sys/kbd/<left>", func(termui.Event) {
		currentPage.OnLeft()
		termui.Clear()
		termui.Render(termui.Body)
	})

	termui.Handle("/sys/kbd/<right>", func(termui.Event) {
		currentPage.OnRight()
		termui.Clear()
		termui.Render(termui.Body)
	})

	termui.Handle("/sys/kbd/<enter>", func(termui.Event) {
		currentPage.OnEnter()
		termui.Clear()
		termui.Render(termui.Body)
	})

	termui.Handle("/sys/kbd/<escape>", func(termui.Event) {
		currentPage.OnEscape()
		termui.Clear()
		termui.Render(termui.Body)
	})

	termui.Handle("/sys/wnd/resize", func(e termui.Event) {
		termui.Body.Width = termui.TermWidth()
		termui.Body.Align()
		termui.Clear()
		termui.Render(termui.Body)
	})

	termui.Handle("/timer/1s", func(e termui.Event) {
		termui.Render(termui.Body)
	})

	termui.Loop()
}
