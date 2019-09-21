package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("You will be shown a chosen number of nonce constructions.")
	fmt.Println("You will then attempt to type them in accurately.")
	fmt.Println("Everything will be lowercase, but be conservative as there are no misakes accepted.")
	fmt.Println()

	var wordsPerTrial uint64
	prompt := ""

	var stillPlaying = true
	for stillPlaying == true {
		getNumberOfWords(&wordsPerTrial)

		prompt = ""
		var i uint64
		for i < wordsPerTrial {
			prompt += genWord()
			prompt += " "

			i++
		}
		fmt.Println(prompt)
		fmt.Println()

		// Get user input
		// Compare to prompt to determine if won; check (len(prompt) - 1) / 5 to determine score (the -1 is for the space at the end)
		// Or just check wordsPerTrial

		// ask if the user wishes to play again
	}

	fmt.Println("Genned Word: ", genWord()) // Debugging
	fmt.Println("Execution Completed")      // Debugging
}

func getNumberOfWords(wordsPerTrial *uint64) {
	*wordsPerTrial = 0
	cin := bufio.NewReader(os.Stdin)

	validUserChoice := false
	for validUserChoice == false {
		validUserChoice = true // Assume input is valid until proven otherwise
		fmt.Println("How many words would you like to memorize?")

		userInput, _ := cin.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		uintUserInput, err := strconv.ParseUint(userInput, 10, 64)
		if err != nil {
			fmt.Println(err)

			validUserChoice = false
		}

		if validUserChoice == true {
			*wordsPerTrial = uintUserInput
		}
	} // User has entered a valid choice
}

// Generate a pseudo-word for the English language
func genWord() string {
	onset := []string{
		"",
		"m",
		"n",
		"p",
		"t",
		"ch",
		"c",
		"b",
		"d",
		"j",
		"g",
		"f",
		"ph",
		"th",
		"s",
		"sh",
		"h",
		"v",
		"z",
		"x",
		"l",
		"r",
		"y",
		"w",
		"pl",
		"bl",
		"cl",
		"ɡl",
		"pr",
		"br",
		"tr",
		"dr",
		"cr",
		"ɡr",
		"fl",
		"sl",
		"fr",
		"phr",
		"shr",
		"sw",
		"sp",
		"st",
		"sk",
		"sm",
		"sn",
		"sph",
		"spl",
		"scl",
		"spr",
		"str",
		"scr",
	}
	nucleus := []string{
		"a",
		"o",
		"au",
		"ou",
		"ao",
		"i",
		"e",
		"u",
		"oo",
		"ue",
		"ei",
		"ey",
		"ay",
		"oa",
		"oe",
		"ee",
		"ea",
		"ie",
		"ay",
		"ai",
		"oi",
		"oy",
		"eu",
		"ow",
		"y",
		"eau",
	}
	coda := []string{
		"",
		"m",
		"n",
		"mn",
		"ng",
		"p",
		"t",
		"ch",
		"ck",
		"b",
		"d",
		"g",
		"ff",
		"ph",
		"th",
		"ss",
		"sh",
		"v",
		"z",
		"ll",
		"r",
		"lp",
		"lb",
		"lt",
		"ld",
		"lch",
		"lge",
		"lk",
		"rp",
		"rb",
		"rt",
		"rd",
		"rch",
		"rge",
		"rk",
		"rg",
		"lf",
		"lve",
		"lth",
		"lse",
		"ls",
		"rsh",
		"lm",
		"ln",
		"rm",
		"rn",
		"rle",
		"mp",
		"nt",
		"nd",
		"nch",
		"nge",
		"nk",
		"mph",
		"nth",
		"nce",
		"ns",
		"ngth",
		"ft",
		"sp",
		"st",
		"sk",
		"fth",
		"pt",
		"ct",
		"pth",
		"pse",
		"tz",
		"dth",
		"x",
		"lpt",
		"lps",
		"ltz",
		"lst",
		"rmth",
		"rpt",
		"rpse",
		"rtz",
		"rst",
		"mpt",
		"mpse",
		"nct",
		"nx",
		"xt",
	}

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
