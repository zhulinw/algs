package utils

import (
	"reflect"
	"testing"
)

func Equal(t *testing.T, want, got interface{}) {
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, but got %v", want, got)
	}
}

func AssertEqual(want, got interface{}) bool {
	if reflect.DeepEqual(want, got) {
		return true
	}
	return false
}
