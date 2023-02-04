package maps

import "errors"

/*
A gotcha with maps is that they can be a nil value.
A nil map behaves like an empty map when reading,
but attempts to write to a nil map will cause a runtime panic
Do not insantiate an empty map like so:
var dictionary map[string]string
var dictionary = map[string]string{}
OR
var dictionary = make(map[string]string)
Both approaches create an empty hash map and point dictionary at it.
Which ensures that you will never get a runtime panic
*/

// Map keys can only be comparable types
// so they can be looked up easily

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")
var ErrAlreadyExist = errors.New("This word is already in our dictionary")

func (d Dictionary) Search(word string) (string, error) {
	// maps return 2 values. Second to indicate if it found the value ("ok")
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, ok := d[word]
	if ok {
		return ErrAlreadyExist
	}

	d[word] = definition
	return nil
}
