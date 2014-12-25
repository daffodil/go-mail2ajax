

package mailadmin

import(

	"fmt"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)

type Forward struct {
	FwdID int `db:"fwd_id" json:"fwd_id"`
	Source string `db:"source" json:"source"`
	Destination string `db:"destination" json:"destination"`
}

type ForwardingsPayload struct {
	Success bool `json:"success"` // keep extjs happy
	Forwardings []Forward `json:"forwardings"`
	Error string `json:"error"`
}

func NewForwardingsPayload() ForwardingsPayload {
	t := ForwardingsPayload{}
	t.Success = true
	t.Forwardings = make([]Forward, 0)
	return t
}

// gets forwardings from database
// TODO filter by domain in source
func GetForwardings(domain string) ([]Forward, error) {
	var rows []Forward
	err := config.DB.Select(&rows, "SELECT * FROM forwardings order by source asc ")
	return rows, err
}

func ForwardingsHandler(resp http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	domain := vars["domain"]
	//as_rows := vars["rows"]
	// TODO check domain exists

	payload := NewForwardingsPayload()

	var err error
	payload.Forwardings, err = GetForwardings(domain)
	if err != nil{
		fmt.Println(err)
	}




	json_str, _ := json.MarshalIndent(payload, "" , "  ")
	fmt.Fprint(resp, string(json_str))
}
