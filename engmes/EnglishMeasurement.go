package engmes


// EnglishMeasurement the struct
type EnglishMeasurement struct {
	fraction Fraction
}

const cFOOT = 12

// NewEnglishMeasurement Create EM constructor
func NewEnglishMeasurement(foot, inch int, frac *Fraction) EnglishMeasurement {
	inches := (foot * cFOOT) + inch
	newFrac := Fraction{
		Numerator:   (inches*frac.Denominator + frac.Numerator),
		Denominator: frac.Denominator,
	}
	newEm := EnglishMeasurement{newFrac}
	return newEm
}

// Add add 2 EMs
func (e *EnglishMeasurement) Add(em EnglishMeasurement) EnglishMeasurement {
	newFrac := em.fraction.Add(&e.fraction)
	newEm := EnglishMeasurement{Fraction{newFrac.Numerator, newFrac.Denominator}}
	return newEm
}

// Subtract 2 EMs
func (e *EnglishMeasurement) Subtract(em EnglishMeasurement) EnglishMeasurement {
	newFrac := em.fraction.Subtract(&e.fraction)
	newEm := EnglishMeasurement{Fraction{newFrac.Numerator, newFrac.Denominator}}
	return newEm
}

// Area of 2 EMs
func (e *EnglishMeasurement) Area(em EnglishMeasurement) EnglishMeasurement {
	newFrac := em.fraction.Multiply(&e.fraction)
	return EnglishMeasurement{
		Fraction{
			Numerator:   newFrac.Numerator,
			Denominator: newFrac.Denominator,
		}}
}

// Divide 2 EMs
func (e *EnglishMeasurement) Divide(em EnglishMeasurement) EnglishMeasurement {
	newFrac := em.fraction.Divide(&e.fraction)

	return EnglishMeasurement{
		Fraction{
			Numerator:   newFrac.Numerator,
			Denominator: newFrac.Denominator,
		}}
}

// GetInch of Fraction from EM
func (e *EnglishMeasurement) GetInch() int {
	return e.fraction.GetInteger() % cFOOT
}

// GetInchArea of Fraction from EM
func (e *EnglishMeasurement) GetInchArea() int {
	return e.fraction.GetInteger() % (cFOOT * cFOOT)
}

// GetFoot of Fraction from EM
func (e *EnglishMeasurement) GetFoot() int {
	return e.fraction.GetInteger() / cFOOT
}

// GetFootArea of Fraction from EM
func (e *EnglishMeasurement) GetFootArea() int {
	return e.fraction.GetInteger() / (cFOOT * cFOOT)
}

// GetFractionInch of Fraction from EM
func (e *EnglishMeasurement) GetFractionInch() Fraction {
	return Fraction{
		Numerator:   e.fraction.Numerator % e.fraction.Denominator,
		Denominator: e.fraction.Denominator,
	}
}
