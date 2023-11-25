package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	//basic
	pattern := "go"
	text := "golang is a powerful lanuage"
	mached, err := regexp.MatchString(pattern, text)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(mached)
	//search
	pattern1 := `(\d{4})-(\d{2})-(\d{2})`
	text1 := "2023-10-25,2022-05-15"
	//search
	re := regexp.MustCompile(pattern1)
	matches := re.FindAllStringSubmatch(text1, -1)
	//replace
	replaced := re.ReplaceAllString(text1, "[$1][$2][$3]")
	fmt.Println("Replaced text:", replaced)

	for _, matche := range matches {
		fmt.Println("Full match:", matche[0])
		fmt.Println("year match:", matche[1])
		fmt.Println("monch match:", matche[2])
		fmt.Println("day match:", matche[3])
	}
}
