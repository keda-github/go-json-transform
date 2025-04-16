package transform

import (
	"reflect"
	"testing"
)

func TestBasicTransform(t *testing.T) {
	input := map[string]interface{}{
		"name": "John Doe",
		"age":  30,
		"contact": map[string]interface{}{
			"email": "john@example.com",
		},
	}

	mapping := Mapping{
		Map: map[string]interface{}{
			"fullName":     "name",
			"userAge":      "age",
			"emailAddress": "contact.email",
		},
	}

	expected := map[string]interface{}{
		"fullName":     "John Doe",
		"userAge":      30,
		"emailAddress": "john@example.com",
	}

	transformer := New()
	result, err := transformer.Transform(input, mapping)

	if err != nil {
		t.Errorf("Transform failed: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestArrayTransform(t *testing.T) {
	input := map[string]interface{}{
		"users": []interface{}{
			map[string]interface{}{"name": "John", "age": 30},
			map[string]interface{}{"name": "Jane", "age": 25},
		},
	}

	mapping := Mapping{
		Map: map[string]interface{}{
			"people": ArrayMapping{
				Path: "users",
				Mapping: map[string]interface{}{
					"fullName": "name",
					"userAge":  "age",
				},
			},
		},
	}

	expected := map[string]interface{}{
		"people": []map[string]interface{}{
			{"fullName": "John", "userAge": 30},
			{"fullName": "Jane", "userAge": 25},
		},
	}

	transformer := New()
	result, err := transformer.Transform(input, mapping)

	if err != nil {
		t.Errorf("Transform failed: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
} 