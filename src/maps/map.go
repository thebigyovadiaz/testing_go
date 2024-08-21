package maps

type Dictionary map[string]string

func (d Dictionary) Search(key string) string {
	value, ok := d[key]
	if !ok {
		return "key does not exist"
	}

	return value
}
