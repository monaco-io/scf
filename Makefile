.PHONY : build
build :
	rm -rf main
	GOOS=linux GOARCH=amd64 go build -o main main.go

.PHONY : zip
zip :
	GOOS=linux GOARCH=amd64 go build -o main main.go
	zip main.zip main
	rm -rf main

.PHONY : clean
clean :
	rm -rf main

