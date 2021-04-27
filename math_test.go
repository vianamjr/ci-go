package main

import "testing"

func TestSum(t *testing.T) {

	total := sum(15, 15)
	expected := 30

	if total != expected {
		t.Errorf("Resultado da soma é invárido: expeted %d; got: %d", expected, total)
	}
}
