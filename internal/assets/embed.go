package assets

import _ "embed"

type Icon struct {
	Idle     []byte
	Disabled []byte
}

func newIcon(idle []byte, disabled []byte) Icon {
	return Icon{
		Idle:     idle,
		Disabled: disabled,
	}
}

var (
	PlayIcon    = newIcon(playFilled, playOutline)
	PauseIcon   = newIcon(pauseFilled, pauseOutline)
	PlusIcon    = newIcon(plusFilled, plusOutline)
	MinusIcon   = newIcon(minusFilled, minusOutline)
	RestartIcon = newIcon(restart, restart)
)

//go:embed icons/play_filled.png
var playFilled []byte

//go:embed icons/play_outline.png
var playOutline []byte

//go:embed icons/pause_filled.png
var pauseFilled []byte

//go:embed icons/pause_outline.png
var pauseOutline []byte

//go:embed icons/plus_filled.png
var plusFilled []byte

//go:embed icons/plus_outline.png
var plusOutline []byte

//go:embed icons/minus_filled.png
var minusFilled []byte

//go:embed icons/minus_outline.png
var minusOutline []byte

//go:embed icons/restart.png
var restart []byte
