package todo

import (
	"encoding/json"
	"os"
)

type Item struct{
	Text string
}

func SaveItems(filename string, items []Item) error {

	b,err := json.Marshal(items)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}