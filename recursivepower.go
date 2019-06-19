package piscine

import "math"

func RecursivePower(nb,power int) int{
	if power < 0{
		return 0
	}
	if power==0{
		return 1
	}
	if power >0{
		if RecursivePower(nb,power-1)*nb  > math.MaxInt32{
			return 0			
		}
		return RecursivePower(nb,power-1)*nb 
	}
	return 0
}


