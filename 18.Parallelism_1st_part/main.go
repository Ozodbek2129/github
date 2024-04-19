package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func fak(s int,ch chan int){
	FakHisob:=1
	for i:=1;i<=s;i++{
		FakHisob*=i
	}
	ch<-FakHisob
}

func main(){
	ch1:=make(chan int)
	var son int
	fmt.Printf("Sonni kiriting: ")
	fmt.Scan(&son)
	go fak(son,ch1)
	time.Sleep(1 * time.Millisecond)
	fmt.Println("Siz kiritgan sonning faktoriali->",<-ch1)

	files:=[]string{"file1.txt","file2.txt"}
	var file string
	for _,v:=range files{
		f,err:=os.ReadFile(v)
		if err!=nil{
			fmt.Println("File ni o'qishda xatolik?")
			return
		}
		for _,i:=range strings.Split(string(f)," "){
			file+=i
			file+=" "
		}
	}
	fmt.Println(file)
}