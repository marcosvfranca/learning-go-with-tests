package main

var (
	ErrorNonExistentWord         = ErrDictionary("word not found")
	ErrorWordExists              = ErrDictionary("existing word")
	ErrorNonExistentWordToUpdate = ErrDictionary("word not found to update")
)

type Dictionary map[string]string

type ErrDictionary string

func (d Dictionary) Search(str string) (string, error) {
	definition, exists := d[str]
	if !exists {
		return "", ErrorNonExistentWord
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNonExistentWord:
		d[word] = definition
	case nil:
		return ErrorWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNonExistentWord:
		return ErrorNonExistentWordToUpdate
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Destroy(word string) {
	delete(d, word)
}

func (e ErrDictionary) Error() string {
	return string(e)
}
