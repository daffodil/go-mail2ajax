
package main

import (
	"fmt"
	//"log"
	"flag"

	"net/http"
	"github.com/gorilla/mux"

	"github.com/daffodil/go-mail2ajax/m2a"
	"github.com/daffodil/go-mail2ajax/mailbox"
)


func main(){

	fWrite := flag.Bool("create", false, "Create config.json file")

	flag.Parse()
	//TODO - write config

	config, db := m2a.Init()
    fmt.Printf("Results: %v\n", config, fWrite)
    
	// gotta be a better way to connect and db etc.. am a newbie
	defer db.Close()
	

	//go mailadmin.Init()

    r := mux.NewRouter()

	mailbox.SetRoutes(r)

    
    fmt.Println("Serving on " + config.HttpListen)
	http.Handle("/", r)
    http.ListenAndServe( config.HttpListen, nil)
    
}

