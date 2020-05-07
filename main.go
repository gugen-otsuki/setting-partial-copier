package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Failed to read from stdin")
		return
	}
	tokens := strings.Split(string(bytes), ".")
	if len(tokens) != 3 {
		log.Fatal("Failed to parse input")
		return
	}
	configBsonString, err := base64.StdEncoding.DecodeString(tokens[1])
	if err != nil {
		log.Fatal("Failed to parse input")
		return
	}

	var parsed map[string]interface{}
	if err = bson.Unmarshal([]byte(configBsonString), &parsed); err != nil {
		log.Fatalf("Failed to parse input: %v\n", err)
		return
	}
	jsonString, err := json.Marshal(parsed)
	if err != nil {
		log.Fatalf("Failed to parse input: %v\n", err)
		return
	}
	fmt.Println(string(jsonString))
}
