package gotest

import "testing"

func Test_add(t *testing.T) {
	if add(1,5) != 6 {
		t.Error(`add(1,5) != 6`)
	}

	if add(0,5) != 5 {
		t.Error(`add(0,5) != 5`)
	}

	if add(1000,5) != 1005 {
		t.Error(`add(1000,5) != 1005`)
	}
}

func Test_sub(t *testing.T) {
	if sub(1,5) != -4 {
		t.Error(`sub(1,5) != -4`)
	}

	if sub(0,5) != -5 {
		t.Error(`sub(0,5) != -5`)
	}

	if sub(1000,5) != 995 {
		t.Error(`sub(1000,5) != 995`)
	}
}

func Test_max(t *testing.T) {
	if max(1,5) != 5 {
		t.Error(`max(1,5) != 5`)
	}

	if max(20,5) != 20 {
		t.Error(`max(20,5) != 20`)
	}

	if max(-1,5) != 5 {
		t.Error(`max(-1,5) != 5`)
	}
}

