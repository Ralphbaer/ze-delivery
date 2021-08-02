package common

import "encoding/json"

// StructToMapExpansive converts a struct (interface) to a map[string]interface{} by using json marshaling (struct) and than
// unmarshaling it back to map[string]interface{}
func StructToMapExpansive(s interface{}) (map[string]interface{}, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}

// MapToStructExpansive converts a map[string]interface{} by using json marshaling and than
// unmarshaling it back to struct
func MapToStructExpansive(m map[string]interface{}, s interface{}) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, s); err != nil {
		return err
	}
	return nil
}
