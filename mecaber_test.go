package mecaber

import (
	"fmt"
	"testing"
)

func TestNewing(t *testing.T) {
	err, m := CreateNew("")
	if err != nil {
		t.Error(err.Error())
	}
	if m == nil {
		t.Error("nil mecab")
	}
	done := make(chan struct{})
	errCh, ch := m.ParseToNode("8月3日に放送された「中居正広の金曜日の スマイルたちへ」(TBS系)で、1日たった5分でぽっこりおなかを解消するというダイエット方法を紹 介。キンタロー。のダイエットにも密着。",
		done)
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
