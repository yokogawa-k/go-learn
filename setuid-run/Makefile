SHELL := bash
BIN := test-server

build:
	go build -o $(BIN)

clean:
	rm -vf $(BIN)

test: build
	sudo chown root $(BIN)
	sudo chmod 4755 $(BIN)
	stat $(BIN)
	./$(BIN) &
	ps o euid,ruid,egid,rgid,comm -C $(BIN)
	sudo ss -tlnp sport = 84
	killall $(BIN)

.PHONY: build clean test
