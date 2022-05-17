package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
	"unicode"
)

func getInput() string {
	result, _ := bufio.NewReader(os.Stdin).ReadString('\n') // получили входную строку
	return result
}

func converString(input string) (result []big.Int, err error) {
	if len(input) == 0 {
		return nil, err // если строка пустая, выводим ошибку
	}
	rs := strings.Split(input, " ") // конвертируем в руновый слайс
	for indx, value := range rs {
		value = strings.TrimSpace(value)
		symbols := []rune(value)
		     for _, value := range symbols{
		         if  unicode.IsDigit(value) != true {
		         	return nil, err // одно из значений - не число, выводим ошибку
		         }
		     }
		if indx > 1 {
			return nil, err // чисел в строке больше двух, выводим ошибку
		}
		bigIntValue  := big.NewInt(1)
		bigIntValue.SetString(value, 10)
		result = append(result, *bigIntValue)
	}
	return result, nil
}

func summ(a *big.Int, b *big.Int) (*big.Int) {
	result  := big.NewInt(0)
	result = result.Add(result, a)
  result = result.Add(result, b)
	return result
}

func main() {
	// Напишите ваш код здесь. Разбейте на функции, если считаете нужным.
	input := getInput()                     // получили входную строку
	bigIntSlice, err := converString(input) // превратили ее в слайс больших чисел
	if err != nil || len(bigIntSlice) < 1 {
		fmt.Println("not digit") // выведем not digit
		return                   //закончим
	}
	a := bigIntSlice[0]
  b := bigIntSlice[1]
	result := summ(&a, &b)
	fmt.Println(result) // выведем результат сложения
	//fmt.Println("RESULT")
}
