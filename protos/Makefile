# Makefile for generating protobuf code
# 用于生成 protobuf 代码的 Makefile

# Mark targets as phony (not files) to ensure they always run
# 标记目标为伪目标（非文件）以确保它们总是执行
.PHONY: install generate clean regen

# Install protoc-gen-go plugin
# 安装 protoc-gen-go 插件
install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Generate protobuf code
# 生成 protobuf 代码
generate:
	protoc --go_out=. --go_opt=paths=source_relative protoenumstatus/protoenumstatus.proto
	protoc --go_out=. --go_opt=paths=source_relative protoenumresult/protoenumresult.proto

# Clean generated files
# 清理生成的文件
clean:
	rm -f protoenumstatus/*.pb.go
	rm -f protoenumresult/*.pb.go

# Regenerate and verify
# 重新生成并验证
regen: clean generate
	@echo "success"
