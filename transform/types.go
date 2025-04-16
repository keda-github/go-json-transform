package transform

// Mapping defines the structure for transformation rules
type Mapping struct {
    Map map[string]interface{}
}

// Field represents a field mapping with additional options
type Field struct {
    Path         string
    DefaultValue interface{}
}

// ArrayMapping represents mapping rules for array transformation
type ArrayMapping struct {
    Path    string
    Mapping map[string]interface{}
}

// Operation represents a custom operation on multiple fields
type Operation struct {
    Fields    []string
    Operation func(...interface{}) interface{}
} 