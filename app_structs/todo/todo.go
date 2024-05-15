package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type todo struct {
	Text string `json:"text"` /* TAG FOR THE JSON KEY */

}

func New(Text string) (todo, error) {

	if Text == "" {
		return todo{}, errors.New("todo text cannot be empty")
	}

	return todo{
		Text: Text,
	}, nil
}

func (todo todo) Display() {
	fmt.Println(todo.Text)
}

func (todo todo) Save() error {

	newfileName := "todo.json"
	jsonData, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(newfileName, jsonData, 0644)
}
