package main

import (
	_ "gfDemo/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gfDemo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
