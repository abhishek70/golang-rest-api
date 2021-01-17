package router

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type MuxRouter struct {

}

var (
	muxDispatcher = mux.NewRouter()
	l *log.Logger
)

func NewMuxRouter(logger *log.Logger) Router {
	l = logger
	return &MuxRouter{}
}

func (*MuxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*MuxRouter) SERVE(port string) {

	l.Println("HTTP server started on port", port)

	// Setting up the web server
	s := &http.Server{
		Addr: ":9090",
		Handler: muxDispatcher,
		ReadHeaderTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
		IdleTimeout: 120*time.Second,
	}

	// GoRoutines
	go func() {

		// Running the web server
		err := s.ListenAndServe()

		if err != nil {
			l.Println("Server shutting down on port ", port)
			os.Exit(1)
		}
	}()

	// Creating a channel
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	l.Println("Receive signal:", sig)


	// Gracefully shutdown / wait for client request to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	// Shut down the server context
	_ = s.Shutdown(ctx)
}


