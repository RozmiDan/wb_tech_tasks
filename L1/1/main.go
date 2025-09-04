package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
Подсказка: используйте композицию (embedded struct), чтобы Action имел все методы Human.
*/

type Human struct {
	Height int
	Weight int
	Nation string
}

func (h *Human) SpeakHumanMethod(word string) {
	fmt.Printf("Human can speak, he said: \"%s\" word\n", word)
}

func (h *Human) WalkHumanMethod(meters int) {
	fmt.Printf("Human can walk, he walked %d meters\n", meters)
}

type Action struct {
	*Human
	FstField  string
	ScndField int
}

func (a *Action) FstMethodAction() {
	fmt.Println("some first action")
}

func (a *Action) ScndMethodAction() {
	fmt.Println("some second action")
}

func (a *Action) WalkHumanMethod(meters int) { // Shadowing example
	fmt.Printf("Action can walk, he walked %d meters\n", meters)
}

func main() {
	person := &Human{
		Height: 53,
		Weight: 196,
		Nation: "russian",
	}
	person.SpeakHumanMethod("hello")
	person.WalkHumanMethod(234)

	act := &Action{
		Human:     &Human{Height: 178, Weight: 86, Nation: "german"},
		FstField:  "aboba",
		ScndField: 23,
	}
	act.SpeakHumanMethod("hi")
	act.WalkHumanMethod(1000)
}
