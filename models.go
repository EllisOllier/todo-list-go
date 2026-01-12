package main

import "sync"

// declares the structure for todos slice
type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

// todos holds the stored in memory list of tasks
var todos = []Todo{
	{ID: 1, Task: "Make my other end points for todo-list-go"},
}

// mu is used to arrange access to the todos slice across goroutines
var mu sync.Mutex // sync.Mutex is a type not a value
