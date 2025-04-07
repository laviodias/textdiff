package diff

import (
	"testing"
)

func TestCompareTexts(t *testing.T) {
	tests := []struct {
		a, b     string
		expected string
	}{
		{
			a:        "my name is Josh and I am a developer",
			b:        "my name is John and I am a dev",
			expected: "my name is [-Josh-] {+John+} and I am a [-developer-] {+dev+}",
		},
		{
			a:        "this should be the same",
			b:        "this should be the same",
			expected: "this should be the same",
		},
		{
			a:        "hello big world",
			b:        "hello there",
			expected: "hello [-big world-] {+there+}",
		},
		{
			a:        "this should be removed",
			b:        "this",
			expected: "this [-should be removed-]",
		},
		{
			a:        "this should",
			b:        "this should be added",
			expected: "this should {+be added+}",
		},
	}

	for _, tt := range tests {
		result := CompareTexts(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("CompareTexts(%q, %q) = %q, expected %q", tt.a, tt.b, result, tt.expected)
		}
	}
}
