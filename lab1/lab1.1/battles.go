package main

import "fmt"

func main() {

	for i := 99; i > 1; i-- {
		ending1 := getEnding(i)
		ending2 := getEnding(i - 1)
		fmt.Printf("%v бутыл%v пива на стене\n%v бутыл%v пива!\nВозьми одну, пусти по кругу\n%v бутыл%v пива на стене!\n\n", i, ending1, i, ending1, i - 1, ending2)
	}

	fmt.Print("Нет бутылок пива на стене!\nНет бутылок пива!\nПойди в магазин и купи ещё,\n99 бутылок пива на стене!")
}

func getEnding(count int) string {
	if count >= 5 && count <= 20 {
		return "ок"
	}

	reminder := count%10
	if reminder == 1 {
		return "ка"
	}
	if reminder >= 2 && reminder <= 4 {
		return "ки"
	}

	return "ок"
}

