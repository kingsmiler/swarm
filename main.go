package main

import (
// 通过使用下划线_别名, 仅执行这些包下面所有文件的init方法,
// 并不会真正的导入这些包.
	_ "github.com/docker/docker/pkg/discovery/file"
	_ "github.com/docker/docker/pkg/discovery/kv"
	_ "github.com/docker/docker/pkg/discovery/nodes"
	_ "github.com/docker/swarm/discovery/token"

	"github.com/docker/swarm/cli"
)

func main() {
	cli.Run()
}
