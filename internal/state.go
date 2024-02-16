package internal

import "sync"

type State struct {
	width  int
	height int
	zoom   int
	paused bool
	speed  int
	m      sync.RWMutex
}

func (s *State) Init() {
	s.m.Lock()
	defer s.m.Unlock()
	s.width = InitialScreenWidth
	s.height = InitialScreenHeight
	s.zoom = InitialZoom
	s.speed = InitialSpeed
	s.paused = true
}

func (s *State) Paused() bool {
	s.m.Lock()
	defer s.m.Unlock()
	return s.paused
}

func (s *State) Speed() int {
	s.m.Lock()
	defer s.m.Unlock()
	return s.speed
}

func (s *State) Width() int {
	s.m.Lock()
	defer s.m.Unlock()
	return s.width
}
func (s *State) Height() int {
	s.m.Lock()
	defer s.m.Unlock()
	return s.height
}
func (s *State) Zoom() int {
	s.m.Lock()
	defer s.m.Unlock()
	return s.zoom
}

func (s *State) Pause() {
	s.m.Lock()
	defer s.m.Unlock()
	s.paused = true
}

func (s *State) Play() {
	s.m.Lock()
	defer s.m.Unlock()
	s.paused = false
}

func (s *State) PlayPause() {
	s.m.Lock()
	defer s.m.Unlock()
	s.paused = !s.paused
}

func (s *State) ZoomIn() {
	s.m.Lock()
	defer s.m.Unlock()
	s.zoom++
}

func (s *State) ZoomOut() {
	s.m.Lock()
	defer s.m.Unlock()
	s.zoom--
}

func (s *State) CanZoomIn() bool {
	s.m.Lock()
	defer s.m.Unlock()
	return s.paused && s.zoom < MaxZoom
}

func (s *State) CanZoomOut() bool {
	s.m.Lock()
	defer s.m.Unlock()
	return s.paused && s.zoom > MinZoom
}

func (s *State) Fast() {
	s.m.Lock()
	defer s.m.Unlock()
	s.speed--
}

func (s *State) Slow() {
	s.m.Lock()
	defer s.m.Unlock()
	s.speed++
}

func (s *State) CanFast() bool {
	s.m.Lock()
	defer s.m.Unlock()
	return !s.paused && s.speed > MinSpeed
}

func (s *State) CanSlow() bool {
	s.m.Lock()
	defer s.m.Unlock()
	return !s.paused && s.speed < MaxSpeed
}
