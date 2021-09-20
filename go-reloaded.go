package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile(os.Args[1])
	toFile, _ := os.Create(os.Args[2])
	if err != nil {
		fmt.Println("No file input", err.Error()) // vb panna error creating a file
	}
	inputStr := string(file)

	chars := []rune(inputStr)
	var words []string
	lastSpace := 0
	for index := 0; index < len(chars); index++ { // Pmst loop-ib üle chars-i ja lisab array-sse kahe space vahel olevad stringid//
		if chars[index] == ' ' {
			words = append(words, string(chars[lastSpace:index]))
			lastSpace = index + 1
		} else if index == len(chars)-1 { // et lisaks ka viimase sõna
			words = append(words, string(chars[lastSpace:index+1])) //
		}
	}

	for index, str := range words {
		if str == "(up," || str == "(low," || str == "(cap," {
			num, err := strconv.Atoi(strings.Trim(words[index+1], ")"))
			if err != nil {
				fmt.Println("Error in converting")
				num = 0
			}
			for i := 1; i <= num; i++ {
				if str == "(up," {
					words[index-i] = strings.ToUpper(words[index-i])
				}
				if str == "(low," {
					words[index-i] = strings.ToLower(words[index-i])
				}
				if str == "(cap," {
					words[index-i] = strings.Title(words[index-i])
				}
			}
			words = remove(words, index)
			words = remove(words, index)
		}
	}

	// teha üks loop kõigi kohta(possibli switchsiga) = possibly efektiivsem ehk parema töötavusega programm

	var s string
	var m string
	for index, j := range words {
		if j == "(hex)" || j == "(bin)" || j == "(up)" || j == "(low)" || j == "(cap)" {
			if j == "(hex)" {
				hex_num, _ := strconv.ParseInt(words[index-1], 16, 64)
				var n int64 = hex_num
				s = strconv.FormatInt(n, 10) // teisendab hex_num-i int64-st stringiks
				words[index] = s             // asendab enne hex-i numbri teisendatud parseint-iga
			} else if j == "(bin)" {
				bin_num, _ := strconv.ParseInt(words[index-1], 2, 64)
				var l int64 = bin_num
				m = strconv.FormatInt(l, 10)
				words[index] = m
			} else if j == "(up)" {
				upper := strings.ToUpper(words[index-1])
				words[index] = upper
			} else if j == "(low)" {
				lower := strings.ToLower(words[index-1])
				words[index] = lower
			} else if j == "(cap)" {
				res := strings.Title(words[index-1])
				words[index] = res
			}
			words = remove(words, index-1) // eemaldab eelmise
		}
	}

	inputStr = strings.Join(words, " ")

	punktpunktpunkt := regexp.MustCompile(`( )(\.\.\.)`)
	inputStr = punktpunktpunkt.ReplaceAllString(inputStr, `$2`)

	mituerinevat := regexp.MustCompile(`( )(!+\?+)`)
	inputStr = mituerinevat.ReplaceAllString(inputStr, `$2`)

	ylakoma := regexp.MustCompile(`( )(')( )(\w)`)
	inputStr = ylakoma.ReplaceAllString(inputStr, ` $2$4`)

	ylakomalõpus := regexp.MustCompile(`(\w+)( )(')($)`)
	inputStr = ylakomalõpus.ReplaceAllString(inputStr, ` $1$3$4`)

	sõnatyhikkoma := regexp.MustCompile(`(\w)( )(,|\?|:|\.|!)(\w)`)
	inputStr = sõnatyhikkoma.ReplaceAllString(inputStr, `$1$3 $4`)

	tyhikkomatyhik := regexp.MustCompile(`( )(,|\?|:|\.|!)( )`)
	inputStr = tyhikkomatyhik.ReplaceAllString(inputStr, `$2 `)

	tyhikylakomatyhik := regexp.MustCompile(`( )(‘)( )(.)`)
	inputStr = tyhikylakomatyhik.ReplaceAllString(inputStr, ` $2$4`)

	tyhikylakomalõpus := regexp.MustCompile(`(\.)( )(‘)`)
	inputStr = tyhikylakomalõpus.ReplaceAllString(inputStr, `$1’`)

	aintoan := regexp.MustCompile(`(a|A)( )([AaEeIiOoUuHh])`)
	inputStr = aintoan.ReplaceAllString(inputStr, `${1}n $3`)

	defer toFile.Close()
	toFile.WriteString(inputStr)
}

func remove(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}
