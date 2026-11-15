package main

import "sync"
import "fmt"

type MutextScoreboardManager struct{
	l sync.RWMutex
	scoreboard map[string]int
}