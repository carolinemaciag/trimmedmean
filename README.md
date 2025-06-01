# trimmedmean

`trimmedmean` is a Go package for computing symmetric and asymmetric trimmed means from slices of float64 values. This program supports specifying separate proportions of data to trim from the lower and upper ends of the sorted dataset.

## Features

- Compute **symmetric trimmed mean** by specifying a single trim proportion.
- Compute **asymmetric trimmed mean** by specifying different proportions for low and high-end trimming.
- Includes unit tests for validation.

## Installation

To use `trimmedmean` in your Go project, run:

```bash
go get github.com/carolinemaciag/trimmedmean
```

## Usage

- Import the package and call the TrimmedMean function:

```go
import "github.com/carolinemaciag/trimmedmean"

data := []float64{1, 2, 3, 4, 5, 100}

// Symmetric trimming: 20% from both ends
resultSym := trimmedmean.TrimmedMean(data, 0.20, 0.20)

// Asymmetric trimming: 10% low-end, 20% high-end
resultAsym := trimmedmean.TrimmedMean(data, 0.10, 0.20)
```
