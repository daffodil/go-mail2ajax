
package main

import (
	"fmt"
	//"log"


	"net/http"
	"github.com/gorilla/mux"

	"github.com/daffodil/go-mail2ajax/mcore"
	"github.com/daffodil/go-mail2ajax/mailbox"
)


func main(){
    

	config := mcore.LoadConfig()
    fmt.Printf("Results: %v\n", config)
    
	// gotta be a better way to connect and db etc.. am a newbie
	defer mcore.Db.Close()
	

	//go mailadmin.Init()

    r := mux.NewRouter()

	mailbox.SetRoutes(r)

    
    fmt.Println("Serving on " + config.WwwPort)
	http.Handle("/", r)
    http.ListenAndServe(":" + config.WwwPort, nil)
    
}

