package main

import (
	"fmt"
	"github.com/poupardm-GhostWrath/gator/internal/config"
)

func main() {
	cfg := config.Read()
	fmt.Printf("db_url: %s\ncurrent_user_name: %s\n", cfg.DBUrl, cfg.CurrentUserName)
	err := cfg.SetUser("Matt")
	if err != nil {
		fmt.Println(err)
		return
	}
	cfg = config.Read()
	fmt.Printf("db_url: %s\ncurrent_user_name: %s\n", cfg.DBUrl, cfg.CurrentUserName)
	return
}