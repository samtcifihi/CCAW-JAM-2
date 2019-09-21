package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Genned Word: ", genWord()) // Debugging
	fmt.Println("Execution Completed")      // Debugging
}

// Generate a pseudo-word for the English language
func genWord() string {
	onset := []string{"1", "2"}
	nucleus := []string{"3", "4"}
	coda := []string{"5", "6"}

	word := genCluster(onset) + genCluster(nucleus) + genCluster(coda)

	return word
}

// Choose a random element from an array
func genCluster(validStrings []string) string {
	// Seeded Random Number Generator
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	cluster := validStrings[randGen.Intn(len(validStrings))]

	return cluster
}
