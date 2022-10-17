package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Data struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Age  int    `json:"Age"`
}

func main() {
	jsonFile, err := os.Open("C:/golang/lesson10/cmd/user.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println(jsonFile)

	defer func() {
		err := jsonFile.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	log.Printf("%#v", jsonFile)

	byteValue, _ := io.ReadAll(jsonFile)
	var users Data
	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(users.Name)
	log.Println(users.Type)
	log.Println(users.Age)

	//2
	file, err := os.Create("read")
	if err != nil {
		log.Println(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	fi, err := file.Stat()
	if err != nil {
		log.Println(err)
	}
	log.Printf("Права доступа к файлу  %v\n", fi.Mode())

	err = file.Chmod(555)
	if err != nil {
		log.Println(err)

	}
	fi, err = file.Stat()
	if err != nil {
		log.Println(err)

	}
	log.Printf("Права доступа к файлу %v\n", fi.Mode())

}
