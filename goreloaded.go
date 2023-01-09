package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// removes command from the output
func remove_commands(s []string) string {
	var str string
	for i, com := range s {
		if com == "(cap," || com == "(low," || com == "(up," || com == "(Cap," || com == "(Low," || com == "(Up," {
			s[i] = ""
			s[i+1] = ""
		} else if com != "(hex)" && com != "(bin)" && com != "(up)" && com != "(cap)" && com != "(low)" && com != "" && com != "(Low)" && com != "(Cap)" && com != "(Up)" && com != "(Bin)" && com != "(Hex)" {
			if i == 0 {
				str += com
			} else {
				str += " " + com
			}
		}
	}
	return str
}

// removes the  space before, after and between the quotation marks
func quotes(s string) string {
	str := ""
	var removeSpace bool
	for i, char := range s {
		if char == 39 && s[i-1] == ' ' { // easier to use 39 than the symbol
			if removeSpace {
				str = str[:len(str)-1]
				str += string(char)
				removeSpace = false
			} else {
				str += string(char)
				removeSpace = true
			}
		} else if i > 1 && s[i-2] == 39 && s[i-1] == ' ' {
			if removeSpace {
				str = str[:len(str)-1]
				str += string(char)
			} else {
				str += string(char)
			}
		} else {
			str += string(char)
		}
	}
	return str
}

//  removes whitespaces
// func remove_whitespace(s string) string {
// 	end := len(s) - 1
// 	if s[end-1] == ' ' {
// 		return remove_whitespace(s[:end])
// 	}
// 	return s[:end-1]
// }

// hex and bin decimal values
func base_decimal_converter(s []string) []string {
	for i, word := range s {
		if strings.Compare(word, "(hex)") == 0 {
			num, _ := strconv.ParseInt(s[i-1], 16, 64)
			s[i-1] = fmt.Sprint(num)
		}
		if strings.Compare(word, "(bin)") == 0 {
			num, _ := strconv.ParseInt(s[i-1], 2, 64)
			s[i-1] = fmt.Sprint(num)
		}
	}
	return s
}

// convert the word before the command to lower case, upper case or title case. If extra parameter (number) is provided
// it will convert the number of words to the given case in a reverse manner
func wordcase(s []string) []string {
	for i, word := range s {
		if strings.Compare(word, "(low)") == 0 {
			s[i-1] = strings.ToLower(s[i-1])
		}
		if strings.Compare(word, "(up)") == 0 {
			s[i-1] = strings.ToUpper(s[i-1])
		}
		if strings.Compare(word, "(cap)") == 0 {
			s[i-1] = strings.Title(strings.ToLower(s[i-1]))
		}
		if strings.Compare(word, "(low,") == 0 {
			s[i-1] = strings.ToLower(s[i-1])
			size := len(s[i+1])
			value := s[i+1][:size-1]
			num, _ := strconv.Atoi(value)
			for j := 1; j <= num; j++ {
				s[i-j] = strings.ToLower(s[i-j])
			}
		}
		if strings.Compare(word, "(up,") == 0 {
			s[i-1] = strings.ToUpper(s[i-1])
			size := len(s[i+1])
			value := s[i+1][:size-1]
			num, _ := strconv.Atoi(value)
			for j := 1; j <= num; j++ {
				s[i-j] = strings.ToUpper(s[i-j])
			}
		}
		if strings.Compare(word, "(cap,") == 0 {
			s[i-1] = strings.Title(strings.ToLower(s[i-1]))
			size := len(s[i+1])
			value := s[i+1][:size-1]
			num, _ := strconv.Atoi(value)
			if num >= len(s[:i]) {
				num = len(s[:i])
			}
			for j := 1; j <= num; j++ {
				s[i-j] = strings.Title(strings.ToLower(s[i-j]))
			}
		}
	}
	return s
}

// converting a into an when the next word is vowel or h
func article_check(s []string) []string {
	for i, word := range s {
		if strings.Compare(word, "A") == 0 && (string(s[i+1][0]) == "a") {
			s[i] = "An"
		}
		if strings.Compare(word, "a") == 0 && (string(s[i+1][0]) == "a") {
			s[i] = "an"
		}
		if strings.Compare(word, "A") == 0 && (string(s[i+1][0]) == "e") {
			s[i] = "An"
		}
		if strings.Compare(word, "a") == 0 && (string(s[i+1][0]) == "e") {
			s[i] = "an"
		}
		if strings.Compare(word, "A") == 0 && (string(s[i+1][0]) == "i") {
			s[i] = "An"
		}
		if strings.Compare(word, "a") == 0 && (string(s[i+1][0]) == "i") {
			s[i] = "an"
		}
		if strings.Compare(word, "A") == 0 && (string(s[i+1][0]) == "o") {
			s[i] = "An"
		}
		if strings.Compare(word, "a") == 0 && (string(s[i+1][0]) == "o") {
			s[i] = "an"
		}
		if strings.Compare(word, "A") == 0 && (string(s[i+1][0]) == "u") {
			s[i] = "An"
		}
		if strings.Compare(word, "a") == 0 && (string(s[i+1][0]) == "u") {
			s[i] = "an"
		}
		if strings.Compare(word, "A") == 0 && (string(s[i+1][0]) == "h") {
			s[i] = "An"
		}
		if strings.Compare(word, "a") == 0 && (string(s[i+1][0]) == "h") {
			s[i] = "an"
		}
	}
	return s
}

// checks the characters and adjust the space before or after
func character_space_check(str string) string {
	var word string
	for i, char := range str {
		if i == len(str)-1 {
			if char == '.' || char == ',' || char == '!' || char == '?' || char == ';' || char == ':' {
				if str[i-1] == ' ' {
					word = word[:len(word)-1]
					word += string(char)
				} else {
					word += string(char)
				}
			} else {
				word += string(char)
			}
		} else if char == '.' || char == ',' || char == '!' || char == '?' || char == ';' || char == ':' {
			if str[i-1] == ' ' {
				word = word[:len(word)-1]
				word += string(char)
			} else {
				word += string(char)
			}
			if str[i+1] != ' ' && str[i+1] != '.' && str[i+1] != '!' && str[i+1] != '?' && str[i+1] != ';' && str[i+1] != ':' && str[i+1] != ',' {
				word += " "
			}
		} else {
			word += string(char)
		}
	}
	return word
}

//  function takes string argument and write the edited output into a file
func Goreloaded(args string) string {
	// reading the input file into data
	data, err := os.ReadFile(args)
	if err != nil {
		log.Fatal(err)
	}
	input := string(data)
	input_slices := strings.Split(input, " ")
	// modifying the input file
	first_step := base_decimal_converter(input_slices)
	second_step := wordcase(first_step)
	third_step := article_check(second_step)
	str := remove_commands(third_step)
	word := character_space_check(str)
	word = quotes(word)
	return word
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Arguments size error")
		os.Exit(0)
	}
	word := Goreloaded(os.Args[1])
	output := []byte(word)
	error := os.WriteFile(os.Args[2], output, 0o666)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("done")
}
