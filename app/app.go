package app

import (
	"github.com/crocus7724/korat/config"
	"github.com/crocus7724/korat/view"
	"github.com/crocus7724/korat/api"
	"fmt"
	"os"
	"os/exec"
)

var (
	userConfig config.Config
)

func Start(c config.Config) {
	userConfig = c
	token := userConfig.User[0].Token
	url := userConfig.User[0].Url
	if token == "" {
		fmt.Println("korat require github token")
		os.Exit(1)
	}
	api.Init(url, token)
	view.Init()
	go ViewerRepositories()

	view.Start()
}

func OpenUrl(url string) {
	err := exec.Command("open", url).Run()
	if err != nil {
		view.ShowError(err)
	}
}
