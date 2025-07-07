.DEFAULT_GOAL := run

# 生成 README.md
.PHONY: generate
generate:
	@printf '\033[0;36m%s\033[0m\n' "=================================================="
	@printf '\033[0;36m%s\033[0m\n' "*                    Generate                    *"
	@printf '\033[0;36m%s\033[0m\n' "=================================================="
	@go run ./tools/generator.go
	@echo

# 预览 README.md
.PHONY: preview
preview:
	@glow -w 0 ./README.md

# 检查所有笔记文件是否需要格式化
.PHONY: check
check:
	@printf '\033[0;36m%s\033[0m\n' "=================================================="
	@printf '\033[0;36m%s\033[0m\n' "*                     Check                      *"
	@printf '\033[0;36m%s\033[0m\n' "=================================================="
	@prettier -c docs

.PHONY: run
run: generate preview check
