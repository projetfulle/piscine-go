package piscine

import "math"

func IterativePower(nb,power int) int{
	if power < 0{
		return 0
	}
	if power >=0{
		result:=1
		for i:=1;i<=power;i++{
			result*=nb
			if result > math.MaxInt32{
				return 0			
			}
		}
		return result
	}
	return 0
}


