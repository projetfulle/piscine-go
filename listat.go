package piscine

func ListAt(l *NodeL, pos int) *NodeL{
	if pos<0{
		return nil
	}
	count:=0
	if l!=nil{
		for l.Next!=nil && count<pos{
			l=l.Next
			count++
		}
		if pos-count>0{
			return nil
		}
		return l
	}
	return nil
}

