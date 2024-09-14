package main

import (
	"encoding/json"
	"fmt"
)

type Agent struct {
	Name        string   `json:"name"`
	Catagory    string   `json:"type"`
	Abilities   []string `json:"abilities"`
	IsAvailable bool     `json:"available"`
}

func EncodeSliceJSON(slice []Agent) string {
	byteCode, err := json.MarshalIndent(slice, "", "\t")

	if err != nil {
		panic(err)
	}
	// fmt.Println(byteCode)
	return string(byteCode)
}

func main() {

	agents := []Agent{
		{"Sage", "Sentinal", []string{"Heal", "Wall", "SlowOrb", "Revibe"}, true},
		{"Omen", "Controller", []string{"Smoke", "Paranoia", "Teleport", "Vanish"}, true},
		{"Reyna", "Duelist", []string{"Blind", "Heal", "Dismiss", "Empress"}, true},
		{"Fade", "Initator", []string{"Eye", "Blindog", "Sucktion", "Fear"}, false},
	}

	content := EncodeSliceJSON(agents)

	// fmt.Printf("Type is %T\n", content)
	fmt.Println(content)

}
