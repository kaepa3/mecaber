package mecaber

import (
	"fmt"

	"github.com/bluele/mecab-golang"
)

const BOSEOS = "BOS/EOS"

type Mecaber struct {
	m *mecab.MeCab
}

func CreateNew(opt ...string) (error, *Mecaber) {
	o := append(opt, "-Owakati")
	m, err := mecab.New(o...)
	if err != nil {
		return err, nil
	}
	return nil, &Mecaber{m: m}
}

func (me *Mecaber) Destroy() {
	me.m.Destroy()
}

func (me *Mecaber) ParseToNode(text string, done chan struct{}) (<-chan error, <-chan string) {
	ch := make(chan string)
	errCh := make(chan error)
	go func() {
		tg, err := me.m.NewTagger()
		if err != nil {
			errCh <- err
			return
		}
		defer tg.Destroy()

		lt, err := me.m.NewLattice(text)
		if err != nil {
			errCh <- err
			return
		}
		defer lt.Destroy()
		node := tg.ParseToNode(lt)
		send(node, ch)
		close(done)
	}()
	return errCh, ch
}

func send(n *mecab.Node, ch chan string) {
	for {
		ch <- fmt.Sprintf("%s,%s", n.Surface(), n.Feature())
		if n.Next() != nil {
			break
		}
	}
}
