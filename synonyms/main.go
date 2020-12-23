package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/leonistor/domainfinder/thesaurus"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	apiKey := goDotEnvVariable("BHT_APIKEY")
	thesaurus := &thesaurus.BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalln("Failed when looking for synonyms for "+word, err)
		}
		if len(syns) == 0 {
			log.Fatalln("Could not find any synonyms for " + word)
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
