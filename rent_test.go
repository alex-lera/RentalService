package main

import "testing"

func TestRental(t *testing.T){

	result:= TestDatabase()

	if result == -1 {
		t.Error("Todo ha salido mal")
	}
}
