package main

import "github.com/whale1017/go-app/v10/pkg/app"

const (
	installApp = "/app/install"
	updateApp  = "/app/update"
)

func handleAppInstall(ctx app.Context, a app.Action) {
	ctx.ShowAppInstallPrompt()
}

func handleAppUpdate(ctx app.Context, a app.Action) {
	ctx.Reload()
}
