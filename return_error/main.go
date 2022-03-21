package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Roulette() (string, error) {
	number := rand.Float64()
	if number < 0.4 {
		return "", errors.New("missed this round")
	}
	return "win!", nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	value, err := Roulette()
	if err != nil {
		fmt.Printf("roulette error: %v\n", err)
		return
	}
	fmt.Println(value)
}
