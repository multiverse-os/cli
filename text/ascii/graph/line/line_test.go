package line

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestLoadDataValid(t *testing.T) {
	testCases := []struct {
		testData string
		expected *dataSet
	}{
		{"x,y", &dataSet{[]point{}, "x", "y", 0, 0, 0, 0}},
		{"x,y\n", &dataSet{[]point{}, "x", "y", 0, 0, 0, 0}},
		{"x,y,ignored", &dataSet{[]point{}, "x", "y", 0, 0, 0, 0}},
		{"x,y\n1,2", &dataSet{[]point{{1, 2}}, "x", "y", 1, 1, 2, 2}},
		{"x,y\n1,2\n3,4", &dataSet{[]point{{1, 2}, {3, 4}}, "x", "y", 1, 3, 2, 4}},
		{"x,y\n3,4\n1,2", &dataSet{[]point{{3, 4}, {1, 2}}, "x", "y", 1, 3, 2, 4}},
		{"x,y\n3,4\n1,2\n2,8", &dataSet{[]point{{3, 4}, {1, 2}, {2, 8}}, "x", "y", 1, 3, 2, 8}},
	}

	for _, test := range testCases {
		if r, err := loadData(strings.NewReader(test.testData)); err != nil {
			t.Errorf("Expected: %v\n Actual Error: %v", test.expected, err)
		} else {
			if !reflect.DeepEqual(test.expected, r) {
				t.Errorf("dataSet for testData: %s was incorrect!\nExpected: %v\nActual: %v", test.testData, test.expected, r)
			}
		}
	}
}

func TestLoadDataInvalid(t *testing.T) {
	testCases := []struct {
		testData string
		expected error
	}{
		{"", &ErrRow{0, "", errors.New("No data found")}},
		{"x", &ErrRow{0, "x", errors.New("Header with 2 elements required")}},
		{"x,", &ErrRow{0, "x,", errors.New("Header with 2 elements required")}},
		{"x,y\n1", &ErrRow{1, "1", errors.New("coordinates require 2 values")}},
		{"x,y\n1,2\n3", &ErrRow{2, "3", errors.New("coordinates require 2 values")}},
		{"x,y\ngarbagex,2", &ErrRow{1, "garbagex,2", errors.New("garbagex is not a number")}},
		{"x,y\n2,garbagey", &ErrRow{1, "2,garbagey", errors.New("garbagey is not a number")}},
	}

	for _, test := range testCases {
		if _, err := loadData(strings.NewReader(test.testData)); err != nil {
			if test.expected.Error() != err.Error() {
				t.Errorf("Expected: %v\n Actual: %v", test.expected, err)
			}
		} else {
			t.Errorf("Expected error, but none found!")
		}
	}
}

func TestParseRowValid(t *testing.T) {
	testCases := []struct {
		row      string
		expected point
	}{
		{"1,2", point{1, 2}},
		{"1123,2098", point{1123, 2098}},
	}

	for _, test := range testCases {
		if r, err := parseRow(test.row); err != nil {
			t.Errorf("Expected: %v\n Actual Error: %v", test.expected, err)
		} else {
			if !(test.expected.x == r.x) || !(test.expected.y == r.y) {
				t.Errorf("Expected: %v\n Actual: %v", test.expected, r)
			}
		}
	}
}

func TestParseRowInvalid(t *testing.T) {
	testCases := []struct {
		row      string
		expected error
	}{
		{"1", errors.New("coordinates require 2 values")},
		{"foo", errors.New("coordinates require 2 values")},
		{"foo,1", errors.New("foo is not a number")},
		{"1,bar", errors.New("bar is not a number")},
		{"foo,1,2", errors.New("foo is not a number")},
		{"1,bar,2", errors.New("bar is not a number")},
		{"foo,bar", errors.New("foo is not a number")},
		{"1.0,2", errors.New("1.0 is not a number")},
		{"1,2.0", errors.New("2.0 is not a number")},
	}

	for _, test := range testCases {
		if _, err := parseRow(test.row); err != nil {
			if test.expected.Error() != err.Error() {
				t.Errorf("Expected: %v\n Actual: %v", test.expected, err)
			}
		} else {
			t.Errorf("Expected error, but none found!")
		}
	}
}
