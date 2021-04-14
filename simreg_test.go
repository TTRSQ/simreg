package simreg

import (
	"math"
	"testing"
)

func checkHasError(val, expect float64) bool {
	allowableError := math.Pow(0.1, 10)
	if math.Abs(val-expect) < allowableError {
		return false
	}
	return true
}

func TestResult(t *testing.T) {
	res, err := Result(
		[]float64{-1, 1},
		[]float64{1, -1},
	)
	if err != nil {
		t.Error(err.Error())
	}

	t.Logf("result:%+v", res)

	slope := -1.0
	intercept := 0.0
	cod := 1.0

	hasError := checkHasError(res.Slope, slope)
	if hasError {
		t.Error("invalid res.Slope")
	}
	hasError = checkHasError(res.Intercept, intercept)
	if hasError {
		t.Error("invalid res.Intercept")
	}
	hasError = checkHasError(res.COD, cod)
	if hasError {
		t.Error("invalid res.COD")
	}
}
