package engmes

import (
	"regexp"
	"strings"
	"strconv"
	"fmt"
)
// Parse to parse string
type Parse struct {
	Input string
}

// FromString parse from string
func ParseString(st string) EnglishMeasurement {

	foot := parseFoot(st)
	inch := parseInch(st)
	fraction := parseFraction(st);
	return NewEnglishMeasurement(foot, inch, &fraction)
}

func parseFoot(st string) (int) {
	footReg := regexp.MustCompile("[0-9]+'");

	if footReg.MatchString(st) {
		replace := strings.Replace(footReg.FindString(st), "'", "", 1);
		i, err := strconv.Atoi(replace)
		if err != nil {
			fmt.Println(err)
		}else {
			return i
		}
	}

	return 0
}

func parseInch(st string) (int) {
	reg := regexp.MustCompile("[0-9]+\"");

	if reg.MatchString(st) {
		replace := strings.Replace(reg.FindString(st), "\"", "", 1);
		i, err := strconv.Atoi(strings.TrimSpace(replace))
		if err != nil {
			fmt.Println(err)
		}else {
			return i
		}
	}

	return 0
}

func parseFraction(st string) (Fraction) {
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

		return Fraction{i,i2}
	}

	return Fraction{0,1}
}
