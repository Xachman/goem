package EnglishMeasurements


type EnglishMeasurement struct {
	fraction Fraction;
}

const cFOOT  = 12
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
func NewEnglishMeasurement(foot, inch int, frac *Fraction) EnglishMeasurement {
	inches := (foot * cFOOT) + inch;
	newFrac := Fraction{
		Numerator: (inches * frac.Denominator + frac.Numerator),
		Denominator: frac.Denominator,
	}
	newEm := EnglishMeasurement{newFrac}
	return newEm
}

func (e *EnglishMeasurement) GetInch() int  {
	return e.fraction.GetInteger() % cFOOT
}

func (e *EnglishMeasurement) GetFoot() int {
	return e.fraction.GetInteger() / cFOOT
}


func (e *EnglishMeasurement) GetFractionInch() Fraction {
	return Fraction{
		Numerator: e.fraction.Numerator % e.fraction.Denominator ,
		Denominator: e.fraction.Denominator,
	}
}