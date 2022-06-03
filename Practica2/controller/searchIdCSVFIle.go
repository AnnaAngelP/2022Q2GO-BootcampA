package controller

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	entities "main.com/entities"
)

func HandlerSearchID(w http.ResponseWriter, r *http.Request) {
	file := "./pokemonFile.csv"
	IdToSearch := 2

	id, err := SearchID(file, IdToSearch)
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	fmt.Printf("The JSON response as a string\n\n")
	fmt.Println(string(id))

	_, e := w.Write(id)
	if e != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Successfully completed the task")
}

func SearchID(file string, id int) ([]byte, error) {
	csvLines, e := ReadData(file)

	if e != nil {
		fmt.Print("Error trying to get the info in CSV file ")
		return nil, e
	}
	fmt.Println("Successfully got the info in CSV file")

	var dataID []entities.Data

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

		if idInt == id {
			dataID = append(dataID, d)
			fmt.Println(dataID)
		}
	}
	fmt.Println("Datos: ", dataID)

	byteArray, err := json.MarshalIndent(dataID, "", "  ")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return byteArray, e
}

func ReadData(fileName string) ([][]string, error) {
	f, err := os.Open(fileName) // open file

	if err != nil {
		return [][]string{}, err
	}

	fmt.Println("Successfully opened the file")
	defer f.Close()               // close the file at the end of the program
	csvReader := csv.NewReader(f) // read csv values using csv.Reader

	// skip first line
	//if _, err := csvReader.Read(); err != nil {
	//	return [][]string{}, err
	//}

	csvLines, err := csvReader.ReadAll()

	if err != nil {
		return nil, err
	}
	return csvLines, nil
}
