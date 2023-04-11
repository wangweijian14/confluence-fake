package main

import (
	_ "confluence_fake/internal/packed"

	_ "confluence_fake/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"confluence_fake/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
