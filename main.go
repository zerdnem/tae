package main

import (
	"./utils"
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"os"
)

var (
	green         = color.New(color.FgGreen).SprintFunc()
	red           = color.New(color.FgRed).SprintFunc()
	successSymbol = green("✔")
	errorSymbol   = red("✖")
)

func dcipherHash(hash, hashType string) (string, error) {
	var temp string
	temp, err := utils.RetrieveHash(hash, hashType)
	if err != nil || temp == "" {
		return temp, errors.New("Hash could not be deciphered")
	}
	return temp, nil
}

func dcipher(h string) (string, error) {
	var response string
	hash, err := utils.FromString(h)
	if err != nil {
		return nil, err
	}
	hashType := string(hash.Algorithm)
	hashValue := fmt.Sprintf("%x", hash.HashValue)

	response, err = dcipherHash(hashValue, hashType)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func displaySymbol(hash interface{}) {
	s, _ := hash.(string)
	result, err := dcipher(s)
	if err != nil {
		fmt.Printf("%s %s", errorSymbol, err)
	}
	fmt.Printf("%s %s", successSymbol, result)
}

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			hash := scanner.Text()
			displaySymbol(hash)
		}
	} else {
		hash := flag.String("hash", "", "Specify a hash to decipher (Required)")
		flag.Parse()
		if *hash == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}
		displaySymbol(*hash)
	}
}
