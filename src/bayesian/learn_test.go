package bayesian

import . "testing"

func TestTrim(t *T) {
	at := trim("@@@")
	if len(at) != 0 {
		t.Errorf("trimming '@@@' resulted in %v", at)
	}
}
