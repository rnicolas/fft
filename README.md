# fft

Fast Fourier Transform in Golang.

A library of fast Fourier transforms / inverse transforms.

## Install

    $ go get github.com/takatoh/fft

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

## License

MIT License
