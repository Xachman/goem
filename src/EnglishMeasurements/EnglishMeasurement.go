package EnglishMeasurements

type EnglishMeasurement struct {
	Fraction Fraction;
}

const cFOOT  = 12
func (e *EnglishMeasurement) Add(em *EnglishMeasurement) (*EnglishMeasurement) {
	newFrac := em.Fraction.Add(&e.Fraction);
	newEm := EnglishMeasurement{Fraction{newFrac.Numerator, newFrac.Denominator}}
	return &newEm
}

func NewEnglishMeasurement(foot, inch int, frac *Fraction) (*EnglishMeasurement) {
	inches := (foot * 12) + inch;
	newFrac := Fraction{
		Numerator: (inches * frac.Denominator),
		Denominator: frac.Denominator,
	}
	newEm := EnglishMeasurement{newFrac}
	return &newEm
}

func (e *EnglishMeasurement) GetInch() int  {
	return e.Fraction.GetInteger()
}

func (e *EnglishMeasurement) GetFoot() int {
	return e.GetInch() / cFOOT
}