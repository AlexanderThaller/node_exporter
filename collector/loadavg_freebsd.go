// +build !noloadavg

package collector

import (
	"bytes"
	"os/exec"
)

func getLoad1() (float64, error) {
	data := new(bytes.Buffer)
	command := exec.Command("sysctl", "-n", "vm.loadavg")
	command.Stdout = data

	err := command.Run()
	if err != nil {
		return 0, err
	}

	return parseLoadFreeBSD(data.String())
}
