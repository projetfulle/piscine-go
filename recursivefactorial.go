package piscine

import "math"

func RecursiveFactorial(nb int) int{
	if nb == 0 || nb==1{
		return 1
	}
	if nb > 1{
		if nb > math.MaxInt32{
			return 0			
		}
		return RecursiveFactorial(nb-1)*nb
	}
	return 0
}


