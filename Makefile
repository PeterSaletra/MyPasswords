HOME := /home/peter

clean:
	@if [ -f /usr/local/bin/mypasswords ]; then \
		sudo rm /usr/local/bin/mypasswords; \
	else echo "Bin already removed"; \
	fi

	@if [ -d $(HOME)/.local/.mypasswords ]; then \
		rm -rf $(HOME)/.local/.mypasswords; \
	else echo "Directory does not exists"; \
	fi
.PHONY: clean

build:
	@go build -o mypasswords main.go
	@sudo cp mypasswords /usr/local/bin/mypasswords
	@rm mypasswords
.PHONY: build

