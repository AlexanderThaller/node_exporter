// +build !noloadavg

package collector

import "io/ioutil"

func getLoad1() (float64, error) {
	data, err := ioutil.ReadFile(procLoad)
	if err != nil {
		return 0, err
	}
	return parseLoadLinux(string(data))
}
