package deck

import "testing"

func Test_newDeck(t *testing.T) {
	d := NewDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 20 but got %v", len(d))
	}
}