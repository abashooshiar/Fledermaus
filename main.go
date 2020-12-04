package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	dimensionsArray := [20][3]int{
		{2, 4, 3}, {1, 5, 6}, {1, 7, 8}, {1, 9, 10}, {2, 10, 11},
		{2, 7, 12}, {3, 6, 13}, {3, 9, 14}, {4, 8, 15}, {4, 5, 16},
		{5, 12, 17}, {6, 11, 18}, {7, 14, 18}, {8, 13, 19}, {9, 16, 19},
		{10, 15, 17}, {11, 16, 20}, {12, 13, 20}, {14, 15, 20}, {17, 18, 19}}

	rand.Seed(time.Now().UnixNano())
	raum := randomint(1, 20)
	zufallszahl := randomint(1, 20)
	slice := make([]int, 4)
	slice[0] = randomint(1, 20)
	slice[1] = randomint(1, 20)
	slice[2] = randomint(1, 20)
	slice[3] = randomint(1, 20)

	startraum(slice, raum, zufallszahl, dimensionsArray)
	//if heimstart == false {
	//	fmt.Println("Zufallszahl(", zufallszahl, ")","Loch(", slice[0], ",", slice[1], ")","Fledermaus(", slice[2], ",", slice[3], ")")
	//	angrenzend(raum, dimensionsArray, zufallszahl, slice)
	//} else if heimstart == true {
	//	fmt.Println("im startraum() ist was komisch!")
	//}
}
func angrenzend(raum int, dimensionsArray [20][3]int, zufallszahl int, slice []int) {
	fmt.Println("Du befindest dich in Raum : ", raum)
	fmt.Println("Angrenzende Räume sind : ", dimensionsArray[raum-1])

	for _, inhalts := range dimensionsArray[raum-1] {
		if inhalts == zufallszahl {
			fmt.Println("Es stinkt bestialisch.")
		}
		if inhalts == slice[0] || inhalts == slice[1] {
			fmt.Println("Du spürst einen Luftzug.")
		}
		if inhalts == slice[2] || inhalts == slice[3] {
			fmt.Println("Du hörst ein Flattern.")
		}
	}

	fmt.Print("Wohin möchtest Du gehen? : ")
	var wohingehen int
	if _, err := fmt.Scan(&wohingehen); err != nil {
		fmt.Print("Bitte geben Sie eine Nummer__ ")
	}

	if wohingehen == zufallszahl {
		fmt.Println("Du hast das Monster gefunden!")
		fmt.Println("drücken Sie 'j', um das Spiel zu beenden.Andernfalls geben Sie eine Zahl ein.")

		var exit int
		if _, err := fmt.Scan(&exit); err != nil {
			os.Exit(0)
		}
	}

	if wohingehen == slice[0] || wohingehen == slice[1] {
		fmt.Println("Du fällst in ein bodenloses Loch und stirbst.")
		os.Exit(0)

	}
	finder(raum, wohingehen, dimensionsArray, zufallszahl, slice)
}

func randomint(min, max int) int {
	return min + rand.Intn(max-min)
}
func finder(raum, wohingehen int, dimensional [20][3]int, zufallszahl int, slice []int) {
	indexOfDimensional := dimensional[raum-1]

	sw, _ := findin(indexOfDimensional[:], wohingehen, zufallszahl, indexOfDimensional[:])
	if sw == true {
		fmt.Println("----------------------------")
		fmt.Println("Zufallszahl*(", zufallszahl, ")", "Loch*(", slice[0], ",", slice[1], ")", "Fledermaus*(", slice[2], ",", slice[3], ")")

		raum = wohingehen
		angrenzend(raum, dimensional, zufallszahl, slice)
	} else {
		fmt.Println("Du kannst von hier nicht in Raum ", wohingehen)
		angrenzend(raum, dimensional, zufallszahl, slice)
	}
}

func findin(arrayinhalt []int, wohingehen int, zufallszahl int, s []int) (bool, int) {
	switchKey := false
	for _, inhalt := range arrayinhalt {
		if inhalt == wohingehen {
			switchKey = true
		}
	}
	return switchKey, 0
}

func startraum(slice []int, raum int, zufallszahl int, dimensionsArray [20][3]int) {

	switchKey := false
	count := 0
	if dimensionsArray[raum-1][0] == slice[0] || dimensionsArray[raum-1][0] == slice[1] || dimensionsArray[raum-1][0] == slice[2] || dimensionsArray[raum-1][0] == slice[3] ||
		dimensionsArray[raum-1][1] == slice[0] || dimensionsArray[raum-1][1] == slice[1] || dimensionsArray[raum-1][1] == slice[2] || dimensionsArray[raum-1][1] == slice[3] ||
		dimensionsArray[raum-1][2] == slice[0] || dimensionsArray[raum-1][2] == slice[1] || dimensionsArray[raum-1][2] == slice[2] || dimensionsArray[raum-1][2] == slice[3] ||
		slice[0] == slice[2] || slice[0] == slice[3] || slice[1] == slice[2] || slice[1] == slice[3] ||
		zufallszahl == slice[0] || zufallszahl == slice[1] || zufallszahl == slice[2] || zufallszahl == slice[3] ||
		zufallszahl == dimensionsArray[raum-1][0] || zufallszahl == dimensionsArray[raum-1][1] || zufallszahl == dimensionsArray[raum-1][2] {

		fmt.Println(count)
		fmt.Println("raum:", raum)
		fmt.Println("zufallszahl:", zufallszahl)
		rand.Seed(time.Now().UnixNano())
		raum = randomint(1, 20)
		zufallszahl = randomint(1, 20)
		slice[0] = randomint(1, 20)
		slice[1] = randomint(1, 20)
		slice[2] = randomint(1, 20)
		slice[3] = randomint(1, 20)
		fmt.Println("raum*:", raum)
		fmt.Println("zufallszahl*:", zufallszahl)
		switchKey = true

	}
	if switchKey == true {
		startraum(slice, raum, zufallszahl, dimensionsArray)

		//reflect.Call() Function with new variable +golang

	} else {
		switchKey = false
		fmt.Println("ok")
		fmt.Println("Zufallszahl(", zufallszahl, ")", "Loch(", slice[0], ",", slice[1], ")", "Fledermaus(", slice[2], ",", slice[3], ")")
		angrenzend(raum, dimensionsArray, zufallszahl, slice)
	}

}
