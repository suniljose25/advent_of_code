package day05

import "testing"

func TestHasAtleastThreeVowels1(t *testing.T) {
	if !hasAtleastThreeVowels("aei") {
		t.Fail()
	}
}

func TestHasAtleastThreeVowels2(t *testing.T) {
	if !hasAtleastThreeVowels("xazegov") {
		t.Fail()
	}
}

func TestHasAtleastThreeVowels3(t *testing.T) {
	if !hasAtleastThreeVowels("aeiouaeiouaeiou") {
		t.Fail()
	}
}

func TestHasAtleastThreeVowels4(t *testing.T) {
	if hasAtleastThreeVowels("qwerty") {
		t.Fail()
	}
}

func TestHasAtleastThreeVowels5(t *testing.T) {
	if hasAtleastThreeVowels("qwertyo") {
		t.Fail()
	}
}

func TestHasContigousRunesReturnsTrueForValidInput(t *testing.T) {
	if !hasContigousRunes("xx") {
		t.Fail()
	}

	if !hasContigousRunes("abcdde") {
		t.Fail()
	}

	if !hasContigousRunes("aabbccdd") {
		t.Fail()
	}
}

func TestHasContigousRunesReturnsFalseForInvalidInput(t *testing.T) {
	if hasContigousRunes("abacdeaf") {
		t.Fail()
	}
}

func TestHasRestrictedStringsReturnsTrueForInvalidInput(t *testing.T) {
	if !hasRestrictedStrings("abacdeaf") {
		t.Fail()
	}

	if !hasRestrictedStrings("aacdeaf") {
		t.Fail()
	}

	if !hasRestrictedStrings("aacpqeaf") {
		t.Fail()
	}

	if !hasRestrictedStrings("aacxyeaf") {
		t.Fail()
	}
}

func TestHasRestrictedStringsReturnsFalseForValidInput(t *testing.T) {
	if hasRestrictedStrings("aandeaf") {
		t.Fail()
	}

	if hasRestrictedStrings("aaceaf") {
		t.Fail()
	}

	if hasRestrictedStrings("aacqeaf") {
		t.Fail()
	}

	if hasRestrictedStrings("aacyeaf") {
		t.Fail()
	}
}

func TestNiceString(t *testing.T) {
	if !isNiceString("ugknbfddgicrmopn") {
		t.Fail()
	}

	if !isNiceString("aaa") {
		t.Fail()
	}
}

func TestNaughtyString(t *testing.T) {
	if isNiceString("jchzalrnumimnmhp") {
		t.Fail()
	}

	if isNiceString("haegwjzuvuyypxyu") {
		t.Fail()
	}

	if isNiceString("dvszwmarrgswjxmb") {
		t.Fail()
	}
}
