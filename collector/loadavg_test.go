package collector

import "testing"

func TestLoadLinux(t *testing.T) {
	load, err := parseLoadLinux("0.21 0.37 0.39 1/719 19737")
	if err != nil {
		t.Fatal(err)
	}

	if want := 0.21; want != load {
		t.Fatalf("want load %f, got %f", want, load)
	}
}

func TestLoadFreeBSD(t *testing.T) {
	load, err := parseLoadFreeBSD("{ 0.21 0.37 0.39 }")
	if err != nil {
		t.Fatal(err)
	}

	if want := 0.21; want != load {
		t.Fatalf("want load %f, got %f", want, load)
	}
}
