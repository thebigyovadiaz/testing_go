package maps

import (
	"errors"
)

var (
	ErrNotFound        = DictionaryErr("key not found")
	ErrKeyExist        = DictionaryErr("key is already found")
	ErrKeyDoesNotExist = DictionaryErr("key not exists")
)

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch {
	case errors.Is(err, ErrNotFound):
		d[key] = value
	case err == nil:
		return ErrKeyExist
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch {
	case errors.Is(err, ErrNotFound):
		return ErrKeyDoesNotExist
	case err == nil:
		d[key] = value
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) {
	_, err := d.Search(key)

	if err == nil {
		delete(d, key)
	}
}
