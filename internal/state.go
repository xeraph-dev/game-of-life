package internal

type State struct {
	zoom   int
	paused bool
}

var state State

func init() {
	state.zoom = 2
	state.paused = true
}
