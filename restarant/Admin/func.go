package test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type Taomlar_ruyxati struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func Taomlar(t Taomlar_ruyxati) {
	file, err := os.OpenFile("taomlar_ruyxat.json", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Xatolik")
		return
	}
	defer file.Close()

	var taomlar []Taomlar_ruyxati

	fileContent, err := os.ReadFile("taomlar_ruyxat.json")
	if err != nil{
		fmt.Println("Xatolik")
		return
	}

	if len(fileContent) > 0 {
		err = json.Unmarshal(fileContent, &taomlar)
		if err != nil {
			fmt.Println("Xatolik?")
			return
		}
	}

	taomlar = append(taomlar, t)

	fileData, err := json.MarshalIndent(taomlar, "", "    ")
	if err != nil {
		fmt.Println("Xatolik")
		return
	}

	err = os.WriteFile("taomlar_ruyxat.json", fileData, 0644)
	if err != nil {
		fmt.Println("Xatolik")
		return
	}

	fmt.Println("Malumotlar ro'yxatga muvaffaqiyatli qo'shildi.")
	fmt.Println()
}

func Admin() {

	var name string
	var price int
	fmt.Printf("Taom nomini kiriting: ")
	fmt.Scan(&name)
	fmt.Printf("Taom narxini kiriting: ")
	fmt.Scan(&price)

	c:=rand.Intn(400) + 1
	for {
		if c==0{
			c=rand.Intn(400) + 1
		}else{
			break
		}
	}
	taom := Taomlar_ruyxati{
		Id:    c,
		Name:  name,
		Price: price,
	}
	Taomlar(taom)
}

func Remove(){
	filename := "taomlar_ruyxat.json"
	var a int
	fmt.Printf("\n\nHaqiqatdan ham biror malumotni o'chirmoqchimisiz\nha -> 1\nyo'q -> 0\n>>>: ")
	fmt.Scan(&a)

	if a==1{
		file,err:=os.ReadFile(filename)
		if err != nil {
			fmt.Println("Xatolik:", err)
			return
		}

		var data []Taomlar_ruyxati
		errr:=json.Unmarshal(file,&data)
		if errr != nil {
			fmt.Println("Xatolik:", err)
			return
		}

		err = os.Remove(filename)
		if err != nil {
			fmt.Println("Faylni o'chirishda xatolik:", err)
			return
		}

		file1,err1:=os.OpenFile(filename,os.O_APPEND|os.O_CREATE,0644)
		if err1!=nil{
			fmt.Println("Xatolik?")
			return
		}
		defer file1.Close()
		var slice []Taomlar_ruyxati
		fmt.Printf("\n\nTaomlar ro'yxati.\n\n")
		for _,i:=range data{
			fmt.Printf("Id -> %d\nName -> %s\nPrice -> %d\n",i.Id,i.Name,i.Price)
		}

		var b int
		fmt.Printf("\n\nQancha taomni o'chirmoqchisiz: ")
		fmt.Scan(&b)
		for i:=0;i<b;i++{
			var d int
			fmt.Printf("\nTaom idsini kiriting: ")
			fmt.Scan(&d)
			fmt.Println()
			for _,o:=range data{
				if o.Id!=d{
					slice = append(slice, o)
				}
			}
		}
	
		fileread,err2:=json.MarshalIndent(slice,"","    ")
		if err2!=nil{
			fmt.Println("Xatoli?")
			return
		}
		err3:=os.WriteFile(filename,fileread,0644)
		if err3!=nil{
			fmt.Println("Xatolik?")
			return
		}
	}else if a==0{
		return
	}
}