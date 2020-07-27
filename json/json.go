// Package json provides helpful methods for working with generic JSON, represented as nested map. Generic JSON can be used as embedded properties in complex JSON structures, leaving generic parts of the JSON as is
package json

import (
	"strings"
)

// Generic represents JSON as arbitary nested map
type Generic map[string]interface{}

// IsKey check if key is present in the generic JSON, dot notation for key is supported
func (g *Generic) IsKey(key string) bool {
	return g.Value(key) != nil
}

// Bool retrieves boolean value by key, dot notation for key is supported
func (g *Generic) Bool(key string) bool {
	if v, ok := g.Value(key).(bool); ok {
		return v
	}

	return false
}

// Float retrieves float value by key, dot notation for key is supported
func (g *Generic) Float(key string) float64 {
	if v, ok := g.Value(key).(float64); ok {
		return v
	}

	return 0.0
}

// String retrieves string value by key, dot notation for key is supported
func (g *Generic) String(key string) string {
	if v, ok := g.Value(key).(string); ok {
		return v
	}

	return ""
}

// Int retrieves int value by key, dot notation for key is supported
func (g *Generic) Int(key string) int {
	return int(g.Float(key))
}

// Value retrieves valus of unknown type, caller should cast it accordingly
func (g *Generic) Value(key string) interface{} {
	return deepValue(*g, strings.Split(key, "."))
}

// deepValue recursively fetches values from a nested map
func deepValue(g Generic, key []string) interface{} {
	value, ok := g[key[0]]
	if !ok {
		return nil
	}

	if len(key) == 1 {
		return value
	}

	switch value.(type) {
	default:
		return nil
	case map[string]interface{}:
		return deepValue(value.(map[string]interface{}), key[1:])
	}
}

