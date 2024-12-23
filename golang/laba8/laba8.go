package laba8

import (
	"fmt"
)

var filePath string = "../golang/laba8/"

func RunLab8() {
	var filename string
	fmt.Println("Введите имя файла, в который вы хотите записать данные:")
	fmt.Scan(&filename)
	name, _ := CreateFile(filename)

	fmt.Println("Введите данные в файл (для завершения ввода введите q):")
	WriteDataToFile(name)

	slice := ReadDataFromFile(name)

	fmt.Printf("Чтение файла %s завершено\n", filename)
	for _, str := range slice {
		fmt.Println(str)
	}

	fmt.Println("Введите строку, которую хотите найти в файле: ")
	var StringForFind string
	fmt.Scan(&StringForFind)
	n, _ := SearchingString(name, StringForFind)
	if n == 1 {
		fmt.Printf("Слово найденно на строке %d", n)
	} else {
		fmt.Println("Слово не найденно")
	}
}
