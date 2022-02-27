package lexer

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func LoadRules(file string) (Rules, error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	// we initialize our Rules array
	var rules Rules

	err = json.Unmarshal(byteValue, &rules)
	return rules, err
}
