package todo

import (
	"encoding/json"
	"errors"
	"os"
)

type List struct {
	items []Item
}

func NewList() *List {
	return &List{}
}

func (l *List) Add(description string) {
	l.items = append(l.items, Item{Description: description})
}

func (l *List) Items() []string {
	var descriptions []string
	for _, item := range l.items {
		descriptions = append(descriptions, item.Description)
	}
	return descriptions
}

func (l *List) Remove(index int) error {
	if index < 0 || index >= len(l.items) {
		return errors.New("index out of range")
	}
	l.items = append(l.items[:index], l.items[index+1:]...)
	return nil
}

func (l *List) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(l.items)
}

func (l *List) Load(filename string) error {
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		// Create an empty file if it does not exist
		emptyFile, err := os.Create(filename)
		if err != nil {
			return err
		}
		emptyFile.Close()
		return nil
	} else if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&l.items)
}
