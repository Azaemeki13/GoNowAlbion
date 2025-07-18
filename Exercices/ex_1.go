package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ressource struct {
	tier     int
	name     string
	quantity int
}

type recipe struct {
	ressource_a ressource
	ressource_b ressource
	ressource_c ressource
}

type city_price struct {
	cityPrice int
	cityType  string
}

type price struct {
	cityPrice city_price
}

type item struct {
	category  string
	name      string
	recipe    recipe
	tier      int
	cityPrice city_price
}

func create_dataset() []item {
	return []item{
		{
			category: "Two_hands",
			name:     "Kingmaker",
			recipe: recipe{
				ressource_a: ressource{tier: 6, name: "Steel_Bar", quantity: 20},
				ressource_b: ressource{tier: 6, name: "Worked_Leather", quantity: 12},
				ressource_c: ressource{tier: 6, name: "Master Remnants of the Old King", quantity: 1},
			},
			tier: 6,
			cityPrice: city_price{
				cityPrice: 200000,
				cityType:  "Thetford",
			},
		},
		{
			category: "CLoth_Armor",
			name:     "Mage_Robe",
			recipe: recipe{
				ressource_a: ressource{tier: 6, name: "Master_Fine_Cloth", quantity: 16},
				ressource_b: ressource{},
				ressource_c: ressource{},
			},
			tier:      6,
			cityPrice: city_price{cityPrice: 60000, cityType: "Thetford"},
		},
		{
			category: "Leather_Helmet",
			name:     "Stalker_Hood",
			recipe: recipe{
				ressource_a: ressource{tier: 6, name: "Worked_Leather", quantity: 8},
				ressource_b: ressource{tier: 6, name: "Master's_Imbued_Visor", quantity: 1},
				ressource_c: ressource{},
			},
			tier:      6,
			cityPrice: city_price{cityPrice: 80000, cityType: "Thetford"},
		},
		{
			category: "Cloth_Boots",
			name:     "Scholar_Sandals",
			recipe: recipe{
				ressource_a: ressource{tier: 6, name: "Ornate Cloth", quantity: 8},
				ressource_b: ressource{},
				ressource_c: ressource{},
			},
			tier:      6,
			cityPrice: city_price{cityPrice: 30000, cityType: "Thetford"},
		},
		{
			category: "Cape",
			name:     "Avalonian_Cape",
			recipe: recipe{
				ressource_a: ressource{tier: 4, name: "Adept's_Cape", quantity: 1},
				ressource_b: ressource{tier: 4, name: "Avalonian's_Crest", quantity: 1},
				ressource_c: ressource{tier: 1, name: "Siphonned_Energy", quantity: 15},
			},
			tier:      6,
			cityPrice: city_price{cityPrice: 180000, cityType: "Thetford"},
		},
	}
}

func controller(prompt string, i *int) {
	switch {
	case prompt == "exit":
		os.Exit(0)
	case prompt == "next":
		*i++
	case prompt == "previous" && *i < 0:
		*i--

	}
}

//https://chatgpt.com/c/68796bc6-8804-8007-9084-362886c8c20e 
// check this to correct ![]item error ;

func main() {

	var index int
	reader := bufio.NewReader((os.Stdin))
	item := create_dataset()
	if ![]item {
		fmt.Println("Error loading the array, exiting program...")
		os.Exit(1)
	}
	fmt.Print("Enter your command :")
	prompt, err := reader.ReadString(' ')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	prompt = strings.TrimSpace(prompt)
	for {
		controller(prompt, &index)
	}

}
