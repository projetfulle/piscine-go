package piscine


func UltimateDivMod(a *int, b *int){
	div:= int(*a / *b)
	mod:= *a%*b
	*a=div
	*b=mod
}
