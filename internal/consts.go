package internal

const Title = "Game of Life"
const PackageName = "game-of-life"

const InitialZoom = 2
const InitialSpeed = 1

const MaxZoom = 5
const MinZoom = 1
const MaxSpeed = 5
const MinSpeed = 1

var InitialResolution = Resolution{640, 480}

var Resolutions, ResolutionList = InitializeResolutionsMap([]Resolution{
	{640, 480},
	{800, 600},
	{960, 720},
	{1280, 960},
	{1600, 1200},
	{1920, 1440},
	{2560, 1920},
	{3840, 2880},
	{5120, 3840},
	{7680, 5760},
})

var MaxResolution = Resolutions[ResolutionList[len(ResolutionList)-1]]
var MinResolution = Resolutions[ResolutionList[0]]
