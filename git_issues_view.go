package main

import (
	"fmt"
	"log"

	"github.com/gizak/termui"
	"github.com/xanzy/go-gitlab"
)

// GitIssuesView displays issues on our git repo
type GitIssuesView struct {
	git          *gitlab.Client
	ls           *termui.List
	scroll       int
	lineSelected int
	issues       []*gitlab.Issue
	items        []string
}

func (view *GitIssuesView) reRender() {
	view.items = make([]string, len(view.issues))
	for index, issue := range view.issues {
		view.items[index] = fmt.Sprintf("%3d %s", issue.IID, issue.Title)
		if index == view.lineSelected {
			view.items[index] = fmt.Sprintf("[%s]%s", view.items[index], "(fg-red)")
		}
	}
	view.ls.Items = view.items
}

func (view *GitIssuesView) OnHighlight() {
	view.ls.BorderLabelFg = termui.ColorBlue
}

func (view *GitIssuesView) OnUnHighlight() {
	view.ls.BorderLabelFg = termui.ColorWhite
}

func (view *GitIssuesView) OnSelect() {
	view.ls.BorderLabelFg = termui.ColorGreen
}

// OnDown moves the cursor down
func (view *GitIssuesView) OnDown() {
	if view.lineSelected < len(view.issues)-1 {
		view.lineSelected++
		view.reRender()
	}
}

// OnUp moves the cursor up
func (view *GitIssuesView) OnUp() {
	if view.lineSelected > 0 {
		view.lineSelected--
		view.reRender()
	}
}

func (view GitIssuesView) OnLeft() {
}

func (view GitIssuesView) OnRight() {
}

// OnEnter selects an issue
func (view *GitIssuesView) OnEnter() {
}

// View implements View interface
func (view GitIssuesView) View() termui.GridBufferer {
	return view.ls
}

// NewGitIssuesView initializes a new instance of GitIssuesView
func NewGitIssuesView(git *gitlab.Client) *GitIssuesView {
	r := GitIssuesView{
		git:          git,
		ls:           termui.NewList(),
		scroll:       0,
		lineSelected: 0,
		issues:       nil,
		items:        []string{},
	}
	r.ls.Height = termui.TermHeight()
	//ls.BorderLabel = fmt.Sprintf("Ongoing Issues (%d)", len(issueTitles))
	r.ls.BorderLabel = "Loading Issues..."

	go func() {
		state := "opened"
		issues, _, err := git.Issues.ListProjectIssues("AsylumVR/Asylum", &gitlab.ListProjectIssuesOptions{
			ListOptions: gitlab.ListOptions{PerPage: 1000, Page: 0},
			State:       &state,
		})

		if err != nil {
			r.ls.BorderLabel = "Error!"
			log.Panic(err)
			r.ls.Items = []string{err.Error()}
		}

		r.ls.BorderLabel = fmt.Sprintf("Ongoing Issues (%d)", len(issues))
		r.issues = issues

		r.reRender()
	}()

	return &r
}
