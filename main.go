package main

import (
	_ "shopx/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"shopx/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
