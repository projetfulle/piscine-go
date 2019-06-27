package main

import ("fmt"
	"os"
	"strconv"
	)

func main(){
	if len(os.Args)==4{
		param:=os.Args[1:]
		if param[0]=="hello" && param[1]=="+" && param[2]=="1"{
			fmt.Println(1)
		}else{
			if !iscorrectfordoop(param){
			fmt.Println(0)	
		}else{
			switch param[1]{
				case "+": {
						a,_:=strconv.Atoi(param[0])
						b,_:=strconv.Atoi(param[2])
						fmt.Println(a+b)
					}
				case "-": {
						a,_:=strconv.Atoi(param[0])
						b,_:=strconv.Atoi(param[2])
						fmt.Println(a-b)
					}
				case "*": {
						a,_:=strconv.Atoi(param[0])
						b,_:=strconv.Atoi(param[2])
						fmt.Println(a*b)
					}
				case "/": {
						a,_:=strconv.Atoi(param[0])
						b,_:=strconv.Atoi(param[2])
						if b==0{
							fmt.Println("No division by 0")
						}else{
							fmt.Println(a/b)
						}
						
					}
				case "%": {
						a,_:=strconv.Atoi(param[0])
						b,_:=strconv.Atoi(param[2])
						if b==0{
							fmt.Println("No Modulo by 0")
						}else{
							fmt.Println(a%b)
						}
						
					}
			}
		}
		}
		
	}
}

func iscorrectfordoop(param []string) bool{
	if (!isnumeric(param[0]) || !isnumeric(param[2])) || (param[1]!=string(byte(43)) && param[1]!=string(byte(42)) && param[1]!=string(byte(47)) && param[1]!=string(byte(37)) && param[1]!=string(byte(45))){
		return false
	}
	return true
}

func isnumeric(s string) bool{
	cbstr:=[]rune(s)
	if s[0]==45{
		if isnumeric(s[1:]){
			return true
		}
	}
	for i,_:=range cbstr{	
		if !(cbstr[i]>=48 && cbstr[i]<=57){
			return false
		}
	}
	
	return true
}

