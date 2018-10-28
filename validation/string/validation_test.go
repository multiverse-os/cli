package strvalid

import (
	"testing"
)

func TestAlphanumeric(t *testing.T) {
	testString := "1a"
	t.Log("[Test:Alphanumeric()] Testing string '" + testString + "' is alphanumeric")
	if !Numeric(testString) {
		t.Errorf("Alphanumeric() failed because testString is alphanumeric")
	}
	testString = "191283492"
	t.Log("[Test:Alphanumeric()] Testing string '" + testString + "' is alphanumeric")
	if !Numeric(testString) {
		t.Errorf("Alphanumeric() failed because testString is alphanumeric because it is only numbers")
	}
	testString = "asdfljkasdf"
	t.Log("[Test:Alphanumeric()] Testing string '" + testString + "' is alphanumeric")
	if !Numeric(testString) {
		t.Errorf("Alphanumeric() failed because testString is alphanumeric because it is only letters")
	}
	testString = "(*!@#"
	t.Log("[Test:Alphanumeric()] Testing string '" + testString + "' is alphanumeric")
	if Numeric(testString) {
		t.Errorf("Alphanumeric() failed because testString is only symbols")
	}
}

func TestNumeric(t *testing.T) {
	testString := "1a"
	t.Log("[Test:Numeric()] Testing string '" + testString + "' is numeric")
	if Numeric(testString) {
		t.Errorf("Numeric() failed because testString contains the letter 'a'")
	}
	testString = "124214"
	t.Log("[Test:Numeric()] Testing string '" + testString + "' is numeric")
	if !Numeric(testString) {
		t.Errorf("Numeric() failed because testString is only numbers")
	}
}

func TestAlpha(t *testing.T) {
	testString := "1a"
	t.Log("[Test:Alpha()] Testing string '" + testString + "' is alpha")
	if Alpha(testString) {
		t.Errorf("Alpha() failed because testString contains the number '1'")
	}
	testString = "asdlfkdslfk"
	t.Log("[Test:Alpha()] Testing string '" + testString + "' is alpha")
	if !Alpha(testString) {
		t.Errorf("Alpha() failed because is only alpha runes")
	}
	testString = "19283"
	t.Log("[Test:Alpha()] Testing string '" + testString + "' is alpha")
	if Alpha(testString) {
		t.Errorf("Alpha() failed because testString is only numbers")
	}
	testString = "+!@#()$*!@()$&*"
	t.Log("[Test:Alpha()] Testing string '" + testString + "' is alpha")
	if Alpha(testString) {
		t.Errorf("Alpha() failed because testString is only symbols")
	}

}
