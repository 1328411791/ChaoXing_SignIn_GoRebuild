package main

import (
	_ "ChaoXing_SignIn_GoRebuild/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"ChaoXing_SignIn_GoRebuild/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
