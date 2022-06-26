package random

import "testing"

func TestRandomName(t *testing.T) {
	randomName := RandomName()
	isPresent := false
	for _, name := range Names {
		if randomName == name {
			isPresent = true
		}
	}

	if !isPresent {
		t.Errorf("name %q is not present in default names", randomName)
	}
}