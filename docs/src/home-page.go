package main

import (
	"github.com/whale1017/go-app/v10/pkg/analytics"
	"github.com/whale1017/go-app/v10/pkg/app"
	"github.com/whale1017/go-app/v10/pkg/ui"
)

type homePage struct {
	app.Compo
}

func newHomePage() *homePage {
	return &homePage{}
}

func (p *homePage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *homePage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
	analytics.Page("home", nil)
}

func (p *homePage) Render() app.UI {
	return newPage().
		Title("go-app").
		Icon("/web/icon.png").
		Index(
			newIndexLink().Title("What is go-app?"),
			newIndexLink().Title("Updates"),
			newIndexLink().Title("Declarative Syntax"),
			newIndexLink().Title("Standard HTTP Server"),
			newIndexLink().Title("Other features"),
			newIndexLink().Title("Built With go-app"),

			app.Div().Class("separator"),

			newIndexLink().Title("Next"),
		).
		Content(
			ui.Flow().
				StretchItems().
				Spacing(84).
				Content(
					newRemoteMarkdownDoc().
						Class("fill").
						Src("/web/documents/what-is-go-app.md"),
					newRemoteMarkdownDoc().
						Class("fill").
						Class("updates").
						Src("/web/documents/updates.md"),
				),

			app.Div().Class("separator"),

			newRemoteMarkdownDoc().Src("/web/documents/home.md"),

			app.Div().Class("separator"),

			newBuiltWithGoapp().ID("built-with-go-app"),

			app.Div().Class("separator"),

			newRemoteMarkdownDoc().Src("/web/documents/home-next.md"),
		)
}
