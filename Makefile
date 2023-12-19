
all:
	cd src; make dynamic
	go build -o bin/main main.go
	bin/main