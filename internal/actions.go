package internal

type Actions interface {
	PlayPause()
	Play()
	Pause()
	ZoomIn()
	ZoomOut()
	Restart()
	Fast()
	Slow()
	Step()
	ResolutionUp()
	ResolutionDown()
}
