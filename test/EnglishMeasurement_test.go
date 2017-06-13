package Test

import "testing"
import (
	. "github.com/xachman/goem/engmes"
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

	testEngMes(3, 9, Fraction{3,4}, em3, t)

}


func TestArea(t *testing.T) {
	em := NewEnglishMeasurement(2, 0, &Fraction{
		Numerator: 0,
		Denominator: 1,
	})

	em2 := NewEnglishMeasurement(2, 0, &Fraction{
		Numerator: 0,
		Denominator: 1,
	})

	em3 := em2.Area(em)

	testValues(4, 0, Fraction{0,1}, em3.GetFootArea(), em3.GetInchArea(), em3.GetFractionInch(), t)

}
func TestAreaInches(t *testing.T) {
	em := NewEnglishMeasurement(2, 6, &Fraction{
		Numerator: 0,
		Denominator: 1,
	})

	em2 := NewEnglishMeasurement(2, 6, &Fraction{
		Numerator: 0,
		Denominator: 1,
	})

	em3 := em2.Area(em)
	testValues(6, 36, Fraction{0,1}, em3.GetFootArea(), em3.GetInchArea(), em3.GetFractionInch(), t)

}

func TestDivide(t *testing.T) {

	em2 := NewEnglishMeasurement(2, 0, &Fraction{
		Numerator: 0,
		Denominator: 1,
	})

	em := NewEnglishMeasurement(12, 0, &Fraction{
		Numerator: 0,
		Denominator: 1,
	})

	em3 := em2.Divide(em)

	testEngMes(6, 0, Fraction{0,1}, em3, t)
}

func TestAreaInchesFrac(t *testing.T) {
	em := NewEnglishMeasurement(2, 6, &Fraction{
		Numerator: 1,
		Denominator: 2,
	})

	em2 := NewEnglishMeasurement(2, 6, &Fraction{
		Numerator: 1,
		Denominator: 4,
	})

	em3 := em2.Area(em)


	testValues(6, 58, Fraction{5,8}, em3.GetFootArea(), em3.GetInchArea(), em3.GetFractionInch(), t)

}
func TestToString(t *testing.T) {
	em := NewEnglishMeasurement(2, 6, &Fraction{
		Numerator: 1,
		Denominator: 2,
	})

	test := "2' 6\" 1/2"

	if test != em.ToString() {
		t.Error(fmt.Sprintf("Expected %s but got %s", test, em.ToString()))
	}
}
func TestToStringNoFrac(t *testing.T) {
	em := NewEnglishMeasurement(2, 6, &Fraction{
		Numerator: 0,
		Denominator: 1,
	})

	test := "2' 6\""

	if test != em.ToString() {
		t.Error(fmt.Sprintf("Expected %s but got %s", test, em.ToString()))
	}
}
func TestToStringArea(t *testing.T) {
	em := NewEnglishMeasurement(0, 0, &Fraction{
		Numerator: 7381,
		Denominator: 8,
	})

	test := "6' 58\" 5/8"

	if test != em.ToStringArea() {
		t.Error(fmt.Sprintf("Expected %s but got %s", test, em.ToStringArea()))
	}
}
func TestToStringAreaNoFrac(t *testing.T) {
	em := NewEnglishMeasurement(0, 0, &Fraction{
		Numerator: 900,
		Denominator: 1,
	})

	test := "6' 36\""

	if test != em.ToStringArea() {
		t.Error(fmt.Sprintf("Expected %s but got %s", test, em.ToStringArea()))
	}
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

func testValues(foot, inch int, frac Fraction, testFoot, testInch int, testFrac Fraction, t *testing.T) {
	if foot != testFoot {
		t.Error(fmt.Sprintf("Expected Foot %d but got ", foot), testFoot);
	}

	if inch != testInch {
		t.Error(fmt.Sprintf("Expected Inch %d but got ", inch), testInch);
	}

	if frac.Denominator != testFrac.Denominator {
		t.Error(fmt.Sprintf("Expected Denominator %d but got ", frac.Denominator), testFrac.Denominator);
	}

	if frac.Numerator != testFrac.Numerator {
		t.Error(fmt.Sprintf("Expected Numerator %d but got ", frac.Numerator), testFrac.Numerator);
	}
}
