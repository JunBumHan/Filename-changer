package main

import (
	replDir "Filename-changer/replaceDir"
	"fmt"
	"io/ioutil"
)

var filePath replDir.DirPath

func main() {
	var userInput string

	fmt.Println("[FILE CHANGER]")
	fmt.Println("Please enter the file path")

	_, err := fmt.Scan(&filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	replaceDir := replDir.NewReplaceDir(filePath)

	files, err := ioutil.ReadDir(string(replaceDir.TargetDir))
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, file := range files {
		fmt.Printf("index[%d] : %s\n", i, file.Name())

	}

	fmt.Println("Is there any file you don't want to change?[N/Y]")
	_, err = fmt.Scan(&userInput)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch userInput {
	case "Y":
		fmt.Println("Enter file name")
		_, err = fmt.Scan(&replaceDir.NoneDirName)
		if err != nil {
			fmt.Println(err)
			return
		}
		replaceDir.NoneDirSlice = append(replaceDir.NoneDirSlice, replaceDir.NoneDirName)

		fmt.Println("Not changed file list")
		for i, notChanged := range replaceDir.NoneDirSlice {
			fmt.Printf("index[%d] : %s\n", i, notChanged)
		}
		fallthrough
	case "N":
		// var fileFullName string
		n := 0
		i := 0

		fmt.Println("OK, Let's Go")
		for _, file := range files {
			fileName := file.Name()
			// fileFullName = string(replaceDir.TargetDir) + "/" + fileName

			for i = 0; i < len(fileName); i++ {
				if fileName[i] == byte('-') {
					break
				}
			}
			fileName = removeStringBeforeIndex(fileName, i+1)
			fmt.Println(fileName)
			// res := fmt.Sprintf("%02d-%s", n, fileName)
			// os.Rename(fileFullName, res)
			n++
		}

	default:
		fmt.Println("Invalid input")
		return
	}

}

func removeStringBeforeIndex(str string, index int) string {
	return str[index:]
}
