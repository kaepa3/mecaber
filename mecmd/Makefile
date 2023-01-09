
build:
	export CGO_LDFLAGS="-L/opt/homebrew/Cellar/mecab/0.996/lib -lmecab -lstdc++" && \
	export CGO_CFLAGS="-I/opt/homebrew/Cellar/mecab/0.996/include"  && \
	go build main.go

run:
	export CGO_LDFLAGS="-L/opt/homebrew/Cellar/mecab/0.996/lib -lmecab -lstdc++" && \
	export CGO_CFLAGS="-I/opt/homebrew/Cellar/mecab/0.996/include"  && \
	go run main.go

