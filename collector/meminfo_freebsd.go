// +build !nomeminfo
package collector

import (
	"bytes"
	"os/exec"
)

func getMemInfo() (map[string]float64, error) {
	data := new(bytes.Buffer)

	command := exec.Command("sysctl", "hw.physmem")
	command.Stdout = data

	err := command.Run()
	if err != nil {
		return nil, err
	}

	command = exec.Command("sysctl", "hw.pagesize")
	command.Stdout = data

	err = command.Run()
	if err != nil {
		return nil, err
	}

	command = exec.Command("sysctl", "vm.stats.vm")
	command.Stdout = data

	err = command.Run()
	if err != nil {
		return nil, err
	}

	return parseMemInfoFreeBSD(data)
}
