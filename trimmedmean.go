package trimmedmean

import (
    "sort"
)

// TrimmedMean computes the trimmed mean of a slice of float64 numbers.
func TrimmedMean(data []float64, lowTrim float64, highTrim float64) float64 {
    if highTrim == -1 {
        highTrim = lowTrim
    }

    if len(data) == 0 {
        return 0
    }

    sorted := make([]float64, len(data))
    copy(sorted, data)
    sort.Float64s(sorted)

    n := len(sorted)
    lowIndex := int(float64(n) * lowTrim)
    highIndex := n - int(float64(n)*highTrim)

    if lowIndex >= highIndex {
        return 0 // Not enough data after trimming
    }

    trimmedData := sorted[lowIndex:highIndex]

    sum := 0.0
    for _, v := range trimmedData {
        sum += v
    }

    return sum / float64(len(trimmedData))
}
