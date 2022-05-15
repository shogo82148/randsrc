package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
)

func main() {
	p, err := generate()
	if err != nil {
		log.Fatal(err)
	}
	q, err := generate()
	if err != nil {
		log.Fatal(err)
	}
	m := new(big.Int).Mul(p, q)
	fmt.Print("[]byte{")
	data := m.Bytes()
	for i, b := range data {
		if i%16 == 0 {
			fmt.Println()
		}
		fmt.Printf("0x%02x, ", b)
	}
	fmt.Println("\n}")
}

func generate() (*big.Int, error) {
	for i := 0; i < 1000; i++ {
		p, err := rand.Prime(rand.Reader, 2048)
		if err != nil {
			return nil, err
		}
		bytes := p.Bytes()
		if bytes[len(bytes)-1]%4 == 3 {
			return p, nil
		}
	}
	return nil, errors.New("failed to generate param")
}
