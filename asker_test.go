package asker_test

import (
	"testing"

	"github.com/AWtnb/go-asker"
)

func TestAsk(t *testing.T) {
	a := asker.Asker{Accept: "y", Reject: "n"}
	a.Ask("is it ok?")
	if a.Accepted() {
		t.Log("accepted!")
	} else {
		t.Log("rejected...")
	}
}

func TestAskPositive(t *testing.T) {
	a := asker.Asker{Accept: "y", Reject: "n", Positive: true}
	a.Ask("is it ok?")
	if a.Accepted() {
		t.Log("accepted!")
	} else {
		t.Log("rejected...")
	}
}
