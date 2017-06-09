package Test

import "testing"
import (
	. "github.com/xachman/goem/engmes"
	"fmt"
)

func TestParse(t *testing.T) {

	teststring := "2' 4\" 1/2"

	em := ParseString(teststring)

	compair(NewEnglishMeasurement(2,4,&Fraction{1,2}), em, t);

}

func TestParseNoSpace(t *testing.T) {

	teststring := "2'4\"1/2"

	em := ParseString(teststring)

	compair(NewEnglishMeasurement(2,4,&Fraction{1,2}), em, t);

}
func compair(em, testEM EnglishMeasurement, t *testing.T) {

	if em.GetInch() != testEM.GetInch() {
		t.Error(fmt.Sprintf("Inch should be %d and its %d ", em.GetInch(), testEM.GetInch()))
	}

	if em.GetFoot() != testEM.GetFoot() {
		t.Error(fmt.Sprintf("Foot should be %d and its %d", em.GetFoot(), testEM.GetFoot()))
	}

	if em.GetFractionInch().Numerator != testEM.GetFractionInch().Numerator {
		t.Error(fmt.Sprintf("Numerator should be %d and its %d", em.GetFractionInch().Numerator, testEM.GetFractionInch().Numerator))
	}
	if em.GetFractionInch().Denominator != testEM.GetFractionInch().Denominator {
		t.Error(fmt.Sprintf("Numerator should be %d and its %d", em.GetFractionInch().Denominator, testEM.GetFractionInch().Denominator))
	}
}
