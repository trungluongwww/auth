package main

import (
	"fmt"
	"github.com/trungluongwww/auth/config"
	"github.com/trungluongwww/auth/db"
	"github.com/trungluongwww/auth/pkg/server"
	"github.com/trungluongwww/auth/register"
)

func main() {
	cfg, err := config.NewEnv()
	if err != nil {
		panic(err)
	}

	gdb, err := db.NewDB(*cfg)
	if err != nil {
		panic(err)
	}

	r := register.NewRegister(gdb, *cfg)

	e := server.Bootstrap(r, *cfg)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
