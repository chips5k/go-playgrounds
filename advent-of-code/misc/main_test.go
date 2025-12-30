package main

import "testing"

func TestIsBalanced(t *testing.T) {

	tests := []struct {
		input string
		want  bool
	}{
		{"", true},
		{"(((", false},
		{")))", false},
		{"()()", true},
		{")()(", false},
		{"())()(", false},
		{"(()())()", true},
		{"()(())", true},
		{"((((()))))", true},
		{"((((())))))", false},
		{"(", false},
		{")", false},
		{"(()", false},
		{"())", false},
		{"()(", false},
		{"(( ))", true},
	}

	for _, tt := range tests {
		got := isBalanced(tt.input)
		if got != tt.want {
			t.Errorf("Want: %v, got %v for %s", tt.want, got, tt.input)
		}
	}

}
