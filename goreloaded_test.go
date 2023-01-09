package main

import "testing"

func TestGoreloaded(t *testing.T) {
    actualString := "10 (bin) files were added, and 1E (hex) arts were made available. It has been 10 (bin) years, my life is in 10 (hex) months."
	answer := Goreloaded(actualString)
    expectedString := "2 files were added, and 30 arts were made available. It has been 2 years, my life is in 16 months."
    if answer != expectedString{
        t.Errorf("Expected String(%s) is not same as"+
         " actual string (%s)", expectedString,actualString)
    }
}
