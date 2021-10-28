package utils

import "testing"

func TestContains(t *testing.T) {
	var params []string
	params = append(params, "-v")

	if Contains(params, "something") {
		t.Fatal("Contains function found unexpected result.")
	}

	if !Contains(params, "-v") {
		t.Fatal("Contains function did not find existing field.")
	}

}
