package main

import (
	"net/http"

	"github.com/Calaghan1/EnrichName/helpers"
	"github.com/Calaghan1/EnrichName/iternal/database"
	"github.com/Calaghan1/EnrichName/pkg/handlers"
)



type App struct {
	Addr string
	router http.Handler
	db database.DB
}
func New_app() *App {
	app := &App {
		Addr: ":3000",
	}
	app.loadroutes()
	return app
}

func (a *App) Start() {
	a.db = database.DB{}
	a.db.Init_database()
	server := &http.Server{
		Addr: a.Addr,
		Handler: a.router,
	}
	err := server.ListenAndServe()
	helpers.CheckErrorFatal(err, "Failed to start server", "Starting server")
}

func (a *App) loadroutes() {
	router := http.NewServeMux()
	h := handlers.Person_handlers{
		Db: &a.db,
	}
	router.HandleFunc("/show_all", h.Show_all)
	router.HandleFunc("/create", h.Create)
	router.HandleFunc("/update", h.UpdateById)
	router.HandleFunc("/delete", h.DeleteByID)
	a.router = router
}

func main() {
	app := New_app()
	app.Start()
}