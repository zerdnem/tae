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

func dcipherHash(hash, hashType string) (string, error) {
	var temp string
	temp, err := utils.RetrieveHash(hash, hashType)
	if err != nil {
		return temp, errors.New("Hash could not be deciphered")
	}
	if temp == "" {
		return temp, errors.New("Hash could not be deciphered")
	}
	return temp, nil
}

func dcipher(h string) (string, error) {
	var response string
	hash, err := utils.FromString(h)
	if err != nil {
		return response, errors.New("Hash type not supported")
	}
	hashType := string(hash.Algorithm)
	hashValue := fmt.Sprintf("%x", hash.HashValue)
	if hashType == "sha1" || hashType == "md5" || hashType == "sha256" {
		response, err = dcipherHash(hashValue, hashType)
		if err != nil {
			return response, errors.New("Hash could not be deciphered")
		}
		return response, nil
	} else {
		return response, errors.New("Hash type not supported")
	}
}

func main() {
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	successSymbol := green("✔")
	errorSymbol := red("✖")

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			hash := scanner.Text()
			result, err := dcipher(hash)
			if err != nil {
				fmt.Printf("%s %s", errorSymbol, err)
			}
			fmt.Printf("%s %s", successSymbol, result)
		}
	} else {
		hash := flag.String("hash", "", "Specify a hash to decipher (Required)")
		flag.Parse()
		if *hash == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}
		result, err := dcipher(*hash)
		if err != nil {
			fmt.Printf("%s %s", errorSymbol, err)
		}
		fmt.Printf("%s %s", successSymbol, result)
	}
}
