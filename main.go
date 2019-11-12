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

func decrypt(h string) string {
	hashtype := utils.HashType(h)
	if hashtype == "" {
		return fmt.Sprintf("%s %s", errorSymbol, "Hash not supported")
	}

	decrypted := utils.DecryptHash(h, hashtype)
	if decrypted == "" {
		return fmt.Sprintf("%s %s", errorSymbol, "Hash not found")
	}
	return fmt.Sprintf("%s decrypted=%s hashtype=%s", successSymbol, decrypted, hashtype)
}

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			hash := scanner.Text()
			symbol := decrypt(hash)
			fmt.Println(symbol)
		}
	} else {
		hash := flag.String("hash", "", "Specify a hash to decipher (Required)")
		flag.Parse()
		if *hash == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}
		symbol := decrypt(*hash)
		fmt.Println(symbol)
	}
}
