package app

import "go.uber.org/fx"

// InjectApp returns an fx.Option that configures dependency injection in a Go application using the go.uber.org/fx package.
func InjectApp() fx.Option {
	return fx.Provide(
		NewApp,
		fx.Annotate(NewHttpClient, fx.As(new(ClientHttp))),
	)
}

func StartApp(){
	fx.New(
		InjectApp(),
		fx.Invoke(func(app *App){
			if err:=app.Run();err!=nil{
				panic(err)
			}
		})
	)
}

