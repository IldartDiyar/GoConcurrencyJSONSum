package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
)

type Data struct {
	A int `json:"a"`
	B int `json:"b"`
}

func main() {
	data := []Data{}
	for i := 0; i <= 1_000_000; i++ {
		b := rand.Intn(21) - 10
		a := rand.Intn(21) - 10
		data = append(data, Data{A: a, B: b})
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Can not marshal due to %v", err)
	}

	os.WriteFile("../jj.json", jsonData, 777)
}
