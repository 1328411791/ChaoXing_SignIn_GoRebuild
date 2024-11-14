package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "ChaoXing_SignIn_GoRebuild/internal/packed"

	_ "ChaoXing_SignIn_GoRebuild/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"ChaoXing_SignIn_GoRebuild/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
