package piscine

import "fmt"

func Raid1b(x,y int) {
    for i:=0; i < y; i++ {//boucle pour la longueur
        for j:=0; j < x; j++ {//boucle pour la largeur
            if i == 0 {//première ligne
                if j == 0 || j == x - 1 {
			if j == 0{//première cologne
				fmt.Print("/")
			}
			if j == x-1{//dernière cologne
				if j!=0{
					fmt.Print("\\")
				}
				
			}                    
                }else{//colognes du milieu
                    fmt.Print("*")
                }
            }else if i == y - 1{//dernière ligne
		if j == 0 || j == x - 1 {
			if j == 0{//première cologne
				fmt.Print("\\")
			}
			if j == x-1{//dernière cologne
				if j!=0{
					fmt.Print("/")
				}
			}                    
                }else{//colognes du milieu
                    fmt.Print("*")
                }
            }else{
                if j == 0 || j == x - 1 {
                    fmt.Print("*")
                }else{
                    fmt.Print(" ")
                }
            }
            if j == x - 1 {
                fmt.Print("\n")        
            }
        }
    }
}
