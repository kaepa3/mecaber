package mecaber

import (
	"fmt"
	"testing"
)

func TestNewing(t *testing.T) {
	err, m := CreateNew()
	if err != nil {
		t.Error(err.Error())
	}
	if m == nil {
		t.Error("nil mecab")
	}
	done := make(chan struct{})
	errCh, ch := m.ParseToNode("吾輩は猫である", done)

LOOP:
	for {
		select {
		case text := <-ch:
			fmt.Println(text)
		case err := <-errCh:
			fmt.Println(err)
			break LOOP
		case <-done:
			break LOOP
		}
	}

}
