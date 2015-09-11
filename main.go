
package main

import (
	"fmt"
	//"log"
	//"flag"
	//"crypto/md5"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/daffodil/go-mail2ajax/mail2ajax"
	//"github.com/daffodil/go-mail2ajax/mailbox"
	"github.com/daffodil/go-mail2ajax/mailadmin"

)


func main(){

	//#fWrite := flag.Bool("create", false, "Create config.toml file")

	//flag.Parse()
	//TODO - write config

	config, err := mail2ajax.Init()
    fmt.Printf("Results: %v\n", config, err)
    
	// gotta be a better way to connect and db etc.. am a newbie
	defer config.DB.Close()
	



    r := mux.NewRouter()
	mailadmin.Configure(config, r)
	//mailbox.Configure(config, r)

    
    fmt.Println("Serving on " + config.HTTPListen)
	http.Handle("/", r)
    http.ListenAndServe( config.HTTPListen, nil)
    
}

