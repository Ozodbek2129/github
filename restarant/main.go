package main

import (
	"fmt"
	test "restarant/Admin"
	test2 "restarant/foydalanuvchi"
)

func main(){
	var x int
	var s int
	fmt.Printf("Admin -> 1\nFoydalanuvchi -> 2\n>>>: ")
	fmt.Scan(&s)
	fmt.Println()
	if s==1{
		var parol string
		fmt.Printf("Parolni kiritng: ")
		fmt.Scan(&parol)
		if parol=="12345"{
			fmt.Println()
			for {
				fmt.Printf("\n\nTaom kiritish -> 1\nBiror taomni o'chirish -> 2\nYakunlash -> 0\n>>>: ")
				fmt.Scan(&x)
				if x==0{
					return
				}else if x==1{
					test.Admin()
				}else if x==2{
					test.Remove()
				}
			}
		}
	}else if s==2{
		var f int 
		fmt.Printf("Menuni ko'rish -> 1\nYakunlash -> 0\n>>>: ")
		fmt.Scan(&f)
		fmt.Println()
		if f==1{
			for {
				var d int
				test2.Taomlar()
				fmt.Printf("\nZakas berish -> 1\nYakunlash -> 0\n>>>: ")
				fmt.Scan(&d)
				fmt.Println()
				if d==1{
					test2.Zakas()
				}else if d==0{
					return
				}
			}
		}
		if f==0{
			return
		}
	}
}