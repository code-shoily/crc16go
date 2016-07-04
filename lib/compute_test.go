package lib

import "testing"

func ComputeTest(t *testing.T) {
	testData := [][]uint16{
		{0x05, 0x01, 0x00, 0x01},
		{0x0D, 0x01, 0x01, 0x23, 0x45, 0x67, 0x89, 0x01, 0x23, 0x45, 0x00, 0x01},
	}

	expectedData := []uint16{0xd9dc, 0x8cdd}

	for i, v := range testData {
		found := Compute(v)
		expected := expectedData[i]

		if found != expected {
			t.Errorf("FOUND: %d\tEXPECTED: %d\n", found, expected)
		}
	}
}
