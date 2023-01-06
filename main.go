package main

import (
	"fmt"
	"strings"

	"github.com/bluele/mecab-golang"
)

const BOSEOS = "BOS/EOS"

func parseToNode(m *mecab.MeCab, text string) error {

	tg, err := m.NewTagger()
	if err != nil {
		return err
	}
	defer tg.Destroy()

	lt, err := m.NewLattice(text)
	if err != nil {
		return err
	}
	defer lt.Destroy()

	node := tg.ParseToNode(lt)
	for {
		features := strings.Split(node.Feature(), ",")
		if features[0] != BOSEOS {
			fmt.Printf("%s %s\n", node.Surface(), node.Feature())
		}
		if node.Next() != nil {
			break
		}
	}
	return nil
}
func main() {
	m, err := mecab.New("-Owakati")
	if err != nil {
		panic(err)
	}
	defer m.Destroy()

	err = parseToNode(m, "すもももももももものうち")
	if err != nil {
		panic(err)
	}
}
