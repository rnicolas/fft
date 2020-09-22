package fft

import (
	"math"
	"math/cmplx"
)

//FFT calculates the Fast Fourier Transform
func FFT(x []complex128, n int) []complex128 {
	y := fft(x, n)
	for i := 0; i < n; i++ {
		y[i] = y[i] / complex(float64(n), 0.0)
	}
	return y
}

//IFFT calculates the Inverse Fast Fourier Transform
func IFFT(x []complex128, n int) []complex128 {
	y := make([]complex128, n)
	for i := 0; i < n; i++ {
		y[i] = cmplx.Conj(x[i])
	}
	y = fft(y, n)
	for i := 0; i < n; i++ {
		y[i] = cmplx.Conj(y[i])
	}
	return y
}

//GetAmplitude alculates frequencies amplitude given a FFT m
func GetAmplitude(m []complex128) []float64 {
	a := make([]float64, (len(m)))
	for index := 0; index < (len(m) / 2); index++ {
		a[index] = math.Sqrt(math.Pow(real(m[index]), 2) + math.Pow(imag(m[index]), 2))
	}
	return a
}

func fft(a []complex128, n int) []complex128 {
	x := make([]complex128, n)
	copy(x, a)

	j := 0
	for i := 0; i < n; i++ {
		if i < j {
			x[i], x[j] = x[j], x[i]
		}
		m := n / 2
		for {
			if j < m {
				break
			}
			j = j - m
			m = m / 2
			if m < 2 {
				break
			}
		}
		j = j + m
	}
	kmax := 1
	for {
		if kmax >= n {
			return x
		}
		istep := kmax * 2
		for k := 0; k < kmax; k++ {
			theta := complex(0.0, -1.0*math.Pi*float64(k)/float64(kmax))
			for i := k; i < n; i += istep {
				j := i + kmax
				temp := x[j] * cmplx.Exp(theta)
				x[j] = x[i] - temp
				x[i] = x[i] + temp
			}
		}
		kmax = istep
	}
}
