# fft

Fast Fourier Transform in Golang.

A library of fast Fourier transforms / inverse transforms.

## Install

    $ go get github.com/rnicolas/fft

## Usage

Use the `fft.FFT` function for the Fourier transform.

```go
y := fft.FFT(x, n)
```

`x` is the number of data (complex number) and` n` is the number of data. `n` must be a power of 2.

Use the `fft.IFFT` function for the inverse transform.

```go
z := fft.IFFT(y, n)
```

Use the `fft.GetAmplitude` function to get the amplitude of a Fourier transform.

```go
y := fft.FFT(x, n)
z := fft.GetAmplitude(y)
```

`x` is the number of data (complex number) and `n` is the number of data. `n` must be a power of 2. `y` is the FFT result. `z` is the array with amplitudes calculated.

## License

MIT License
