package internal

const ScreenWidth = 640
const ScreenHeight = 480

type State struct {
	Zoom         int
	Paused       bool
	ScreenWidth  int
	ScreenHeight int
}

var state State

func init() {
	state.Zoom = 1
	state.Paused = true
	state.ScreenWidth = ScreenWidth
	state.ScreenHeight = ScreenHeight
}
