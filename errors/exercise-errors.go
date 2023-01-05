package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("%v is negative number. Sqrt returning with error.\n", float64(e))
}

func Sqrt(x float64) (float64, error) {
	fmt.Printf("\n--\nSqrt called with value: %v\n", x)
	if math.Signbit(x) {
		return 0, ErrNegativeSqrt(x)
	}
	z := float64(1)
	fmt.Printf("Sqrt approximation of %v:\n", x)
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %v, value = %v\n", i, z)
	}
	return z, nil
}
func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
