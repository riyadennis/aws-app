run:
	go run main.go
deps:
	glide install
clean:
	rm -rf vendor/
	rm userdata.db