package test1

import "testing"

// Charles Myers

func TestDouble(t *testing.T) {

	tests := []struct {
		in       int
		expected int
	}{
		{
			in:       23,
			expected: 46,
		},
		{
			in:       1,
			expected: 2,
		},
	}

	for ii, test := range tests {
		rr := DoubleValue(test.in)
		if rr != test.expected {
			t.Errorf("Test %d, expected %d got %d\n", ii, test.expected, rr)
		}
	}

}

// add test for TripleValue at this point

func TestTripple(t *testing.T) {

	tests := []struct {
		in       int
		expected int
	}{
		{
			in:       23,
			expected: 69,
		},
		{
			in:       1,
			expected: 3,
		},
	}

	for ii, test := range tests {
		rr := TrippleValue(test.in)
		if rr != test.expected {
			t.Errorf("Test %d, expected %d got %d\n", ii, test.expected, rr)
		}
	}

}