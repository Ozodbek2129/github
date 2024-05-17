package test2

import (
	"encoding/json"
	"fmt"
	"os"
	test "restarant/Admin"
	"strconv"
)

func Taomlar() {
	file, err := os.ReadFile("taomlar_ruyxat.json")
	if err != nil {
		fmt.Println("Malumotlani o'qishda xatolik?")
		return
	}

	var taom []test.Taomlar_ruyxati
	er := json.Unmarshal(file, &taom)
	if er != nil {
		fmt.Println("Json file ni o'qishda xatolik?")
		return
	}

	for _, v := range taom {
		if len(taom)==0{
			fmt.Println("Bizda taomlar tugab qoldi.")
			return
		}
		fmt.Printf("\n\n----------------------------------------------------")
		fmt.Printf("Id -> %d\nName -> %s\nNarx -> %d\n", v.Id, v.Name, v.Price)
	}
}

func Zakas() {
	var taom2 []test.Taomlar_ruyxati
	file, er := os.ReadFile("taomlar_ruyxat.json")
	if er != nil {
		fmt.Println("File ni o'qishda xatolik?")
		return
	}
	errr := json.Unmarshal(file, &taom2)
	if errr != nil {
		fmt.Println("Xatolik?")
		return
	}

	var slice []string
	cnt:=0
	for {
		var h int
		fmt.Printf("\n\nYakunlash -> 0\nTaom id sini kiriting: ")
		fmt.Scan(&h)
		fmt.Println()
		if h == 0 {
			break
		}
		found := false
		for _, j := range taom2 {
			if j.Id == h {
				l := strconv.Itoa(j.Id)
				l1 := strconv.Itoa(j.Price)
				slice = append(slice, l, j.Name, l1)
				found = true
				cnt +=j.Price
				break
			}
		}
		if !found {
			fmt.Println("Ming afsuski bunaqa aydi yo'q?")
		}

	}
	fmt.Printf("\n\nSizdan %d so'm bo'ldi.\nSiz buyurgan taomlar ruyxati -> %v", cnt, slice)
}
