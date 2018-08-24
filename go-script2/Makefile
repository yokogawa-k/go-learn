SHELL := bash
GOFILE := main.go

test:
	sudo chown root $(GOFILE)
	sudo chmod 4755 $(GOFILE)
	stat $(GOFILE)
	./$(GOFILE)

.PHONY: build clean test
