package dock

import (
	"bufio"
	"os"
)

const IBMDockDocked = "/sys/devices/platform/dock.2/docked"
const IBMDockModAlias = "/sys/devices/platform/dock.2/modalias"
const IBMDockId = "acpi:IBM0079:PNP0C15:LNXDOCK:"

/**
 * Check if the ThinkPad is physically docked
 * into the UltraDock or the UltraBase
 * @return true if the ThinkPad is inside the dock
 */
func IsDocked() (docked bool) {
	f, err := os.Open(IBMDockDocked)
	if err != nil {
		return
	}
	defer f.Close()

	status := make([]byte, 1)
	_, err = f.Read(status)
	if err != nil {
		return
	}
	docked = status[0] == '1'
	return
}

/**
 * Probes the dock if it is an IBM dock and if the
 * dock is sane and ready for detection/state changes
* return true if the dock is sane and valid
*/
func Probe() (ok bool) {
	f, err := os.Open(IBMDockModAlias)
	if err != nil {
		return
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	data, err := reader.ReadString('\n')
	ok = data == IBMDockId
	return
}
