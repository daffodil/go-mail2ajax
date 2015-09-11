

package mailadmin

import(

	"fmt"
	"net/http"
	"encoding/json"

	//"github.com/gorilla/mux"
)

type Domain struct {
	//DomainID int `db:"domain_id" json:"domain_id"`
	Domain string 		`db:"domain" json:"domain"`
	Description string 	`db:"description" json:"description"`
	Aliases int 		`db:"aliases" json:"aliases"`
	Mailboxes int 		`db:"mailboxes" json:"mailboxes"`
	MaxQuota int 		`db:"maxquota" json:"maxquota"`
	Quota int 			`db:"quota" json:"quota"`
	Transport string	`db:"transport" json:"transport"`
	BackupMx int 		`db:"backupmx" json:"backupmx"`
	Created string		`db:"created" json:"created"`
	Modified string		`db:"modified" json:"modified"`
	Active int 			`db:"active" json:"active"`
}

type DomainPayload struct {
	Success bool `json:"success"` // keep extjs happy
	Domain Domain `json:"domain"`
	Error string `json:"error"`
}



func NewDomainPayload() DomainPayload {
	payload := DomainPayload{}
	payload.Success = true
	payload.Domain = Domain{}
	return payload
}



func GetDomain() (Domain, error) {
	var row Domain
	err := config.DB.QueryRow(&row, "SELECT domain, description, aliases, mailboxes, maxquota, quota, transport, backupmx, created, modified, active FROM domain order by domain asc ")

	return row, err
}


// Handles /ajax/domain/example.com
func DomainAjaxHandler(resp http.ResponseWriter, req *http.Request) {

	//_ := mux.Vars(req)
	// TODO auth

	payload := NewDomainPayload()

	var err error
	payload.Domain, err = GetDomain()
	if err != nil{
		fmt.Println(err)
		payload.Error = "DB Error: " + err.Error()
	}


	json_str, _ := json.MarshalIndent(payload, "" , "  ")
	fmt.Fprint(resp, string(json_str))
}
