package eightbitctl

import (
	"fmt"
	"strings"
	"time"

	"github.com/holoplot/go-evdev"
)

// TODO: make behavior
func New() *Eightbit {
	return &Eightbit{
		listeners: []Listener{},
	}
}

type Listener func(KeyCode)

type Eightbit struct {
	listeners []Listener
}

func (e *Eightbit) On(listener func(KeyCode)) {
	e.listeners = append(e.listeners, listener)
}

func (e *Eightbit) StartWait() error {
	var err error

	for range 120 {
		err = e.Start()
		if err == nil {
			fmt.Println("connected")
			return nil
		}
		time.Sleep(1 * time.Second)
	}
	return err
}

func (e *Eightbit) Start() error {
	devpath, err := e.find()
	if err != nil {
		return err
	}
	fmt.Println("find: ", devpath)

	dev, err := evdev.Open(devpath.Path)
	if err != nil {
		return err
	}
	go e.listen(dev)

	return nil
}

func (e *Eightbit) find() (*evdev.InputPath, error) {
	devpaths, err := evdev.ListDevicePaths()
	if err != nil {
		return nil, fmt.Errorf("failed to list")
	}
	for _, devpath := range devpaths {
		if strings.Contains(devpath.Name, "8BitDo") {
			return &devpath, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func (e *Eightbit) listen(dev *evdev.InputDevice) error {
	devname, err := dev.Name()
	if err != nil {
		return err
	}
	fmt.Println("listen: ", devname)

	for {
		event, err := dev.ReadOne()
		if err != nil {
			return err
		}

		if event.Type == evdev.EV_KEY && event.Value == 1 {
			value, ok := keymap[int(event.Code)]
			if ok {
				for _, listener := range e.listeners {
					listener(value)
				}
			}
		}
		if event.Type == evdev.EV_ABS && event.Code == 0 {
			value, ok := horizontalmap[int(event.Value)]
			if ok {
				for _, listener := range e.listeners {
					listener(value)
				}
			}
		}
		if event.Type == evdev.EV_ABS && event.Code == 1 {
			value, ok := verticalmap[int(event.Value)]
			if ok {
				for _, listener := range e.listeners {
					listener(value)
				}
			}
		}
	}
}
