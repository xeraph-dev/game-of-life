package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
)

type config struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Zoom   int `json:"zoom"`
	Speed  int `json:"speed"`
	file   string
	ok     bool
}

func (c *config) ensure() (err error) {
	defer func(c *config) {
		c.ok = err == nil
	}(c)

	var userConfigDir string
	var f *os.File

	if userConfigDir, err = os.UserConfigDir(); err != nil {
		err = fmt.Errorf("getting the user config dir: %w", err)
		return
	}

	configDir := path.Join(userConfigDir, PackageName)
	if err = os.MkdirAll(configDir, os.ModePerm); err != nil {
		err = fmt.Errorf("creating game config directory: %w", err)
		return
	}

	c.file = path.Join(configDir, "config.json")
	if _, err = os.Stat(c.file); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			err = fmt.Errorf("getting config file stat: %w", err)
			return
		}

		if f, err = os.Create(c.file); err != nil {
			err = fmt.Errorf("creating config file: %w", err)
			return
		}
		if err = f.Close(); err != nil {
			err = fmt.Errorf("closing config file: %w", err)
			return
		}
	}

	return
}

func (c *config) load() (err error) {
	if !c.ok {
		return
	}

	var configData []byte

	if configData, err = os.ReadFile(c.file); err != nil {
		err = fmt.Errorf("reading config file: %w", err)
		return
	}

	if err = json.Unmarshal(configData, c); err != nil {
		err = fmt.Errorf("decoding config data: %w", err)
		return
	}

	return
}

func (c config) save() (err error) {
	if !c.ok {
		return
	}

	var data []byte

	if data, err = json.Marshal(c); err != nil {
		err = fmt.Errorf("encoding config data: %w", err)
		return
	}

	if err = os.WriteFile(c.file, data, os.ModePerm); err != nil {
		err = fmt.Errorf("writing config file: %w", err)
		return
	}

	return
}
