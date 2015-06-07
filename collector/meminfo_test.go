package collector

import (
	"os"
	"testing"
)

func TestMemInfoLinux(t *testing.T) {
	file, err := os.Open("fixtures/meminfo_linux")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	memInfo, err := parseMemInfoLinux(file)
	if err != nil {
		t.Fatal(err)
	}

	if want, got := 3831959552.0, memInfo["MemTotal"]; want != got {
		t.Errorf("want memory total %f, got %f", want, got)
	}

	if want, got := 3787456512.0, memInfo["DirectMap2M"]; want != got {
		t.Errorf("want memory directMap2M %f, got %f", want, got)
	}
}

func TestMemInfoFreeBSD(t *testing.T) {
	file, err := os.Open("fixtures/meminfo_freebsd")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	memInfo, err := parseMemInfoFreeBSD(file)
	if err != nil {
		t.Fatal(err)
	}

	if want, got := 16829874176.0, memInfo["hw_physmem"]; want != got {
		t.Errorf("want memory total %f, got %f", want, got)
	}

	if want, got := 46846.0, memInfo["vm_stats_vm_v_active_count"]; want != got {
		t.Errorf("want memory active %f, got %f", want, got)
	}
}
