
test:
	export CGO_LDFLAGS="-L/opt/homebrew/Cellar/mecab/0.996/lib -lmecab -lstdc++" && \
	export CGO_CFLAGS="-I/opt/homebrew/Cellar/mecab/0.996/include"  && \
	go test

build:
	export CGO_LDFLAGS="-L/opt/homebrew/Cellar/mecab/0.996/lib -lmecab -lstdc++" && \
	export CGO_CFLAGS="-I/opt/homebrew/Cellar/mecab/0.996/include"  && \
	go build main.go

install:
	export CGO_LDFLAGS="-L/opt/homebrew/Cellar/mecab/0.996/lib -lmecab -lstdc++" && \
	export CGO_CFLAGS="-I/opt/homebrew/Cellar/mecab/0.996/include"  && \
	go install github.com/kaepa3/mecaber/mecmd@latest
