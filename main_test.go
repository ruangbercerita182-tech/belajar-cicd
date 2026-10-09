package main

import "testing"

func TestTambah(t *testing.T) {
	hasil := tambah(10, 20)

	if hasil != 30 {
		t.Errorf("ingin 30, mendapat %d", hasil)
	}
}
