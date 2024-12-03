package Solutions

import (
	"reflect"
	"testing"
)

func TestMovingAverage(t *testing.T) {
	tests := []struct {
		name       string
		size       int
		operations []int
		expected   []float64
	}{
		{
			name:       "test01",
			size:       3,
			operations: []int{1, 10, 3, 5},
			expected:   []float64{1.0, 5.5, 4.666666666666667, 6.0},
		},
		{
			name:       "test02",
			size:       1,
			operations: []int{1, 2, 3},
			expected:   []float64{1.0, 2.0, 3.0},
		},
		{
			name:       "test03",
			size:       2,
			operations: []int{5, 10, -5, 0},
			expected:   []float64{5.0, 7.5, 2.5, -2.5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize the MovingAverage object
			movingAverage := MovingAverageConstructor(tt.size)
			var result []float64

			// Perform operations and store results
			for _, val := range tt.operations {
				result = append(result, movingAverage.Next(val))
			}

			// Compare results with the expected values
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Test %s failed: expected %v, got %v", tt.name, tt.expected, result)
			}
		})
	}
}