package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	entities "main.com/entities"
)

func HandlerReadFile(w http.ResponseWriter, r *http.Request) {
	file := "./pokemonFile.csv"

	information, e := GetData(file)
	if e != nil {
		fmt.Println("Error: ", e)
		os.Exit(1)
	}

	fmt.Printf("The JSON response as a string\n\n")
	fmt.Println(string(information))

	_, err := w.Write(information)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("Successfully completed the task")
}

func GetData(file string) ([]byte, error) {
	csvLines, e := ReadData(file)

	if e != nil {
		fmt.Print("Error trying to get the info in CSV file ")
		return nil, e
	}
	fmt.Println("Successfully got the info in CSV file")

	var dataInfo []entities.Data
	for _, line := range csvLines {
		idInt, e := strconv.Atoi(line[0])

		if e != nil {
			fmt.Printf("%T \n %v", idInt, idInt)
			return nil, e
		}
		d := entities.Data{
			ID:   idInt,
			Name: line[1],
		}
		dataInfo = append(dataInfo, d)
	}

	byteArray, err := json.MarshalIndent(dataInfo, "", " ")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return byteArray, e
}
