// +build !nomeminfo
package collector

import "os"

func getMemInfo() (map[string]float64, error) {
	file, err := os.Open(procMemInfo)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return parseMemInfoLinux(file)
}
