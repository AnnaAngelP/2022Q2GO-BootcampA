package controller

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	entity "main.com/entities"
)

func APIHandler(w http.ResponseWriter, r *http.Request) {
	url := "https://pokeapi.co/api/v2/pokedex/kanto/"
	var responseObject entity.Response

	responseData, err := GetDataUrl(url)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		fmt.Println(err)
	}
	// create a file
	csvFile, err := os.Create("pokemonFile.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	// initialize csv writer
	csvwriter := csv.NewWriter(csvFile)

	var Id string
	var NamePoke string
	PokeArray := [][]string{}

	for _, d := range responseObject.PokemonEntries {
		PArray := []string{}
		Id = strconv.Itoa(d.ID)
		NamePoke = d.Species.Name

		PArray = append(PArray, Id, NamePoke)
		PokeArray = append(PokeArray, PArray)
		fmt.Println("array: ", PArray)
		_ = csvwriter.Write(PArray) // write single row
	}

	fmt.Println("Poke array: ", PokeArray)

	csvwriter.Flush()
	csvFile.Close()

	fmt.Println("The task was completed successfully.")
}

func GetDataUrl(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData, err
}
