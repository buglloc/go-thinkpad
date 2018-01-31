package kbd_backlight

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const SysfsKdbBacklight = "/sys/class/leds/tpacpi::kbd_backlight/brightness"

/**
 * Check if the  Keyboard Backlight is currently on
 */
func IsOn() (isOn bool) {
	b, err := Brightness()
	isOn = err == nil && b > 0
	return
}

/**
 * Probes the Keyboard Backlight driver for validity
 */
func Probe() (ok bool) {
	_, err := os.Stat(SysfsKdbBacklight)
	ok = err == nil
	return
}

/**
 * Returns current Keyboard Backlight brightness
 */
func Brightness() (brightness int, err error) {
	f, err := os.Open(SysfsKdbBacklight)
	if err != nil {
		err = errors.New("keyboard backlight not present")
		return
	}
	defer f.Close()

	status := make([]byte, 1)
	_, err = f.Read(status)
	if err != nil {
		err = fmt.Errorf("can't read brightness: %s", err.Error())
		return
	}
	brightness, err = strconv.Atoi(string(status))
	return
}
