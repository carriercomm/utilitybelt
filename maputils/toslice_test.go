package maps

import (
	"testing"
)

func TestStringInterfaceToSlices(t *testing.T) {
	tests := []struct {
		name           string
		m              map[string]interface{}
		expectedKeys   []string
		expectedValues []interface{}
		expectedErr    string
	}{
		{name: "empty slice", expectedErr: ""},
		{
			name: "slice with 2 keys",
			m: map[string]interface{}{
				"hello": "goodbye",
				"salut": "au revoir",
			},
			expectedKeys:   []string{"hello", "salut"},
			expectedValues: []interface{}{"goodbye", "au revoir"},
			expectedErr:    "",
		},
	}

	for _, test := range tests {
		keys, values := StringInterfaceToSlices(test.m)
		if len(keys) != len(values) {
			t.Errorf("Mismatched slices: key length was %s; values length was %s")
			continue
		}
		for i, key := range keys {
			val, ok := test.m[key]
			// first check to see if what we received is in original
			if !ok {
				t.Errorf("Key %s, which was extracted from the passed map was not found in it.", key)
				continue
			}
			if val != values[i] {
				t.Errorf("Unexpected value extracted from map for %s: %s received,  %s expected", key, val, values[i])
			}
			var found bool
			// then check to see if it is in expected
			for _, k := range test.expectedKeys {
				if key == k {
					found = true
					continue
				}
			}
			if !found {
				t.Errorf("%s not found in the expected keys", key)
			}
			found = false
			for _, v := range test.expectedValues {
				if val == v {
					found = true
					continue
				}
			}
			if !found {
				t.Errorf("%s not found in the expected values", val)
			}
		}
	}
}

func TestStringStringToSlices(t *testing.T) {
	tests := []struct {
		name           string
		m              map[string]string
		expectedKeys   []string
		expectedValues []string
		expectedErr    string
	}{
		{name: "empty slice", expectedErr: ""},
		{
			name: "slice with 2 keys",
			m: map[string]string{
				"hello": "goodbye",
				"salut": "au revoir",
			},
			expectedKeys:   []string{"hello", "salut"},
			expectedValues: []string{"goodbye", "au revoir"},
			expectedErr:    "",
		},
	}

	for _, test := range tests {
		keys, values := StringStringToSlices(test.m)
		if len(keys) != len(values) {
			t.Errorf("Mismatched slices: key length was %s; values length was %s")
			continue
		}
		for i, key := range keys {
			val, ok := test.m[key]
			// first check to see if what we received is in original
			if !ok {
				t.Errorf("Key %s, which was extracted from the passed map was not found in it.", key)
				continue
			}
			if val != values[i] {
				t.Errorf("Unexpected value extracted from map for %s: %s received,  %s expected", key, val, values[i])
			}
			var found bool
			// then check to see if it is in expected
			for _, k := range test.expectedKeys {
				if key == k {
					found = true
					continue
				}
			}
			if !found {
				t.Errorf("%s not found in the expected keys", key)
			}
			found = false
			for _, v := range test.expectedValues {
				if val == v {
					found = true
					continue
				}
			}
			if !found {
				t.Errorf("%s not found in the expected values", val)
			}
		}
	}
}

func TestStringIntToSlices(t *testing.T) {
	tests := []struct {
		name           string
		m              map[string]int
		expectedKeys   []string
		expectedValues []int
		expectedErr    string
	}{
		{name: "empty slice", expectedErr: ""},
		{
			name: "slice with 2 keys",
			m: map[string]int{
				"a": 1,
				"b": 2,
			},
			expectedKeys:   []string{"a", "b"},
			expectedValues: []int{1, 2},
			expectedErr:    "",
		},
	}

	for _, test := range tests {
		keys, values := StringIntToSlices(test.m)
		if len(keys) != len(values) {
			t.Errorf("Mismatched slices: key length was %s; values length was %s")
			continue
		}
		for i, key := range keys {
			val, ok := test.m[key]
			// first check to see if what we received is in original
			if !ok {
				t.Errorf("Key %s, which was extracted from the passed map was not found in it.", key)
				continue
			}
			if val != values[i] {
				t.Errorf("Unexpected value extracted from map for %s: %s received,  %s expected", key, val, values[i])
			}
			var found bool
			// then check to see if it is in expected
			for _, k := range test.expectedKeys {
				if key == k {
					found = true
					continue
				}
			}
			if !found {
				t.Errorf("%s not found in the expected keys", key)
			}
			found = false
			for _, v := range test.expectedValues {
				if val == v {
					found = true
					continue
				}
			}
			if !found {
				t.Errorf("%s not found in the expected values", val)
			}
		}
	}
}

func TestStringBoolToSlices(t *testing.T) {
	tests := []struct {
		name           string
		m              map[string]bool
		expectedKeys   []string
		expectedValues []bool
		expectedErr    string
	}{
		{name: "empty slice", expectedErr: ""},
		{
			name: "slice with 2 keys",
			m: map[string]bool{
				"a": true,
				"b": false,
			},
			expectedKeys:   []string{"a", "b"},
			expectedValues: []bool{true, false},
			expectedErr:    "",
		},
	}

	for _, test := range tests {
		keys, values := StringBoolToSlices(test.m)
		if len(keys) != len(values) {
			t.Errorf("Mismatched slices: key length was %s; values length was %s")
			continue
		}
		for i, key := range keys {
			val, ok := test.m[key]
			// first check to see if what we received is in original
			if !ok {
				t.Errorf("Key %s, which was extracted from the passed map was not found in it.", key)
				continue
			}
			if val != values[i] {
				t.Errorf("Unexpected value extracted from map for %s: %s received,  %s expected", key, val, values[i])
			}
			var found bool
			// then check to see if it is in expected
			for _, k := range test.expectedKeys {
				if key == k {
					found = true
					continue
				}
			}
			if !found {
				t.Errorf("%s not found in the expected keys", key)
			}
			found = false
			for _, v := range test.expectedValues {
				if val == v {
					found = true
					continue
				}
			}
			if !found {
				t.Errorf("%s not found in the expected values", val)
			}
		}
	}
}
