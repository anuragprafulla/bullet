package main

import (
	"log"
	"os"
)

func main() {
	args := Args{
		conn: "postgres://user:password@fullstack-postgres:5432/fullstack_api?sslmode=disable",
		port: ":8080",
	}

	/*if conn := os.Getenv("DB_CONN"); conn != "" {
		args.conn = conn
	}
	*/
	if port := os.Getenv("PORT"); port != "" {
		args.port = ":" + port
	}
	// run server
	if err := Run(args); err != nil {
		log.Println(err)
		log.Println(args.conn)
	}

	log.Println("Router initialzied")

}
