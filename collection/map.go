package collection

func IsExistInMap(mapping map[interface{}]interface{}, key interface{}) bool {
	_, ok := mapping[key]
	return ok
}

func WhenExistInMap(mapping map[interface{}]interface{}, key interface{}, run func(interface{}) error) error {
	if val, ok := mapping[key]; ok {
		return run(val)
	}
	return nil
}
