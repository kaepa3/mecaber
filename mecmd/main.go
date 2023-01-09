package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/kaepa3/mecaber"
)

const (
	confFile = ".mecaber"
)

func readConfig() {
	paths := []string{
		filepath.Join(".env"),
		filepath.Join(os.Getenv("HOME"), confFile),
		filepath.Join(os.Getenv("HOMEPATH"), confFile),
	}
	for _, v := range paths {
		if err := godotenv.Load(v); err == nil {
			return
		}
		log.Println(v)
	}
	log.Println("Error loading .env file", paths)
	return
}

func getDic() string {
	readConfig()
	dic := os.Getenv("DICTIONARY")
	if len(dic) == 0 {
		return ""
	}
	return "-d " + dic
}

func main() {
	flag.Parse()
	err, m := mecaber.CreateNew(getDic())
	if err != nil {
		panic(err)
	}
	defer m.Destroy()

	done := make(chan struct{})
	errCh, node := m.ParseToNode(flag.Args()[0], done)
	if err != nil {
		panic(err)
	}

LOOP:
	for {
		select {
		case text := <-node:
			fmt.Println(text)
		case err := <-errCh:
			fmt.Println(err)
			break LOOP
		case <-done:
			break LOOP
		}
	}
}
