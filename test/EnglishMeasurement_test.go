package Test

import "testing"
import (
	. "EnglishMeasurements"
	"fmt"
)

func TestAdd(t *testing.T) {
	em := NewEnglishMeasurement(4, 6, &Fraction{
		Numerator: 1,
		Denominator: 2,
	})

	em2 := NewEnglishMeasurement(8,4, &Fraction{
		Numerator: 1,
		Denominator: 4,
	})

	em3 := em.Add(em2);

	fmt.Println(em3.GetFractionInch())
	testEngMes(12, 10, Fraction{3,4}, em3, t)

}

func TestSubtract(t *testing.T) {
	em := NewEnglishMeasurement(4, 6, &Fraction{
		Numerator: 1,
		Denominator: 2,
	})

	em2 := NewEnglishMeasurement(8,4, &Fraction{
		Numerator: 1,
		Denominator: 4,
	})

	em3 := em2.Subtract(em);

	fmt.Println(em3.GetFractionInch())
	testEngMes(3, 9, Fraction{3,4}, em3, t)

}
func testEngMes(foot, inch int, frac Fraction, em EnglishMeasurement, t *testing.T) {
	if foot != em.GetFoot() {
		t.Error(fmt.Sprintf("Expected Foot %d but got ", foot), em.GetFoot());
	}

	if inch != em.GetInch() {
		t.Error(fmt.Sprintf("Expected Inch %d but got ", inch), em.GetInch());
	}

	if frac.Denominator != em.GetFractionInch().Denominator {
		t.Error(fmt.Sprintf("Expected Denominator %d but got ", frac.Denominator), em.GetFractionInch().Denominator);
	}

	if frac.Numerator != em.GetFractionInch().Numerator {
		t.Error(fmt.Sprintf("Expected Numerator %d but got ", frac.Numerator), em.GetFractionInch().Numerator);
	}
}
