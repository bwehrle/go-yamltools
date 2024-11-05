package go_yamltools

import (
	"testing"

	"gopkg.in/yaml.v2"
)

func accumulator(v string, state *[]string) bool {
	*state = append(*state, v)
	return true
}

func TestTraverseMapSlice(t *testing.T) {
	tests := []struct {
		name      string
		tree      yaml.MapSlice
		processor NodeValueProcessor
		expected  []string
	}{
		{"Example tree",
			yaml.MapSlice{
				{Key: "key1", Value: yaml.MapSlice{
					{Key: "key1.1", Value: "value1.1"},
					{Key: "key1.2", Value: "value1.2"},
				}},
				{Key: "key2", Value: "value2"},
			},
			accumulator,
			[]string{"value1.1", "value1.2", "value2"},
		},
	}
	for _, tt := range tests {
		state := make([]string, 0, 100)
		t.Run(tt.name, func(t *testing.T) {
			TraverseMapSlice(tt.tree, &state, tt.processor)
		})
		if len(state) != len(tt.expected) {
			t.Errorf("TraverseMapSlice() = %v, want %v", state, tt.expected)
		}
		for i, v := range state {
			if v != tt.expected[i] {
				t.Errorf("TraverseMapSlice() = %v, want %v", state, tt.expected)
			}
		}
	}
}
