package maps

import (
	"errors"
)

type Dictionary map[string]string

var (
	wordMissingError  = errors.New("Word not in dictionary.")
	missingValueError = errors.New("No definition for word.")
	existingWordError = errors.New("Cannot overwrite existing definition for word.")
)

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]

	if !exists {
		return "", missingValueError
	}

	return value, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, exists := d[word]
	if exists {
		return existingWordError
	}

	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, exists := d[word]
	if !exists {
		return wordMissingError
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
