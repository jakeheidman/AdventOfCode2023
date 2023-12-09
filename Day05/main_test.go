package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	got := Part1("test_input.txt")
	want := 35
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestCreateConversion(t *testing.T) {
	got := createConversion("50 98 2")
	want := new(conversion)
	want.destStart = 50
	want.sourceStart = 98
	want.length = 2
	if got.destStart != want.destStart {
		t.Errorf("dest error: got %d want %d", got.destStart, want.destStart)
	}
	if got.sourceStart != want.sourceStart {
		t.Errorf("source error: got %d want %d", got.sourceStart, want.sourceStart)
	}
	if got.length != want.length {
		t.Errorf("length error: got %d want %d", got.length, want.length)
	}
}

func TestTranslateSeed(t *testing.T) {
	c := createConversion("50 98 2")
	seed := 99
	got := c.translateSeed(seed)
	want := 51
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestGetSeeds(t *testing.T) {
	line := "seeds: 4106085912 135215567 529248892 159537194 1281459911 114322341 1857095529 814584370 2999858074 50388481 3362084117 37744902 3471634344 240133599 3737494864 346615684 1585884643 142273098 917169654 286257440"
	got := getSeeds(line)
	var want []int
	want = append(want, 4106085912)
	want = append(want, 135215567)
	want = append(want, 529248892)
	want = append(want, 159537194)
	want = append(want, 1281459911)
	want = append(want, 114322341)
	want = append(want, 1857095529)
	want = append(want, 814584370)
	want = append(want, 2999858074)
	want = append(want, 50388481)
	want = append(want, 3362084117)
	want = append(want, 37744902)
	want = append(want, 3471634344)
	want = append(want, 240133599)
	want = append(want, 3737494864)
	want = append(want, 346615684)
	want = append(want, 1585884643)
	want = append(want, 142273098)
	want = append(want, 917169654)
	want = append(want, 286257440)
	for idx, _ := range want {
		if want[idx] != got[idx] {
			t.Errorf("got %d want %d at idx %d", got[idx], want[idx], idx)
		}
	}
}
