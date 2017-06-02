package Test


import "testing"
import "EnglishMeasurements"


func TestFractionAdd(t *testing.T) {

	fraction1 := EnglishMeasurements.Fraction{
		Numerator: 1,
		Denominator: 2,
	}
	fraction2 := EnglishMeasurements.Fraction{
		Numerator: 1,
		Denominator: 2,
	}

	fraction3 := fraction1.Add(&fraction2);

	if fraction3.Denominator != 1 {
		t.Error("Expected Denominator 1 but got ", fraction3.Denominator);
	}

	if fraction3.Numerator != 1 {
		t.Error("Expected Numerator 1 but got ", fraction3.Numerator);
	}
}
