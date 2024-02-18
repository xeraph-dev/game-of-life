package assets

type Icon struct {
	Idle     []byte
	Disabled []byte
}

func newIcon(idle []byte) (i Icon) {
	i.Idle = idle
	return
}

func (i Icon) disabled(disabled []byte) Icon {
	i.Disabled = disabled
	return i
}

var (
	PlayIcon       = newIcon(playFilled).disabled(playOutline)
	PauseIcon      = newIcon(pauseFilled).disabled(pauseOutline)
	PlusIcon       = newIcon(plusFilled).disabled(plusOutline)
	MinusIcon      = newIcon(minusFilled).disabled(minusOutline)
	SlowIcon       = newIcon(slowFilled).disabled(slowOutline)
	FastIcon       = newIcon(fastFilled).disabled(fastOutline)
	StepIcon       = newIcon(stepFilled).disabled(stepOutline)
	RestartIcon    = newIcon(restart).disabled(restart)
	KeyRIcon       = newIcon(keyR)
	KeyCommaIcon   = newIcon(keyComma)
	KeyPeriodIcon  = newIcon(keyPeriod)
	KeyGreaterIcon = newIcon(keyGreater)
	KeyLowerIcon   = newIcon(keyLower)
	KeyPlusIcon    = newIcon(keyPlus)
	KeyMinusIcon   = newIcon(keyMinus)
	KeySpaceIcon   = newIcon(keySpace)
	LeftIcon       = newIcon(leftFilled).disabled(leftOutline)
	RightIcon      = newIcon(rightFilled).disabled(rightOutline)
)
