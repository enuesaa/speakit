package main

import (
	"fmt"
	"strings"

	"github.com/holoplot/go-evdev"
)

// 8bitdo
var keymap = map[int]string{
	304: "A",
	305: "B",
	307: "X",
	308: "Y",
	310: "L",
	311: "R",
	312: "2L",
	313: "2R",
}

var verticalmap = map[int]string{
	0: "UP",
	255: "DOWN",
}

var horizontalmap = map[int]string{
	0: "LEFT",
	255: "RIGHT",
}

func Run() {
	devpath, err := find()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", devpath)

	dev, err := evdev.Open(devpath.Path)
	if err != nil {
		panic(err)
	}
	fmt.Println(dev.Name())

	for {
		event, err := dev.ReadOne()
		if err != nil {
			panic(err)
		}

		switch event.Type {
		case evdev.EV_KEY:
			// fmt.Printf("[KEY] %d ... %d\n", event.Code, event.Value)
			if event.Value == 1 {
				value, ok := keymap[int(event.Code)]
				if ok {
					fmt.Printf("clicked: %s\n", value)
					continue
				}
			}
			// ignore event.Value == 0 as this seems `after-clicked` event
		case evdev.EV_ABS:
			if event.Code == 1 {
				value, ok := verticalmap[int(event.Value)]
				if ok {
					fmt.Printf("clicked: %s\n", value)
					continue
				}
			}
			if event.Code == 0 {
				value, ok := horizontalmap[int(event.Value)]
				if ok {
					fmt.Printf("clicked: %s\n", value)
					continue
				}
			}
		}
	}
}

func find() (*evdev.InputPath, error) {
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
