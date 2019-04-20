package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	name := prompt("Site name:")
	author := prompt("Author:")
	needJSFolder := prompt("Do you want a folder for JavaScript?")
	needCSSFolder := prompt("Do you want a folder for CSS?")
	//debug
	fmt.Println(name, author, needJSFolder, needCSSFolder)

	if dirpath := "./awesomeco"; isNotExist(dirpath) {
		err := os.Mkdir(dirpath, os.ModePerm)
		check(err)
	}

	if filepath := "./awesomeco/index.html"; isNotExist(filepath) {
		outputFile, err := os.Create(filepath)
		check(err)
		defer outputFile.Close()

		html := fmt.Sprintf(htmlTemplate(), name, author)
		outputFile.Write([]byte(html))
	}

	if needJSFolder == "y" {
		jspath := "./awesomeco/js"
		if isNotExist(jspath) {
			err := os.Mkdir(jspath, os.ModePerm)
			check(err)
		}
	}
	if needCSSFolder == "y" {
		csspath := "./awesomeco/css"
		if isNotExist(csspath) {
			err := os.Mkdir(csspath, os.ModePerm)
			check(err)
		}
	}
}

func prompt(msg string) (input string) {
	if !strings.HasSuffix(msg, " ") {
		msg += " "
	}
	fmt.Print(msg)

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input += scanner.Text()
	}

	return
}

func isNotExist(path string) bool {
	info, err := os.Stat(path)
	fmt.Println("info ", info, "\n", "err", err, "\n")
	ret := os.IsNotExist(err)
	fmt.Println("path is not exist: ", ret)
	return ret
}

func htmlTemplate() string {
	filename := "./template/indexhtml"
	buf, err := ioutil.ReadFile(filename)
	check(err)

	return string(buf)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
