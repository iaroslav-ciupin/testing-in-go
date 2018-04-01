package code 

type App struct {
	restClient RESTClient
	dbClient DBClient
}

func (a *App) Start() {
	// initialise all components: 
	// controllers, services, repositories, clients using *App
	// start web-server
}

func (a *App) Shutdown() {
	// gracefully shutdown web-server
}

func main() {
	app := App{}
	// set real implementations to app
	go app.Start()
	// listen for OS shutdown signal: e.g. SIGTERM
	app.Shutdown() 
}