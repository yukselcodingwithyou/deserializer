package main

import (
	"fmt"
	deserializer "github.com/Trendyol/proteo/deserializer"
	"io/ioutil"
	"log"
)

func main() {
	s := struct {
		Name       string `yaml:"name"`
		Profession string `yaml:"profession"`
	}{}

	d := deserializer.New(deserializer.YAML)
	err := d.Deserialize(readFile(), &s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
}

func readFile() []byte {
	content, err := ioutil.ReadFile("example.yml")

	if err != nil {
		log.Fatal(err)
	}
	return content
}
