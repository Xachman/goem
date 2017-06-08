package engmes

type Fraction struct {
	Numerator int
	Denominator int
}

func (f *Fraction) Add(fraction *Fraction) *Fraction {
	commonDenominator := f.Denominator * fraction.Denominator;
	numeratorOne := fraction.Numerator * f.Denominator;
	numeratorTwo := f.Numerator * fraction.Denominator;
	addNumerator := numeratorOne + numeratorTwo;
	newFrac := Fraction{addNumerator, commonDenominator}
	return f.simplifyF(&newFrac);
}

func (f *Fraction) Divide(fraction *Fraction) *Fraction {
	numerator := f.Numerator * fraction.Denominator;
	denominator := fraction.Numerator * f.Denominator;
	return f.simplify(numerator, denominator);
}

func (f *Fraction) Multiply(fraction *Fraction) *Fraction {
	numerator := f.Numerator * fraction.Numerator;
	denominator := fraction.Denominator * f.Denominator;
	return f.simplify(numerator, denominator);
}


func (f *Fraction) Subtract(fraction *Fraction) *Fraction {
	commonDenominator := f.Denominator * fraction.Denominator;
	numeratorOne := fraction.Numerator * f.Denominator;
	numeratorTwo := f.Numerator * fraction.Denominator;
	subNumerator := numeratorOne - numeratorTwo;
	return f.simplify(subNumerator, commonDenominator);
}

func (f *Fraction) simplify(num, den int) *Fraction {
	greatestCommonDivisor := gcd(num, den);
	newFrac := Fraction{num / greatestCommonDivisor, den / greatestCommonDivisor};
	return &newFrac
}

func (f *Fraction) simplifyF(frac *Fraction) *Fraction {
	return f.simplify(frac.Numerator, frac.Denominator)
}

func (f *Fraction) GetInteger() int {
	return f.Numerator / f.Denominator
}

func gcd(a, b int) int {
	if(b == 0) {
		return a
	};
	return gcd(b, a%b);
}


