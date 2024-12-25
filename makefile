all: build 

build:
	touch db/estoque.db
	go build -o api
	./api


clear:
	rm db/estoque.db