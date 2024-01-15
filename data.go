package main

import "github.com/charmbracelet/bubbles/list"

// Provides the mock data to fill the kanban board

func (b *Board) initLists() {
	b.cols = []column{
		newColumn(todo),
		newColumn(inProgress),
		newColumn(done),
	}
	// Init To Do
	b.cols[todo].list.Title = "To Do"
	b.cols[todo].list.SetItems([]list.Item{
		Task{status: todo, title: "learn Rust", description: "rust programming language"},
		Task{status: todo, title: "eat sushi", description: "negitoro roll, miso soup, rice"},
		Task{status: todo, title: "fold laundry", description: "or wear wrinkly t-shirts"},
		Task{status: todo, title: "read a programming book", description: "choose a bestseller on Amazon and start reading"},
		Task{status: todo, title: "complete a coding challenge", description: "find a coding challenge on a platform like HackerRank or LeetCode and solve it"},
	})
	// Init in progress
	b.cols[inProgress].list.Title = "In Progress"
	b.cols[inProgress].list.SetItems([]list.Item{
		Task{status: inProgress, title: "write code", description: "don't worry, it's Go"},
		Task{status: inProgress, title: "build a CLI app in Go", description: "follow a tutorial and build a command-line application"},
	})
	// Init done
	b.cols[done].list.Title = "Done"
	b.cols[done].list.SetItems([]list.Item{
		Task{status: done, title: "stay cool", description: "as a cucumber"},
		Task{status: done, title: "contribute to an open-source project", description: "find a project on GitHub and make a contribution"},
	})
}
