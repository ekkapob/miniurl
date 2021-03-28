package utils

import "testing"

func TestGetIntFromMap(t *testing.T) {
	t1 := map[string]string{"data": "123"}
	v, err := GetIntFromMap(t1, "data")
	if err != nil || v != 123 {
		t.Fatalf("expect %v but got %v", 123, v)
	}

	t2 := map[string]string{"data": "abc"}
	v, err = GetIntFromMap(t2, "data")
	if err == nil {
		t.Fatalf("expect to have error but not")
	}
}

func TestGetInt(t *testing.T) {
	defaultValue := 10
	tests := map[string]int{
		"a":  10,
		"20": 20,
		"":   10,
	}

	for k, v := range tests {
		r := GetInt(k, defaultValue)
		if GetInt(k, defaultValue) != v {
			t.Fatalf("expect %v but got %v", v, r)
		}
	}
}
