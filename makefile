all: build 

build:
	touch db/estoque.db
	go build -o api


clear:
	rm db/estoque.db