package ntf

import "math"

var (
	FyneGuiMsg = make(chan interface{}, math.MaxInt8)
)

func Tell(m interface{}) {
	FyneGuiMsg <- m
}
