package bifid

import (
	"strings"
)

var slice = [5][5]string{
		{"R","A","N","C","H"},
		{"V","W","X","Y","Z"},
		{"O","B","D","E","F"},
		{"P","Q","U","S","T"},
		{"G","I","K","L","M"},
	}

func key(letter string) (int, int) {	
	positions := make(map[string]struct{i, j int})

	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			positions[slice[i][j]] = struct{i, j int}{i, j}
		}
	}

	position, _ := positions[letter]
	return position.i, position.j
}

func Encrypt(message string) string {
	message = strings.ToUpper(message)
	var encrypted string
	var arr_i []int
	var arr_j []int
	var arr []int

	for _, char := range message {
		i, j := key(string(char))
		arr_i = append(arr_i, i)
		arr_j = append(arr_j, j)
	}
	arr = append(arr, arr_i...)
	arr = append(arr, arr_j...)

	for i := 0; i < len(arr); i += 2 {
		letter := slice[arr[i]][arr[i + 1]]
		encrypted += letter
	}

	return encrypted
}

func Decrypt(message string) string {
	message = strings.ToUpper(message)
	var decrypted string
	var arr1 []int
	var arr2 []int

	for _, char := range message {
		i, j := key(string(char))
		arr1 = append(arr1, i, j)
	}
	lenSplit := len(arr1) / 2
	for j := 0; j < lenSplit; j++ {
		arr2 = append(arr2, arr1[j], arr1[j + lenSplit])
	} 

	for i := 0; i < len(arr2); i += 2 {
		letter := slice[arr2[i]][arr2[i + 1]]
		decrypted += letter
	}

	return decrypted
}