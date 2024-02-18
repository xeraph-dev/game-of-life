package internal

import (
	"fmt"
	"sync"
)

type State struct {
	res     Resolution
	zoom    int
	paused  bool
	speed   int
	m       sync.RWMutex
	_config struct {
		ok   bool
		file string
	}
}

var state State

func (s *State) Init() {
	defer func() {
		go s.save()
	}()
	s.m.Lock()
	defer s.m.Unlock()

	s.res = InitialResolution
	s.zoom = InitialZoom
	s.speed = InitialSpeed
	s.paused = true

	var err error
	config := s.config()
	if err = config.ensure(); err != nil {
		fmt.Println("error ensuring config", err)
		return
	}
	if config.load(); err != nil {
		fmt.Println("error loading config", err)
		return
	}
	s.fromConfig(config)
}

func (s *State) fromConfig(config config) {
	if !config.ok {
		return
	}
	s.res = NewResolution(config.Width, config.Height)
	s.zoom = config.Zoom
	s.speed = config.Speed
	s._config.ok = config.ok
	s._config.file = config.file
}

func (s *State) config() (config config) {
	config.Width = s.res.Width
	config.Height = s.res.Height
	config.Zoom = s.zoom
	config.Speed = s.speed
	config.ok = s._config.ok
	config.file = s._config.file
	return
}

func (s *State) save() {
	s.m.Lock()
	defer s.m.Unlock()
	if err := s.config().save(); err != nil {
		fmt.Println("saving config", err)
	}
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
	defer func() {
		go s.save()
	}()
	s.m.Lock()
	defer s.m.Unlock()
	s.zoom++
}

func (s *State) ZoomOut() {
	defer func() {
		go s.save()
	}()
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
	defer func() {
		go s.save()
	}()
	s.m.Lock()
	defer s.m.Unlock()
	s.speed--
}

func (s *State) Slow() {
	defer func() {
		go s.save()
	}()
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

func (s *State) Resolution() Resolution {
	s.m.Lock()
	defer s.m.Unlock()
	return s.res
}

func (s *State) ResolutionUp() {
	defer func() {
		go s.save()
	}()
	s.m.Lock()
	defer s.m.Unlock()
	s.res.Up()
}

func (s *State) ResolutionDown() {
	defer func() {
		go s.save()
	}()
	s.m.Lock()
	defer s.m.Unlock()
	s.res.Down()
}

func (s *State) CanResolutionUp() bool {
	s.m.Lock()
	defer s.m.Unlock()
	return s.paused && s.res.Lower(MaxResolution)
}

func (s *State) CanResolutionDown() bool {
	s.m.Lock()
	defer s.m.Unlock()
	return s.paused && s.res.Greater(MinResolution)
}
