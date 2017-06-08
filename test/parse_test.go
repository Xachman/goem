package Test

import "testing"
import (
	. "github.com/xachman/goem/engmes"
	"fmt"
)

func TestParse(t *testing.T) {

	teststring := "2' 4\" 1/2"

	em := ParseString(teststring)

	if em.GetInch() != 4 {
		t.Error(fmt.Sprintf("Inch should be 4 and its %d ", em.GetInch()))
	}
}
