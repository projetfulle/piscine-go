package piscine


func StrLen(str string) int{
	nb:=0
	for _,n := range str {
		nb++
		n++
	}
	return nb
}
