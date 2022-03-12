package maps

const (
	ErrKeyNotFound = DictionaryErr("key not found in the dictionary")
	ErrKeyExists   = DictionaryErr("key already exists in the dictionary")
)

type StringDict map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d StringDict) Search(key string) (string, error) {
	value, ok := d[key]

	if !ok {
		return "", ErrKeyNotFound
	}

	return value, nil
}

func (d StringDict) Add(key, value string) error {
	_, err := d.Search(key)

	if err == nil {
		return ErrKeyExists
	}

	d[key] = value
	return nil
}

func (d StringDict) Update(key, value string) error {
	_, err := d.Search(key)

	if err == ErrKeyNotFound {
		return err
	}

	d[key] = value
	return nil
}

func (d StringDict) Delete(key string) error {
	_, err := d.Search(key)

	if err == ErrKeyNotFound {
		return err
	}

	delete(d, key)
	return nil
}
