package asker

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Asker struct {
	Accept   string
	Reject   string
	Negative bool
	answer   string
}

func (a *Asker) Ask(prompt string) {
	f := "%s (%s/%s)"
	if a.Negative {
		fmt.Printf(f, prompt, a.Accept, strings.ToUpper(a.Reject))
	} else {
		fmt.Printf(f, prompt, strings.ToUpper(a.Accept), a.Reject)
	}
	scn := bufio.NewScanner(os.Stdin)
	scn.Scan()
	a.answer = strings.TrimSpace(scn.Text())
}

func (a Asker) Accepted() bool {
	if len(a.answer) < 1 {
		return !(a.Negative)
	}
	return strings.EqualFold(a.answer, a.Accept)
}
