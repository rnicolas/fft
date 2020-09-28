package main

import (
	"fmt"
	"math"

	"github.com/rnicolas/fft"
)

func main() {
	x0 := []float64{
		31,
		31,
		26,
		26,
		24,
		24,
		21,
		21,
		16,
		15,
		20,
		20,
		14,
		14,
		14,
		14,
		16,
		16,
		17,
		17,
		13,
		13,
		15,
		15,
		12,
		12,
		11,
		11,
		13,
		13,
		12,
		12,
		8,
		8,
		12,
		11,
		18, 18, 13, 13, 13, 13, 16, 16, 11, 11, 10, 9, 16, 16, 14, 14, 8, 8, 16, 16, 15, 15, 10, 10, 16, 15, 11, 11, 11, 11, 14, 14, 8, 8, 7, 7, 10, 10, 4, 4, 8, 8, 6, 6, 4, 4, 9, 9, 5, 5, 7, 7, 11, 11, 3, 4, 4, 4, 6, 6, 3, 3, 6, 6, 6, 6, 4, 4, 6, 6, 4, 4, 3, 3, 5, 5, 4, 4, 4, 4, 5, 5, 3, 3, 9, 9, 10, 10, 6, 6, 12, 12, 10, 10, 8, 8, 13, 13, 4, 4, 9, 9, 6, 7, 1, 1, 5, 5, -2, -2, -3, -3, 3, 3, 1, 1, 4, 4, 6, 6, 2, 2, 3, 3, 5, 5, 8, 9, 7, 7, 13, 13, 13, 13, 15, 15, 7, 7, 8, 7, 12, 12, 10, 10, 14, 14, 12, 12, 7, 8, 9, 9, 4, 4, 5, 5, 7, 7, 3, 3, 3, 3, -2, -2, -3, -3, -3, -3, -1, -1, -3, -3,
		-4,
		-4,
		-1,
		-1,
		-4,
		-4,
		-5,
		-5,
		-7,
		-7,
		-9,
		-9,
		-5,
		-6,
		-4,
		-4,
		-7,
		-7,
		-4,
		-4,
		0,
		0,
		-2,
		-2,
		-3,
		-4,
		1,
		1,
		-1,
		-1,
		-4,
		-4,
		1, 1, 6, 6, 5, 5, 14, 13,
		13, 13, 10, 10, 17, 17, 10, 10, 10, 10, 15, 15, 22, 22, 16, 16, 19, 19, 18, 18, 17, 17, 20, 20, 16, 16, 18, 18, 16, 16, 12, 12, 18, 18, 13, 13, 12, 12, 16, 16, 15, 15, 15, 15, 14, 14, 11, 11, 13, 13, 12, 12, 10, 10, 10, 10, 17, 17, 16, 16, 16, 16, 17, 17, 15, 15, 16, 16, 20, 20, 17, 17, 17, 17, 19, 19, 16, 16, 14, 14, 12, 12, 19, 19, 18, 18, 15, 15, 12, 12, 13, 14, 9, 9, 10, 10, 7, 7, 8, 8, 8, 8, 5, 5, 2, 2, 5, 5, 5, 5, 4, 4, 5, 5, 6, 6, 7, 7, 2, 2, 5, 5, 1, 1, 3, 3, 7, 7, 6, 6, 7, 7, 11, 11, 10, 10,
		6, 6, 9, 9, 10, 10, 13, 13, 13, 13, 10, 10, 11, 11, 7, 7, 10, 10, 10, 10, 9, 8, 10, 10, 7, 7, 9, 9, 5, 6, 6, 6, 6, 6, 4, 4, 6, 6, 4, 4, 6, 6, 3, 3, 2, 2, 8, 8, 8, 8, 7, 7, 17, 17, 12, 12, 9, 9, 17, 16, 16, 16, 17, 17, 20, 20, 15, 15, 19, 19, 15, 15, 17, 16, 20, 20, 18, 18, 22, 22, 17, 17, 17, 17, 19, 19, 15, 15, 19, 19, 17, 17, 23, 23, 19, 19, 20, 20, 17, 17, 12, 12, 16, 16, 11, 11, 16, 16, 14, 14, 4, 4, 9, 9, 9, 9, 5, 5, 13, 12, 9, 9, 2, 2, 10, 10, 7, 8, 0, 0, 1, 2, -1, -1, 1, 1, 5, 5, 8, 8, 5, 5, 4, 4, 6, 6, 10, 10, 7, 7, 4, 4, 6, 6, 6, 6, 5, 5, 8, 8, 7, 7, 6, 6, 6, 6, 11, 11, 11, 11, 10, 10, 8, 8, 6, 6, 8, 8, 4, 4, 3, 3, 7, 7, 7, 8, 7, 6, 10, 10, 13, 13, 10, 10, 11, 11, 15, 15, 13, 13, 20, 20, 20, 21, 20, 20, 27, 27, 26, 26, 30, 30, 31, 32, 28, 27, 33, 33, 32, 32, 35, 35, 37, 38, 35, 35, 40, 40, 32, 32, 34, 34, 35, 35, 27, 27, 33, 33, 28, 29, 22, 22, 25, 25, 20, 20, 14, 14, 18, 18, 10, 10, 8, 8, 5, 5, 2, 2, 6, 6, 3, 3, -5, -5, -4, -4, -4, -4, -3, -3, 6, 5, 2, 2, 2, 2, 8, 8, 3, 3, 1, 1, 4, 4, 6, 6, 7, 7, 12, 12, 16, 16, 13, 13, 14, 14, 16, 16, 18, 18, 18, 18, 14, 14, 18, 18, 16, 16, 15, 15, 15, 15, 13, 13, 11, 11, 10, 10, 12, 12, 13, 13, 15, 15, 8, 8, 6, 5, 8, 8, 2, 2, 3, 3, 6, 6, 3, 3, -1, -1, 1, 1, 0, 0, 0, 0, -5, -4, -7, -7, -3, -3, -2, -2, -2, -2, -3, -3, -7, -7, -5, -5, -4, -4, -4, -4, 3, 3, -2, -1, -5, -5, 0, 0, 1, 1, 1, 1, 1, 1, 4, 4, 3, 3, 5, 5, 6, 6, 12, 12, 10, 10, 12, 12, 16, 16, 10, 10, 14, 14, 17, 17, 13, 13, 20, 20, 20, 20, 15, 15, 19, 19, 15, 15, 18, 18, 19, 19, 19, 19, 16, 16, 15, 15, 13, 13, 19, 19, 17, 17, 9, 9, 9, 9, 7, 7, 8, 8, 3, 3, 1, 1, 8, 8, 4, 4, 3, 3, 6, 6, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 1, 1, -4, -4, -4, -4, -4, -4, -8, -8, -3, -3, -4, -4, -9, -9, -6, -6, -9, -9, -10, -10, -6, -6, -12, -11, -10, -10, -8, -8, -12, -12, -9, -9, -12, -12, -7, -7, -4, -4, -10, -10, -5, -5, -5, -5, -5, -5, -6, -6, -2, -2, -1, -1, 4, 4, 3, 3, 3, 3, 4, 4, 6, 6, 8, 8, 7, 7, 10, 10, 7, 7, 8, 8, 12, 12, 14, 14, 11, 11, 12, 12, 14, 14, 8, 8, 15, 14, 21, 21, 12, 12, 17, 17, 15, 15, 15, 15, 19, 19, 16, 16, 18, 18, 24, 24, 18, 18, 24, 24, 22, 23, 16, 16, 16, 16, 19, 19, 19, 19, 21, 21, 20, 20, 15, 15, 17, 17, 18, 18, 13, 13, 12, 11, 16, 16, 12, 12, 7, 7, 9, 9, 7, 7, 7, 7, 5, 5, 8, 8, 5, 5, 1, 1, 0, 0, 2, 2, -4, -4, 0, 0, 1, 1, -4, -4, -1, -1, -2, -2, -6, -6, -5, -5, -1, -1, 0, 0, -1, -1,
	}
	n := len(x0)
	x := make([]complex128, n)
	for k := 0; k < n; k++ {
		x[k] = complex(x0[k], 0.0)
	}

	y := fft.FFT(x, n)
	z := fft.IFFT(y, n)
	fmt.Println(" K   DATA  FOURIER TRANSFORM  INVERSE TRANSFORM  MAGNITUDE")
	for k := 0; k < n; k++ {
		fmt.Printf("%4d %6.1f  %8.3f%8.3f   %8.3f%8.3f   %8.3f\n",
			k, x0[k], real(y[k]), imag(y[k]), real(z[k]), imag(z[k]), math.Sqrt(math.Pow(real(y[k]), 2)+math.Pow(imag(y[k]), 2)))
	}
}
