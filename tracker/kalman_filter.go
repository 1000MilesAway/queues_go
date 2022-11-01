package tracker

import (
	"gonum.org/v1/gonum/mat"
	// "gonum.org/v1/gonum/matrix"
)

const ndim, dt = 4, 1.
var Motion_mat = DiagOnes(ndim)

func DiagOnes(ndim int) mat.Dense {
	data := make([]float64, ndim*ndim)
	for i := 0; i < ndim; i++ {
		for j := 0; j < ndim; j++ {
			if i == j {
				data[i*j] = 1.0
			}
		}
	}
	return *mat.NewDense(ndim, ndim, data)
} 


type KalmanFilter struct {
	Score float64 
}
