package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type IndxSlice struct {
	x uint
	y uint
	z uint
}

func getInput() string {
	result, _ := bufio.NewReader(os.Stdin).ReadString('\n') // получили входную строку
	return result
}

func converString(input string) (result []int, err error) {
	if len(input) == 0 {
		return nil, err // если строка пустая, выводим ошибку
	}
	rs := strings.Split(input, ",") // конвертируем в руновый слайс
	for _, value := range rs {
		value = strings.TrimSpace(value)
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		result = append(result, intValue)
	}
	//fmt.Println(result)
	return result, nil
}

func sortSlice(input []int) (result []int, err error) {
	var indxSlice []string
	var flag bool
	var flag0 bool
	if len(input) == 0 {
		return nil, err // если строка пустая, выводим ошибку
	}
	for indx, value := range input {
		for indx1, value1 := range input {
			for indx2, value2 := range input {
				if value+value1+value2 == 0 {
					sliceForSort := []int{indx, indx1, indx2}
					sort.Ints(sliceForSort)
					if value1 == 0 && value == 0 && value2 == 0 && flag0 == true {
						continue
					}
					if value1 == 0 && value == 0 && value2 == 0 {
						flag0 = true
					}
					textForIndxSlice := strconv.Itoa(sliceForSort[0]) + strconv.Itoa(sliceForSort[1]) + strconv.Itoa(sliceForSort[2])
					for _, valueIndx := range indxSlice {
						if valueIndx == textForIndxSlice {
							flag = true
							break
						}
					}
					if flag == true {
						continue
					}
					indxSlice = append(indxSlice, textForIndxSlice)
					sliceForSort = []int{value, value1, value2}
					sort.Ints(sliceForSort)
					result = append(result, sliceForSort[0])
					result = append(result, sliceForSort[1])
					result = append(result, sliceForSort[2])
				}
			}
		}
		//	result = append(result, intValue )
	}
	return result, nil
}

func main() {
	// Напишите ваш код здесь. Разбейте на функции, если считаете нужным.
	input := getInput()
	IntSlice, err := converString(input)
	if err != nil {
		return
	}
	resultSlice, err := sortSlice(IntSlice)
	if err != nil {
		return
	}
	for indx, value := range resultSlice {
		if (indx+1)%3 == 0 {
			fmt.Print(value, " ")
			continue
		}
		fmt.Print(value, ",")
	}
	//fmt.Println("RESULT")
}
