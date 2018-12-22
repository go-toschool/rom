APP=rom
BIN=$(PWD)/$(APP)
GO ?= go

build b: clean templates
	@echo "[buil] Building..."
	@$(GO) build -o $(BIN)\

run r: build
	@echo "[run] Running..."
	@$(BIN)

clean:
	@echo "[clean] Removing $(BIN)..."
	@rm -rf templates/*.qtpl.go

templates:
	@echo "[templates] Generating..."
	@qtc templates

.PHONY: build run clean templates
