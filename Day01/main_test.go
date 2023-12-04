package main

import "testing"

func TestPart1(t *testing.T) {
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

func TestPart2(t *testing.T) {
	got := Part2("test_input2.txt")
	want := 281
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

// func TestGetUpdatedCalibrationValue1(t *testing.T) {
// 	got := GetUpdatedCalibrationValue("two1nine")
// 	want := 29
// 	if got != want {
// 		t.Errorf("got %d want %d", got, want)
// 	}
// }

// func TestGetUpdatedCalibrationValue2(t *testing.T) {
// 	got := GetUpdatedCalibrationValue("eightwothree")
// 	want := 83
// 	if got != want {
// 		t.Errorf("got %d want %d", got, want)
// 	}
// }

// func TestGetUpdatedCalibrationValue3(t *testing.T) {
// 	got := GetUpdatedCalibrationValue("abcone2threexyz")
// 	want := 13
// 	if got != want {
// 		t.Errorf("got %d want %d", got, want)
// 	}
// }

// func TestGetUpdatedCalibrationValue4(t *testing.T) {
// 	got := GetUpdatedCalibrationValue("xtwone3four")
// 	want := 24
// 	if got != want {
// 		t.Errorf("got %d want %d", got, want)
// 	}
// }

// func TestGetUpdatedCalibrationValue5(t *testing.T) {
// 	got := GetUpdatedCalibrationValue("4nineeightseven2")
// 	want := 42
// 	if got != want {
// 		t.Errorf("got %d want %d", got, want)
// 	}
// }

// func TestGetUpdatedCalibrationValue6(t *testing.T) {
// 	got := GetUpdatedCalibrationValue("zoneight234")
// 	want := 14
// 	if got != want {
// 		t.Errorf("got %d want %d", got, want)
// 	}
// }

// func TestGetUpdatedCalibrationValue7(t *testing.T) {
// 	got := GetUpdatedCalibrationValue("7pqrstsixteen")
// 	want := 76
// 	if got != want {
// 		t.Errorf("got %d want %d", got, want)
// 	}
// }

func TestGetWordNumber1(t *testing.T) {
	got, err := getWordNumber("two1nine")
	want := 2
	if err != nil {
		t.Errorf(err.Error())
	} else if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestGetWordNumber2(t *testing.T) {
	got, err := getWordNumber("eightwothree")
	want := 8
	if err != nil {
		t.Errorf(err.Error())
	} else if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
