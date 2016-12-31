package kio

import "testing"

func TestNewKio(t *testing.T) {

	first := Kio{
		name:   "name",
		input:  "input.in",
		output: "output.out",
	}

	second := NewKio("name", "input.in", "output.out")

	if first.name != second.name {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", first.name, second.name)
	}

	if first.input != second.input {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", first.input, second.input)
	}

	if first.output != second.output {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", first.output, second.output)
	}

}

func TestExecute(t *testing.T) {
	//todo
}
