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

	testEngMes(12, 10, Fraction{3,4}, em3, t)

}

func testEngMes(foot, inch int, frac Fraction, em *EnglishMeasurement, t *testing.T) {
	if foot != em.GetFoot() {
		t.Error(fmt.Sprintf("Expected Foot %d but got ", foot), em.GetFoot());
	}

	if inch != em.GetInch() {
		t.Error(fmt.Sprintf("Expected Inch %d but got ", inch), em.GetInch());
	}

	if frac.Denominator != em.Fraction.Denominator {
		t.Error(fmt.Sprintf("Expected Denominator %d but got ", frac.Denominator), em.Fraction.Denominator);
	}

	if frac.Numerator != em.Fraction.Numerator {
		t.Error(fmt.Sprintf("Expected Denominator %d but got ", frac.Numerator), em.Fraction.Numerator);
	}
}
