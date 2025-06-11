package main

import (
	_ "shopx/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"

	"shopx/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
