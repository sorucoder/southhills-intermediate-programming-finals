package main

import "fmt"

var dictionary map[string]string = map[string]string{
	"pig":      "Any of various mammals of the family Suidae, having short legs, hooves with two weight-bearing toes, bristly hair, and a cartilaginous snout used for digging, including the domesticated hog (Sus scrofa subsp. domestica syn. S. domesticus) and wild species such as the bushpig.",
	"computer": "A device that computes, especially a programmable electronic machine that performs high-speed mathematical or logical operations or that assembles, stores, correlates, or otherwise processes information.",
	"wallet":   "A flat pocket-sized folding case, usually made of leather, for holding paper money, cards, or photographs; a billfold.",
}

func main() {
	// If the map has the key, it returns the value:
	fmt.Println(dictionary["pig"])

	// If the value does not exist, it returns nil:
	fmt.Println(dictionary["word"])

	// You can check if a value exists by doing this:
	if definition, exists := dictionary["computer"]; exists {
		fmt.Println(definition)
	}
	if _, exists := dictionary["word"]; !exists {
		fmt.Println(`"word" does not exist`)
	}
}
