package engmes

import (
	"regexp"
	"strings"
	"strconv"
	"fmt"
	"errors"
)
// Parse to parse string
type Parse struct {
	Input string
}

// FromString parse from string
func ParseString(st string) (EnglishMeasurement, error) {

	foot, footErr := parseFoot(st)
	inch, inchErr := parseInch(st)
	fraction, fracErr := parseFraction(st);
	if(footErr != nil && inchErr != nil && fracErr != nil) {
		return NewEnglishMeasurement(foot, inch, &fraction), errors.New("Bad Value");
	}
	return NewEnglishMeasurement(foot, inch, &fraction), nil
}

func parseFoot(st string) (int, error) {
	footReg := regexp.MustCompile("[0-9]+'");

	if footReg.MatchString(st) {
		replace := strings.Replace(footReg.FindString(st), "'", "", 1);
		i, err := strconv.Atoi(replace)
		if err != nil {
			fmt.Println(err)
		}else {
			return i, nil
		}
	}

	return 0, errors.New("No Value")
}

func parseInch(st string) (int, error) {
	reg := regexp.MustCompile("[0-9]+\"");

	if reg.MatchString(st) {
		replace := strings.Replace(reg.FindString(st), "\"", "", 1);
		i, err := strconv.Atoi(strings.TrimSpace(replace))
		if err != nil {
			fmt.Println(err)
		}else {
			return i, nil
		}
	}

	return 0, errors.New("No Value")
}

func parseFraction(st string) (Fraction, error) {
	reg := regexp.MustCompile("[0-9]+/[0-9]+");

	if reg.MatchString(st) {
		split := strings.Split( reg.FindString(st), "/")
		i, err := strconv.Atoi(strings.TrimSpace(split[0]))
		if err != nil {
			fmt.Println(err)
		}
		i2, err2 := strconv.Atoi(strings.TrimSpace(split[1]))
		if err2 != nil {
			fmt.Println(err)
		}

		return Fraction{i,i2}, nil
	}

	return Fraction{0,1}, errors.New("No Value")
}
