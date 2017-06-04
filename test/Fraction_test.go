package Test


import "testing"
import (
	"EnglishMeasurements"
	"fmt"
)


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


func TestFractionDivide(t *testing.T) {
	fraction1 := EnglishMeasurements.Fraction{
		Numerator: 2,
		Denominator: 3,
	}

	fraction2 :=  EnglishMeasurements.Fraction{
		Numerator: 1,
		Denominator: 4,
	}

	fraction3 := fraction1.Divide(&fraction2);

	testNumDen(8, 3, fraction3, t)

}

func TestFractionMultiply(t *testing.T) {
	fraction1 := EnglishMeasurements.Fraction{
		Numerator: 2,
		Denominator: 3,
	}

	fraction2 :=  EnglishMeasurements.Fraction{
		Numerator: 5,
		Denominator: 7,
	}

	fraction3 := fraction1.Multiply(&fraction2);

	testNumDen(10, 21, fraction3, t)

}

func TestFractionSubtract(t *testing.T) {
	fraction1 := EnglishMeasurements.Fraction{
		Numerator: 3,
		Denominator: 4,
	}

	fraction2 :=  EnglishMeasurements.Fraction{
		Numerator: 8,
		Denominator: 10,
	}

	fraction3 := fraction1.Subtract(&fraction2);

	testNumDen(1, 20, fraction3, t)

}

func testNumDen(num, den int, f *EnglishMeasurements.Fraction, t *testing.T) {

	if den != f.Denominator {
		t.Error(fmt.Sprintf("Expected Denominator %d but got ", den), f.Denominator);
	}

	if f.Numerator != num {
		t.Error(fmt.Sprintf("Expected Numerator %d but got ", num), f.Numerator);
	}
}

