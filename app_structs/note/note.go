package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"` /* TAG FOR THE JSON KEY */
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title string, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("Note title & Content cannot be empty")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content: \n%v \n\n", note.Title, note.Content)
}

func (note Note) Save() error {

	newfileName := fmt.Sprintf("Note_%v", note.Title)
	newfileName = strings.ReplaceAll(newfileName, " ", "_")
	newfileName = strings.ToLower(newfileName) + ".json"

	jsonData, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(newfileName, jsonData, 0644)
}
