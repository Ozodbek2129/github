package main

import (
	"encoding/json"
	"fmt"
	"os"
)


type Employee struct{
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Age       int     `json:"age"`
	Position  string  `json:"postion"`
}

func main(){
	file,err:=os.ReadFile("employees.json")
	if err!=nil{
		fmt.Println("File ni o'qishda xatolik?")
		return
	}

	var employees []Employee
	er:=json.Unmarshal(file,&employees)
	if er!=nil{
		fmt.Println("Json malumotlarni umarshal qilshda xatolik?")
		return
	}

	for i:=range employees{
		if employees[i].Id==3{
			employees[i].Id=50
			break
		}
	}

	neweemployees:=Employee{
		Id:    6,
		Name:  "Alex",
		Age:   20,
		Position: "Project Manager",
	}

	employees=append(employees, neweemployees)

	new, err := json.MarshalIndent(employees, "", "  ")
	if err != nil {
		fmt.Println("Json ma'lumotlarini marshal qilishda xatolik?")
		return
	}

	errr:=os.WriteFile("employes_new.json",new,0644)
	if errr!=nil{
		fmt.Println("Yangi filega yozishda xatolik?")
		return
	}
}