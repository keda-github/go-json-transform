package transform

import (
	"fmt"
	"strings"
)

// Transformer handles JSON transformations
type Transformer struct{}

// New creates a new Transformer instance
func New() *Transformer {
	return &Transformer{}
}

// Transform applies the mapping rules to the input data
func (t *Transformer) Transform(input map[string]interface{}, mapping Mapping) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	for outKey, mapRule := range mapping.Map {
		value, err := t.processRule(input, mapRule)
		if err != nil {
			return nil, fmt.Errorf("error processing key %s: %w", outKey, err)
		}
		result[outKey] = value
	}

	return result, nil
}

// processRule handles different types of mapping rules
func (t *Transformer) processRule(input map[string]interface{}, rule interface{}) (interface{}, error) {
	switch r := rule.(type) {
	case string:
		return t.getValueFromPath(input, r)
	case Field:
		value, err := t.getValueFromPath(input, r.Path)
		if err != nil || value == nil {
			return r.DefaultValue, nil
		}
		return value, nil
	case ArrayMapping:
		return t.processArrayMapping(input, r)
	case Operation:
		return t.processOperation(input, r)
	default:
		return nil, fmt.Errorf("unsupported mapping rule type: %T", rule)
	}
}

// getValueFromPath retrieves a value from nested structure using dot notation
func (t *Transformer) getValueFromPath(input map[string]interface{}, path string) (interface{}, error) {
	parts := strings.Split(path, ".")
	current := input

	for i, part := range parts {
		if i == len(parts)-1 {
			return current[part], nil
		}

		next, ok := current[part]
		if !ok {
			return nil, nil
		}

		current, ok = next.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid path: %s is not an object", part)
		}
	}

	return nil, nil
}

// processArrayMapping handles array transformations
func (t *Transformer) processArrayMapping(input map[string]interface{}, mapping ArrayMapping) (interface{}, error) {
	value, err := t.getValueFromPath(input, mapping.Path)
	if err != nil {
		return nil, err
	}

	arr, ok := value.([]interface{})
	if !ok {
		return nil, fmt.Errorf("value at path %s is not an array", mapping.Path)
	}

	result := make([]map[string]interface{}, len(arr))
	for i, item := range arr {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("array item at index %d is not an object", i)
		}

		transformed, err := t.Transform(itemMap, Mapping{Map: mapping.Mapping})
		if err != nil {
			return nil, err
		}
		result[i] = transformed
	}

	return result, nil
}

// processOperation applies custom operations on multiple fields
func (t *Transformer) processOperation(input map[string]interface{}, op Operation) (interface{}, error) {
	values := make([]interface{}, len(op.Fields))
	for i, field := range op.Fields {
		value, err := t.getValueFromPath(input, field)
		if err != nil {
			return nil, err
		}
		values[i] = value
	}

	return op.Operation(values...), nil
} 