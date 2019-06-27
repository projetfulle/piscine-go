package piscine

func ListSize(l *List) int {
	if l.Head!=nil{
		count:=1
		for l.Head.Next!=nil{
			l.Head=l.Head.Next
			count++
		}
		return count
	}	
	return 0
}


