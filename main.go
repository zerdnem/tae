package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"./utils"
	"github.com/fatih/color"
)

type hashes struct {
	*utils.Hashes
}

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

func printHashes(text string, ha *hashes) {
	hashes := utils.GenerateHash(text)
	ha.Hashes = &hashes
	fmt.Print(
		"md5    "+ha.Md5+"\n",
		"sha1   "+ha.Sha1+"\n",
		"sha256 "+ha.Sha256+"\n",
		"sha384 "+ha.Sha384+"\n",
		"sha512 "+ha.Sha512+"\n",
	)
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
		text := flag.String("generate", "", "Generate hash from string")
		var ha hashes
		flag.Parse()
		if *hash != "" {
			symbol := dcipher(*hash)
			fmt.Println(symbol)
		}
		if *text != "" {
			printHashes(*text, &ha)
		}
		symbol := decrypt(*hash)
		fmt.Println(symbol)
	}
}
