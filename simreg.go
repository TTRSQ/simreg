package simreg

import "errors"

type result struct {
	Slope     float64 // a of y = ax + b
	Intercept float64 // b of y = ax + b
	COD       float64 // coefficient of determination
}

func (r *result) Predict(x float64) float64 {
	return x*r.Slope + r.Intercept
}

// Result .. O(N)
func Result(xList, yList []float64) (*result, error) {
	if len(xList) != len(yList) {
		return nil, errors.New("The lengths of xList and yList are different")
	}
	if len(xList) < 2 {
		return nil, errors.New("List length must be 2 or more")
	}
	length := float64(len(xList))
	meanX := 0.0
	meanY := 0.0
	for i := range xList {
		meanX += xList[i] / length
		meanY += yList[i] / length
	}
	xxS := 0.0
	xyS := 0.0
	yyS := 0.0
	for i := range xList {
		xxS += (xList[i] - meanX) * (xList[i] - meanX)
		xyS += (xList[i] - meanX) * (yList[i] - meanY)
		yyS += (yList[i] - meanY) * (yList[i] - meanY)
	}
	if xxS == 0.0 || yyS == 0.0 {
		return nil, errors.New("Sxx or Syy is zero")
	}
	slope := xyS / xxS
	intercept := meanY - slope*meanX

	r := result{
		Slope:     slope,
		Intercept: intercept,
	}
	ppS := 0.0
	for i := range xList {
		ppS += (r.Predict(xList[i]) - yList[i]) * (r.Predict(xList[i]) - yList[i])
	}
	r.COD = 1 - ppS/yyS
	return &r, nil
}
