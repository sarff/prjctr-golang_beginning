package common

import "time"

type Direction string

const (
	Left   Direction = "Left"
	Right  Direction = "Right"
	Top    Direction = "Top"
	Bottom Direction = "Bottom"
)

type HistoryItem struct {
	Time      time.Time
	Direction Direction
	ID        int
}
