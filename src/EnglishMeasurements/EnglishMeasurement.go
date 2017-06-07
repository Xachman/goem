package EnglishMeasurements

import "fmt"

type EnglishMeasurement struct {
	fraction Fraction;
}

const cFOOT  = 12
func NewEnglishMeasurement(foot, inch int, frac *Fraction) EnglishMeasurement {
	inches := (foot * cFOOT) + inch;
	newFrac := Fraction{
		Numerator: (inches * frac.Denominator + frac.Numerator),
		Denominator: frac.Denominator,
	}
	newEm := EnglishMeasurement{newFrac}
	return newEm
}
func (e *EnglishMeasurement) Add(em EnglishMeasurement) EnglishMeasurement {
	newFrac := em.fraction.Add(&e.fraction)
	newEm := EnglishMeasurement{Fraction{newFrac.Numerator, newFrac.Denominator}}
	return newEm
}

func (e *EnglishMeasurement) Subtract(em EnglishMeasurement) EnglishMeasurement {
	newFrac := em.fraction.Subtract(&e.fraction)
	newEm := EnglishMeasurement{Fraction{newFrac.Numerator, newFrac.Denominator}}
	return newEm
}

func (e *EnglishMeasurement) Area(em EnglishMeasurement) EnglishMeasurement {
	newFrac := em.fraction.Multiply(&e.fraction)
	return EnglishMeasurement{
		Fraction{
			Numerator: newFrac.Numerator,
			Denominator: newFrac.Denominator,
		}}
}

func (e *EnglishMeasurement) Divide(em EnglishMeasurement) EnglishMeasurement {
	newFrac := em.fraction.Divide(&e.fraction)

	fmt.Println(e.fraction)
	fmt.Println(em.fraction)
	fmt.Println(newFrac)

	return EnglishMeasurement{
		Fraction{
			Numerator: newFrac.Numerator,
			Denominator: newFrac.Denominator,
		}}
}

func (e *EnglishMeasurement) GetInch() int  {
	return e.fraction.GetInteger() % cFOOT
}

func (e *EnglishMeasurement) GetInchArea() int {
	return e.fraction.GetInteger() % (cFOOT * cFOOT)
}

func (e *EnglishMeasurement) GetFoot() int {
	return e.fraction.GetInteger() / cFOOT
}

func (e *EnglishMeasurement) GetFootArea() int  {
	return e.fraction.GetInteger() / (cFOOT * cFOOT)
}


func (e *EnglishMeasurement) GetFractionInch() Fraction {
	return Fraction{
		Numerator: e.fraction.Numerator % e.fraction.Denominator ,
		Denominator: e.fraction.Denominator,
	}
}