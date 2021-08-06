package zmock

import (
	"fmt"
	"reflect"
)

// FieldValue represents the attributes needed to return a gomock.Matcher based interface
type FieldValue struct {
	FieldName string
	Value     interface{}
}

// FieldValueMatcher returns a matcher that check the value of a field in a struct.
// It implements the gomock.Matcher interface
func FieldValueMatcher(fieldName string, value interface{}) *FieldValue {
	return &FieldValue{
		FieldName: fieldName,
		Value:     value,
	}
}

// Matches returns whether x is a match.
func (s FieldValue) Matches(x interface{}) bool {
	v := reflect.ValueOf(x)
	if reflect.TypeOf(x).Kind() == reflect.Ptr {
		v = reflect.ValueOf(x).Elem()
	}
	val := v.FieldByName(s.FieldName).Interface()
	return reflect.DeepEqual(val, s.Value)
}

// String describes what the matcher matches.
func (s FieldValue) String() string {
	return fmt.Sprintf("%s=%s", s.FieldName, s.Value)
}
