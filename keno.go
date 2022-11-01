// keno
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	//"math/rand"
	//"time"
	"github.com/doun/terminal/color"
)

var count int = 0
var tabla [10000][20]int64
var kilenc [10000]int64 //Az adott húzás találatszáma
var matrix2 [81][81]int64

func main() {

	file, err := os.Open("keno.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	//A record egy kétdimenziós tömb,a 2. dim. egy 1 elemű stringtömb
	//Kovertálás egy 10000*20 méretű numerikus tömbbe
	for _, item := range record {
		s := strings.Split(item[0], ";")
		for i := 4; i < 24; i++ {
			tabla[count][i-4], _ = strconv.ParseInt(s[i], 0, 0)
		}
		count++
	}
	//kilencbolnulla()
	kilences()
	//keverttabla()
	kettesek()
}

//Ellenőrző kiiratás
func ellenorzes() {
	fmt.Println(count)
	for i := 0; i < count; i++ {
		for j := 0; j < 20; j++ {
			fmt.Print(tabla[i][j], " ")
		}
		fmt.Println()
	}
}

func kilencbolnulla() {
	szamok := []int64{3, 10, 13, 24, 35, 39, 44, 56, 67}
	var talalat bool = true
	var mennyi int64 = 0

	for i := 0; i < 40; i++ { //az utolsó 40 húzás
		talalat = true
		for j := 0; j < 20; j++ {
			for k := 0; k < 9; k++ {
				if tabla[i][j] == szamok[k] {
					talalat = false
				}
			}
		}
		if talalat == true {
			fmt.Println(i+1, ". húzásnál 0 találat!")
			mennyi++
		}
	}
	fmt.Println("Öszzesen ", mennyi, " alkalommal lett nulla.")
}
func kilences() {
	var talalat int64
	darab := []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	szamok := []int64{3, 10, 13, 24, 35, 39, 44, 56, 67}

	for i := 0; i < 40; i++ { //az utolsó 40 húzás
		for j := 0; j < 20; j++ {
			for k := 0; k < 9; k++ {
				if tabla[i][j] == szamok[k] {
					talalat++
				}
			}
		}
		kilenc[i] = talalat
		darab[talalat]++
		fmt.Println(i+1, ". húzásnál ", talalat, " találat!")
		talalat = 0
	}
	for i := 0; i < 10; i++ {
		fmt.Println("Kombináció: ", i, ", talalat ", darab[i], " esetben.")
	}
}

func keverttabla() {
	ktabla := []int64{47, 16, 23, 75, 22, 55, 4, 11, 14, 50, 9, 40, 25, 68, 6, 33, 21, 66, 26, 31, 74, 73, 76, 71, 80, 51, 24, 15, 62, 78, 13, 65, 67, 60, 18, 19, 64, 46, 38, 61, 20, 44, 34, 35, 41, 42, 1, 5, 54, 49, 79, 77, 12, 37, 8, 17, 48, 53, 32, 59, 69, 3, 70, 39, 57, 7, 27, 30, 56, 45, 10, 43, 72, 28, 63, 58, 52, 36, 2, 29}
	var red bool

	//for i:=0;i<80;i++{
	//	ktabla[i]=int64(i+1)
	//}
	//rand.Seed(time.Now().UnixNano())
	//rand.Shuffle(len(ktabla), func(i, j int) {ktabla[i], ktabla[j] = ktabla[j], ktabla[i] })

	//fmt.Println(ktabla)

	for k := 0; k < 20; k++ {
		for i := 0; i < 79; i++ {
			red = false
			for j := 0; j < 20; j++ {
				if ktabla[i] == tabla[k][j] {
					red = true
				}
			}
			if red {
				color.Print("@r", ktabla[i], " ")
			} else {
				fmt.Print(ktabla[i], " ")
			}
		}
		fmt.Println()
		fmt.Scanln()
	}
}
func kettesek() {
	for i := 0; i < 20; i++ { //az utolsó 20 húzás
		for j := 0; j < 19; j++ {
			for k := j + 1; k < 20; k++ {
				matrix2[tabla[i][j]][tabla[i][k]]++
			}
		}
	}
	for i := 1; i < 81; i++ {
		for j := 1; j < 81; j++ {
			fmt.Print(matrix2[i][j], " ")
		}
		fmt.Println()
	}
}
