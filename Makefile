.DEFAULT_GOAL := run

# 生成 README.md
.PHONY: generate
generate:
	@printf '\033[0;36m%s\033[0m\n' "=================================================="
	@printf '\033[0;36m%s\033[0m\n' "*                    Generate                    *"
	@printf '\033[0;36m%s\033[0m\n' "=================================================="
	@go run ./tools/note.go generate
	@go run ./tools/note.go list

# 检查 generator 代码
.PHONY: check-go
check-go:
	@printf '\033[0;36m%s\033[0m\n' "=================================================="
	@printf '\033[0;36m%s\033[0m\n' "*                   Check Go                     *"
	@printf '\033[0;36m%s\033[0m\n' "=================================================="
	(cd ./tools; golangci-lint run)
	(cd ./tools; gofumpt -l .)
	@echo

# 检查所有笔记文件是否需要格式化
.PHONY: check-md
check-md:
	@printf '\033[0;36m%s\033[0m\n' "=================================================="
	@printf '\033[0;36m%s\033[0m\n' "*                   Check Md                     *"
	@printf '\033[0;36m%s\033[0m\n' "=================================================="
	@prettier -c docs

.PHONY: run
run: check-go generate check-md
