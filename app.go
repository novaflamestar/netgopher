package main

type App struct {
	Args *argBundle
	Net  NetConnectors
}

func (a *App) SetArgs(args *argBundle) {
	a.Args = args
}

func (a *App) SetUpNetConn() {
	switch a.Args.FType {
	case "server":
		a.Net = Server{Addr: a.Args.Adder, Port: a.Args.Port}
	case "client":
		a.Net = Client{Addr: a.Args.Adder, Port: a.Args.Port}
	default:
		panic("Unknown Application Type")
	}
}
