package main

import ("log"
"github.com/emmanuelalo52/proglog")


func main (){
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}