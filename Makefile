.PHONY: generate
generate:
	@printf '\033[0;36m%s\033[0m\n' "=================================================="
	@printf '\033[0;36m%s\033[0m\n' "*  Generate                                      *"
	@printf '\033[0;36m%s\033[0m\n' "=================================================="
	@go run generator.go
	@echo

.PHONY: preview
preview:
	@glow -w 0 ./README.md

.PHONY: check
check:
	@printf '\033[0;36m%s\033[0m\n' "=================================================="
	@printf '\033[0;36m%s\033[0m\n' "*  Check                                         *"
	@printf '\033[0;36m%s\033[0m\n' "=================================================="
	@prettier -c docs


.PHONY: run
run: generate preview check
