package util

func ParseInterfaceMap(data interface{}, keys []string) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		if value, ok := v[keys[0]]; ok {
			if len(keys) == 1 {
				return v[keys[0]]
			} else {
				return ParseInterfaceMap(value, keys[1:])
			}
		}
	case []interface{}:
		for _, value := range v {
			if len(keys) == 1 {
				return value.(map[string]interface{})[keys[0]]
			} else {
				return ParseInterfaceMap(value, keys[1:])
			}
		}
	}
	return nil
}
