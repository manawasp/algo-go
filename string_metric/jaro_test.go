package sdist

import (
	"fmt"
	"math"
	"testing"
)

func TestJaroCase1(t *testing.T) { testJaro(t, "MARTHA", "MARHTA", 0.944444) }
func TestJaroCase2(t *testing.T) { testJaro(t, "DIXON", "DICKSONX", 0.766667) }
func TestJaroCase3(t *testing.T) { testJaro(t, "A", "A", 1) }
func TestJaroCase4(t *testing.T) { testJaro(t, "C", "E", 0) }
func TestJaroCase5(t *testing.T) { testJaro(t, "ABCDEF", "123456", 0) }
func TestJaroCase6(t *testing.T) { testJaro(t, "AAAAAABCCCC", "AAAAAABCCCC", 1) }
func TestJaroCase7(t *testing.T) { testJaro(t, "AL", "LA", 0) }

/// PRIVATE

func testJaro(t *testing.T, str1, str2 string, res float64) {
	result := Jaro(str1, str2)
	pow := 100000.0
	if math.Ceil(result*pow) != math.Ceil(res*pow) {
		fmt.Printf("Jaro mismatch resultat : %f - %f.", res, result)
		t.Fail()
	}
}
