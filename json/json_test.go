package json

import (
	"encoding/json"
	"testing"
)

func TestValues(t *testing.T) {
	input := `{
		"bool": true,
		"int": 1,
		"string": "test",
		"float": 123.45,
		"map": {
			"bool": true,
			"int": 2,
			"string": "test2",
			"float": 234.56,
			"deep": {
				"value": 42
			}
		}
	}`

	var j Generic
	if err := json.Unmarshal([]byte(input), &j); err != nil {
		t.Errorf("failed to unmarshal input '%s': %s", input, err)
	}

	if j.Bool("bool") != true {
		t.Errorf("failed to retrieve 'bool' value")
	}
	if j.Int("int") != 1 {
		t.Errorf("failed to retrieve 'int' value")
	}
	if j.String("string") != "test" {
		t.Errorf("failed to retrieve 'string' value")
	}
	if j.Float("float") != 123.45 {
		t.Errorf("failed to retrieve 'float' value")
	}
	if j.Bool("map.bool") != true {
		t.Errorf("failed to retrieve 'map.bool' value")
	}
	if j.Int("map.int") != 2 {
		t.Errorf("failed to retrieve 'map.int' value")
	}
	if j.String("map.string") != "test2" {
		t.Errorf("failed to retrieve 'map.string' value")
	}
	if j.Float("map.float") != 234.56 {
		t.Errorf("failed to retrieve 'map.float' value")
	}
	if j.Int("map.deep.value") != 42 {
		t.Errorf("failed to retrieve 'map.deep.value' value")
	}
}

