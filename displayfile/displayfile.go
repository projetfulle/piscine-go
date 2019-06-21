package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        fmt.Println(e)
    }
}

func read(filename string) string {
    data, err := ioutil.ReadFile(filename)
    check(err)
    return string(data)
}

func main() {
	filename:=os.Args
	if len(filename)==1{
		fmt.Println("File name missing")
	}else if len(filename)>2{
		fmt.Println("Too many arguments")
	}else{
		data:=read(filename[1])
		fmt.Print(data)
	}
}

