package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"./utils"
	"github.com/fatih/color"
)

var (
	green         = color.New(color.FgGreen).SprintFunc()
	red           = color.New(color.FgRed).SprintFunc()
	successSymbol = green("✔")
	errorSymbol   = red("✖")
)

func dcipher(h string) string {
	hash, err := utils.FromString(h)
	if err != nil {
		return fmt.Sprintf("%s %s", errorSymbol, "Only md5, sha1, sha256 are supprted")
	}
	hashType := string(hash.Algorithm)
	hashValue := fmt.Sprintf("%x", hash.HashValue)

	result := utils.RetrieveHash(hashValue, hashType)
	if result == "" {
		return fmt.Sprintf("%s %s", errorSymbol, "Hash not found")
	}
	return fmt.Sprintf("%s %s", successSymbol, result)
}

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			hash := scanner.Text()
			symbol := dcipher(hash)
			fmt.Println(symbol)
		}
	} else {
		hash := flag.String("hash", "", "Specify a hash to decipher (Required)")
		flag.Parse()
		if *hash == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}
		symbol := dcipher(*hash)
		fmt.Println(symbol)
	}
}
