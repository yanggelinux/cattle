# ----------------------------------------
# Go Cross-Platform Build Makefile
# ----------------------------------------

# 输出目录和文件
BIN_DIR := cmd/bin
APP_NAME := cattle
SRC_FILE := cmd/cattle/cattle.go

# 通用编译参数（去符号表与调试信息）
LDFLAGS := -ldflags="-s -w"

# 自动检测操作系统
ifeq ($(OS),Windows_NT)
	RM := del /Q
	MKDIR := if not exist $(BIN_DIR) mkdir $(BIN_DIR)
	EXT := .exe
else
	RM := rm -rf
	MKDIR := mkdir -p $(BIN_DIR)
	EXT :=
endif

# 默认目标
.DEFAULT_GOAL := help

# ----------------------------------------
# 🧱 Build Commands
# ----------------------------------------

## build.unix: 构建 Linux 版本
build.unix:
	@echo "🔧 Building for Linux amd64..."
	@$(MKDIR)
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BIN_DIR)/$(APP_NAME)$(EXT) $(SRC_FILE)
	@echo "✅ Build done -> $(BIN_DIR)/$(APP_NAME)$(EXT)"

## build.windows: 构建 Windows 版本
build.windows:
	@echo "🔧 Building for Windows amd64..."
	@$(MKDIR)
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BIN_DIR)/$(APP_NAME).exe $(SRC_FILE)
	@echo "✅ Build done -> $(BIN_DIR)/$(APP_NAME).exe"

## build.local: 根据当前系统构建
build.local:
	@echo "🔧 Building for local system..."
	@$(MKDIR)
	@go build $(LDFLAGS) -o $(BIN_DIR)/$(APP_NAME)$(EXT) $(SRC_FILE)
	@echo "✅ Local build done -> $(BIN_DIR)/$(APP_NAME)$(EXT)"

# ----------------------------------------
# 🧹 Clean Command
# ----------------------------------------

## clean: 清理构建产物
clean:
	@echo "🧹 Cleaning build files..."
	@go clean
	@$(RM) $(BIN_DIR)/$(APP_NAME)*$(EXT) 2>nul || true
	@echo "✅ Clean done."

# ----------------------------------------
# 📘 Help
# ----------------------------------------

help:
	@echo ""
	@echo "📦 Available Commands:"
	@echo "  make build.unix     - Build Linux amd64 binary"
	@echo "  make build.windows  - Build Windows amd64 binary"
	@echo "  make build.local    - Build binary for current OS"
	@echo "  make clean          - Clean build outputs"
	@echo "  make help           - Show this help"
	@echo ""
