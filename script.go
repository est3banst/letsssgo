package main

func main() {
	cards := newDeckFromFile("my_cards2")
	cards.print()

	greeting := "Hi there"

	println([]byte(greeting))
}
