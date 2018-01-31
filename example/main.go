package main

import (
	"fmt"

	"github.com/buglloc/go-thinkpad/acpi"
	"github.com/buglloc/go-thinkpad/dock"
	"github.com/buglloc/go-thinkpad/kbd-backlight"
)

func main() {
	brightness, err := kbd_backlight.Brightness()
	if err != nil {
		fmt.Printf("failed to get KBD brightness: %s\n", err)
	} else {
		fmt.Printf("KBD brightness: %d\n", brightness)
	}

	if dock.IsDocked() {
		fmt.Println("Docked")
	} else {
		fmt.Println("Undocked")
	}

	acpi.Debug = true
	err = acpi.HandleMulti(func(eventType acpi.EventType) error {
		fmt.Printf("event: %d\n", eventType)
		return nil
	})

	if err != nil {
		panic(err)
	}
}
