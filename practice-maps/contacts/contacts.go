package contacts

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errNameExists = errors.New("That name already exists")

// Search for name
func (d Dictionary) Search(word string) (string, error) {
	mobile, exists := d[word]
	if exists {
		return mobile, nil
	}
	return "", errNotFound
}

// Add to Dictionary
func (d Dictionary) Add(name, mobile string) error {
	_, err := d.Search(name)
	switch err {
	case errNotFound:
		d[name] = mobile
	case nil:
		return errNameExists
	}
	return nil
}

// Delete from Dictionary
func (d Dictionary) Delete(name string) error {
	_, err := d.Search(name)
	if err != nil {
		return errNotFound
	}
	delete(d, name)
	return nil
}
