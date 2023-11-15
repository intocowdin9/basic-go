package main

import (
	"fmt"
	"regexp"
)

func main() {
	// penerapan regexp
	var text = "apple burger soup"
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	var res1 = regex.FindAllString(text, 2)
	fmt.Printf("%#v \n", res1)

	var res2 = regex.FindAllString(text, -1)
	fmt.Printf("%#v \n", res2)

	// method MatchString()
	var text2 = "grape into the sky"
	var regex2, _ = regexp.Compile(`[a-z]+`)

	var str = regex2.FindString(text2)
	fmt.Println(str)

	// method FindStringIndex()
	var text3 = "manggo top level"
	var regex3, _ = regexp.Compile(`[a-z]+`)

	var idx = regex3.FindStringIndex(text3)
	fmt.Println(idx)

	var str3 = text3[0:6]
	fmt.Println(str3)

	// method FindAllString()
	var text4 = "coffee of darkness"
	var regex4, _ = regexp.Compile(`[a-z]+`)

	var str4 = regex4.FindAllString(text4, -1)
	fmt.Println(str4)

	var strr5 = regex4.FindAllString(text4, 1)
	fmt.Println(strr5)

	// method ReplaceAllString()
	var text5 = "teh hangat"
	var regex5, _ = regexp.Compile(`[a-z]+`)

	var str5 = regex5.ReplaceAllString(text5, "iemon")
	fmt.Println(str5)

	// method ReplaceAllStringFunc()
	var text6 = "air terjun"
	var regex6, _ = regexp.Compile(`[a-z]+`)

	var str6 = regex6.ReplaceAllStringFunc(text6, func(each string) string {
		if each == "terjun" {
			return "hujan"
		}
		return each
	})
	fmt.Println(str6)

	// method Split()
	var text7 = "apple, rasa, anggur"
	var regex7, _ = regexp.Compile(`[a-z]+`)

	var str7 = regex7.Split(text7, -1)
	fmt.Printf("%#v \n", str7)
}
