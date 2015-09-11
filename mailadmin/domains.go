

package mailadmin

import(

	"fmt"
	"net/http"
	"encoding/json"

	//"github.com/gorilla/mux"
)


type DomainsPayload struct {
	Success bool `json:"success"` // keep extjs happy
	Domains []Domain `json:"domains"`
	Error string `json:"error"`
}



func NewDomainsPayload() DomainsPayload {
	t := DomainsPayload{}
	t.Success = true
	t.Domains = make([]Domain, 0)
	return t
}

// gets forwardings from database
// TODO filter by domain in source
func GetDomains() ([]Domain, error) {
	var rows []Domain
	err := config.DB.Select(&rows, "SELECT domain, description, aliases, mailboxes, maxquota, quota, transport, backupmx, created, modified, active FROM domain order by domain asc ")
	return rows, err
}

// Handles /ajax/domains
func DomainsAjaxHandler(resp http.ResponseWriter, req *http.Request) {

	//_ := mux.Vars(req)
	// TODO auth

	payload := NewDomainsPayload()

	var err error
	payload.Domains, err = GetDomains()
	if err != nil{
		fmt.Println(err)
		payload.Error = "DB Error: " + err.Error()
	}


	json_str, _ := json.MarshalIndent(payload, "" , "  ")
	fmt.Fprint(resp, string(json_str))
}
