package engmes

// Parse to parse string
type Parse struct {
	Input string
}

// FromString parse from string
func ParseString(st string) EnglishMeasurement {
	return EnglishMeasurement{
		Fraction{
			Numerator:   1,
			Denominator: 2,
		}}
}
