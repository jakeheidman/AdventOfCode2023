package main

import "testing"

func TestDay01(t *testing.T) {
	got := Part1("test_input1.txt")
	want := 142
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestGetCalibrationValue1(t *testing.T) {
	got := GetCalibrationValue("1abc2")
	want := 12
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestGetCalibrationValue2(t *testing.T) {
	got := GetCalibrationValue("pqr3stu8vwx")
	want := 38
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestGetCalibrationValue3(t *testing.T) {
	got := GetCalibrationValue("a1b2c3d4e5f")
	want := 15
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestGetCalibrationValue4(t *testing.T) {
	got := GetCalibrationValue("treb7uchet")
	want := 77
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
