package piscine

import "fmt"

func PrintCombN(n int){
	nb:=firstnumber(n)
	last:=lastnumber(n)
	for !same(nb,last){
		printnb(nb,last)
		increment(nb, len(nb)-1, last, true)
	}
	printnb(last,last)
	fmt.Print("\n")
}

func firstnumber(n int) []int{
	nb:=make([]int,n)
	for i:=range nb{
		nb[i]=i
	}
	return nb
}

func lastnumber(n int) []int{
	nb:=make([]int,n)
	j:=10-n
	for i:=range nb{
		nb[i]=j
		j++
	}
	return nb
}

func printnb(nb []int, max []int){
	if same(nb,max){
		for i:= range nb{
			fmt.Printf("%d",nb[i])		
		}
	}else{
		for i:= range nb{
			fmt.Printf("%d",nb[i])		
		}
		fmt.Print(", ")
	}
}

func same(a, b []int) bool {
        if len(a) != len(b) {
                return false
        }
        for i, v := range a {
                if v != b[i] {
                        return false
                }
        }
        return true
}

func increment(nb []int, index int, lastnumber []int, do bool) {	
	if nb[index]!=lastnumber[index]{
		nb[index]++
	}else{
		if index!=0{
			if nb[index-1]!=lastnumber[index-1]{
				if do{
					nb[index-1]++
				}
				nb[index]=nb[index-1]+1		
			}else{	
				increment(nb,index-1,lastnumber,true)
			}
		}
		if index<len(nb)-1{
			if nb[index+1]!=nb[index]+1{	
				increment(nb,index+1,lastnumber,false)
			}		
		}
	}
}



