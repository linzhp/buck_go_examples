package goTestDeps

import "testing"

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

func TestUnmarshall(t *testing.T) {
	t.Log("Given the need to test unmarshalling data")
	result := T{}
	err := Unmarshall([]byte(data), &result)
	if err == nil {
		t.Log("Success!")
	} else {
		t.Errorf("Error: %v", err)
	}
}
