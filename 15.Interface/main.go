package main

import (
	"fmt"
	"os"
	"strings"
)

type User struct{
	Name        string
	Age         int
	Occupation  string
}

func main(){
	file,err:=os.ReadFile("sample.txt")
	if err!=nil{
		fmt.Println("File o'qishda xatolik?")
		return
	}

	var slice []string
	for _,v:=range strings.Split(string(file),"\n"){
		slice=append(slice,v)
	}

	var users []User
	var user User
	for _,v:=range slice{
		if v==""{
			continue
		}

		s:=strings.Split(v,":")

		k:=strings.TrimSpace(s[0])
		v:=strings.TrimSpace(s[1])

		if k=="Name"{
			user.Name=v
		}else if k=="Age"{
			fmt.Sscanf(v,"%d",&user.Age)
		}else if k=="Occupation"{
			user.Occupation=v
		}
		// if user.Name!="" && user.Age!=0 && user.Occupation!=""{
		// 	users=append(users, user)
		// }
		users=append(users, user)
	}

	for _,v:=range users{
		if v.Name!="" && v.Age!=0 && v.Occupation!=""{
			fmt.Println("Name: ",v.Name)
			fmt.Println("Age: ",v.Age)
			fmt.Println("Occupation: ",v.Occupation)
			fmt.Println()
		}
	}
}