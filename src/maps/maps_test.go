package maps

import (
	"testing"
)

var dictionary = Dictionary{"test": "this is just a test"}

func TestSearch(t *testing.T) {
	t.Run("search map - key test found", func(t *testing.T) {
		key := "test"
		got, _ := dictionary.Search(key)
		want := "this is just a test"
		assertString(t, got, want)
	})

	t.Run("search map - key test2 not found", func(t *testing.T) {
		key := "test2"
		_, err := dictionary.Search(key)
		if err == nil {
			t.Error("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add map - does not found key", func(t *testing.T) {
		newDic := Dictionary{}
		key := "test2"
		value := "this is just a test2"
		err := newDic.Add(key, value)

		assertError(t, err, nil)
		assertDefinition(t, newDic, key, value)
	})

	t.Run("add map - found key", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		newDic := Dictionary{key: value}

		err := newDic.Add(key, value)

		assertError(t, err, ErrKeyExist)
		assertDefinition(t, newDic, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update map - key found", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		newDic := Dictionary{key: value}
		newValue := "new key value"

		err := newDic.Update(key, newValue)

		assertError(t, err, nil)
		assertDefinition(t, newDic, key, newValue)
	})

	t.Run("update map - key does not found", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		newDic := Dictionary{key: value}

		newKey := "test2"
		newValue := "new key value"

		err := newDic.Update(newKey, newValue)
		assertError(t, err, ErrKeyDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete map - key found", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		newDic := Dictionary{key: value}
		newDic.Delete(key)

		_, err := newDic.Search(key)
		assertError(t, err, ErrNotFound)
	})

	t.Run("delete map - key does not found", func(t *testing.T) {
		key := "test"
		newDic := Dictionary{}
		newDic.Delete(key)

		_, err := newDic.Search(key)
		assertError(t, err, ErrNotFound)
	})
}
