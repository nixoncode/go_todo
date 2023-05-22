package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/cobra"
)

func NewServeCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "serve",
		Short: "Starts the web server in http://127.0.0.1:8080",
		Run: func(cmd *cobra.Command, args []string) {
			err := server()

			if err != nil {
				log.Fatalln(err)
			}
		},
	}

	return command
}

func server() error {
	router := chi.NewRouter()

	// mount routes
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	server := &http.Server{Addr: ":8080", Handler: router}
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	log.Println("Starting http server on port 8080")
	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, ss := context.WithTimeout(serverCtx, 30*time.Second)
		if ss != nil {
			log.Fatal("Can't wait anymore", ss)
		}

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			panic(err)
		}
		serverStopCtx()
	}()

	// Run the server
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()

	return nil
}
