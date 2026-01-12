package main

import "sync"

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var todos = []Todo{
	{ID: 1, Task: "Make my other end points for todo-list-go"},
}

var mu sync.Mutex // sync.Mutex is a type not a value
