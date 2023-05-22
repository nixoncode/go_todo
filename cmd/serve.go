package cmd

import (
	"log"
	"net/http"

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

	log.Println("Starting http server on port 8080")
	return http.ListenAndServe(":8080", router)

}
