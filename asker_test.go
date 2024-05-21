package asker_test

import (
	"testing"

	"github.com/AWtnb/go-asker"
)

func TestAsk(t *testing.T) {
	a := asker.Asker{Accept: "y", Reject: "n", Negative: true}
	a.Ask("is it ok?")
	if a.Accepted() {
		t.Log("accepted!")
	} else {
		t.Log("rejected...")
	}
}

func TestAskNonNegative(t *testing.T) {
	a := asker.Asker{Accept: "y", Reject: "n", Negative: false}
	a.Ask("is it ok?")
	if a.Accepted() {
		t.Log("accepted!")
	} else {
		t.Log("rejected...")
	}
}
