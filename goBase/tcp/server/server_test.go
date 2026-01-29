package server

import "testing"

/*
问题: dyld[1692]: missing LC_UUID load command
原因:
	1. Go 编译器生成的二进制文件缺少 LC_UUID 加载命令
	2. macOS 的动态链接器 (dyld) 期望这个命令
	3. 通常与缓存、符号表或编译缓存相关
解决方案:
	方案1：清理 Go 编译缓存（最有效）
	# 1. 清理所有 Go 编译缓存
	go clean -cache -testcache

	# 2. 清理特定包的缓存
	go clean -cache -testcache ./...

	# 3. 强制重新编译
	go test -a ./...

	# 4. 如果还有问题，清理整个 Go 模块缓存
	go clean -modcache
	方案2: 删除临时目录文件
	# macOS 临时目录可能包含损坏的文件
	rm -rf /private/var/folders/go-build*
	rm -rf /tmp/go-build*

*/

func TestCreateServer(t *testing.T) {
	CreateServer()
}
