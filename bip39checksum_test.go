package main

import (
	"testing"
	"reflect"
)


func TestMain(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{"crowd flip sign buyer truly sail balcony amazing spy garbage cool during decide female phone lawn faith rely barrel celery urban bridge width", []string{"word1", "word2", "word3", "word4", "word5", "word6", "word7", "word8", "word9", "word10", "word11"}},
		{"twenty three words mnemonic", []string{"word1", "word2", "word3", "word4", "word5", "word6", "word7", "word8", "word9", "word10", "word11", "word12", "word13", "word14", "word15", "word16", "word17", "word18", "word19", "word20", "word21", "word22", "word23"}},
		{"invalid mnemonic", []string{}},
	}

	for _, test := range tests {
		got := brute(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("brute(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}
