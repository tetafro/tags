package tags

import "reflect"

// Parse parses struct tags and returns a map of parsed values.
func Parse(obj interface{}, tagName string) (tags []string, values []interface{}) {
	if obj == nil {
		return nil, nil
	}

	typ := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj)

	// Get actual data from pointer
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}

	// Check if object is a struct
	if typ.Kind() != reflect.Struct {
		return nil, nil
	}

	// Iterate over fields
	for i := 0; i < typ.NumField(); i++ {
		tf := typ.Field(i)

		// Skip unexported field
		if tf.PkgPath != "" {
			continue
		}

		// Skip not tagged fields
		tag := tf.Tag.Get(tagName)
		if tag == "" || tag == "-" {
			continue
		}

		tags = append(tags, tag)
		values = append(values, val.Field(i).Interface())
	}
	return tags, values
}
