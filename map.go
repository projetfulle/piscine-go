package piscine

func Map(f func(int) bool, arr []int) []bool {
	var result []bool
	for _,i:=range arr{
		result=append(result,f(i))
	}
	return result
}
