

package mailadmin

import(

	"fmt"
	"net/http"
	"encoding/json"

	//"github.com/gorilla/mux"
)

type Domain struct {
	DomainID int `db:"domain_id" json:"domain_id"`
	Domain string `db:"domain" json:"domain"`
}

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
	err := config.DB.Select(&rows, "SELECT id as dsomain_id, name as domain FROM virtual_domains order by name asc ")
	return rows, err
}

func DomainsHandler(resp http.ResponseWriter, req *http.Request) {

	//_ := mux.Vars(req)


	payload := NewDomainsPayload()

	var err error
	payload.Domains, err = GetDomains()
	if err != nil{
		fmt.Println(err)
	}


	json_str, _ := json.MarshalIndent(payload, "" , "  ")
	fmt.Fprint(resp, string(json_str))
}
