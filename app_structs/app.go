package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"structs/note"
	"structs/todo"
)

type saver interface {
	Save() error
}

// type outputable interface {
// 	Save() error
// 	Display()
// }

/*  OR  */

type outputable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}
	OutputData(userNote)

	/* TODO ACTIONS */
	todoText := getUserInput("Todo Text: ")
	todoData, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	OutputData(todoData)

}
func getUserInput(prompt string) string {
	fmt.Print(prompt)
	// fmt.Scanln(&val)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return ""
	}

	text = strings.TrimSpace(text)
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func getNoteData() (title string, content string) {
	title = getUserInput("Note Title: ")
	content = getUserInput("Note Content: ")

	return title, content
}

func SaveData(data saver) error {
	err := data.Save()

	if err != nil {
		return err
	}

	fmt.Println("Saved Successfully!")
	return nil
}

func OutputData(data outputable) error {
	data.Display()

	err := SaveData(data)

	if err != nil {
		fmt.Println("Failed to save!")
		return err
	}

	return nil
}
