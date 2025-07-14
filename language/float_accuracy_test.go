package code

import (
	"math"
	"testing"
)

func TestFloatAccuracy(t *testing.T) {
	tests := []struct {
		name     string
		input    float32
		expected float32
	}{
		{"Zero", 0.0, 0.0},
		{"Positive Float", 3.14159, 3.14159},
		{"Negative Float", -2.71828, -2.71828},
		{"transform Float", float32(math.Round(0.992*100*10) / 10), 99.2},
		{"transform Float Calculate", float32(math.Round(float64(0.092)*100*10)/10) - float32(math.Round(float64(0.112)*100*10)/10), float32(-2.0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
