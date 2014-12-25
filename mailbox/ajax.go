

package mailbox

import(

	"fmt"
	"net/http"
	"bytes"
	"net/mail"
	"encoding/json"
	//"database/sql"

	"github.com/gorilla/mux"
	"code.google.com/p/go-imap/go1/imap"

	"github.com/daffodil/go-mail2ajax/m2a"
)

var config *m2a.Config

func Configure(cfg *m2a.Config, router *mux.Router){
	config = cfg
	router.HandleFunc("/ajax/mailbox/{address}/summary", SummaryHandler)
	//mux.Get("/rpc/mailbox/summary", mailajax.SummaryHandler)

	router.HandleFunc("/ajax/mailbox/{address}/folders", FoldersHandler)

	router.HandleFunc("/ajax/mailbox/{address}/folder/{folder}/message/{uid}", MessageHandler)
	//mux.Post("/rpc/mailbox/mb_id/{mb_id:[0-9]+}", mailadmin.MailBoxPostHandler)
	//mux.Get("/rpc/mailboxes", mailadmin.MailBoxesHandler)
}

//===============================================
type ErrorPayload struct {
	Success bool `json:"success"`
	Error string `json:"error"`
}


func SendErrorPayload(err string, resp http.ResponseWriter){

	payload := new(ErrorPayload)
	payload.Success = true
	payload.Error = err

	json_str, _ := json.MarshalIndent(payload, "" , "  ")
	fmt.Fprint(resp, string(json_str))
}

// Content-Type to json, and allow origin
func SetAjaxHeaders(w http.ResponseWriter){

	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	//w.Header().Set("Access-Control-Allow-Origin", "*") // cache control ?
	//w.Header().Set("Access-Control-Allow-Origin", "*") // no cache ?
}



//===================================================================
type Header struct {
	Uid uint32 `json:"uid"`
	
	FromName string `json:"from_name"`
	FromEmail string `json:"from_email"`
	Subject string `json:"subject"`
	
	Date string `json:"date"`
	
	Seen bool `json:"seen"`
	Flagged bool `json:"flagged"`
	Answered bool `json:"answered"`
}


type SummaryPayload struct {
	Headers[] *Header `json:"headers"`
	Success bool `json:"success"`
	Folders[] *Folder `json:"folders"`
	Uids[] uint32 `json:"uids"`
	
	Error string `json:"error"`
}



func SummaryHandler(resp http.ResponseWriter, req *http.Request) {
	
	
	
	//= Setup AJAX Payload
	SetAjaxHeaders(resp)

	client := GetImapClient(resp, req)
	if client == nil {
		return
	}
	defer func() { client.Logout(0) }()
	
	payload := new(SummaryPayload)
	payload.Folders = make([]*Folder, 0)
	payload.Uids = make([]uint32, 0)
	payload.Headers = make([]*Header, 0)
	payload.Success = true
	
	var err error
	payload.Folders, err = GetFolders(client)
	if err != nil {
		payload.Error = payload.Error + err.Error() + "\n"
	}
	
	//= Select inbox
	payload.Uids, err = GetUIDs("INBOX", client)
	

	uidlist := GetLastUIDs(payload.Uids)
	
	//----------------------------------------------
	//== Fetch last few  messages
	cmd, err := imap.Wait( client.UIDFetch(uidlist, "FLAGS", "INTERNALDATE", "RFC822.SIZE", "RFC822.HEADER") )
	if err != nil {
		fmt.Println("#################", err)
	}
	
	for _, rsp  := range cmd.Data {
		
		header := imap.AsBytes(rsp.MessageInfo().Attrs["RFC822.HEADER"])
		mm := new(Header)
		mm.Uid = rsp.MessageInfo().UID
		
			
		for flag, boo := range  rsp.MessageInfo().Flags {
			//fmt.Println( boo, flag, flag == "\\Seen" )
			if flag == "\\Seen" && boo {
				mm.Seen = true
			}
			if flag == "\\Flagged" && boo {
				mm.Flagged = true
			}
		}
        if msg, _ := mail.ReadMessage(bytes.NewReader(header)); msg != nil {
        
			//fmt.Println(" msh_err= ", msg_err)
			mm.Subject = msg.Header.Get("Subject")
			
			
			// From
			from, fro_err := mail.ParseAddress(msg.Header.Get("From"))
			if fro_err != nil {
				fmt.Println("address ettot")
			} else {
				mm.FromName = from.Name
				mm.FromEmail = from.Address
			}
			
			// Date
			dat := imap.AsDateTime(rsp.MessageInfo().Attrs["INTERNALDATE"] )
			//if dat_err != nil {
			//	fmt.Println("date error", dat_err)
			//} else {
			mm.Date = dat.Format("2006-01-02 15:04:05")
			//} 
			
			payload.Headers = append(payload.Headers, mm)
        }
		//mm := cmd.Data[midx].MessageInfo() 
			
			
			//fmt.Println(mm.Attrs, mm.InternalDate)
			
			//payload.Messages = append(payload.Messages, mess)
		//}
	}
	

	json_str, _ := json.MarshalIndent(payload, "" , "  ")   
	fmt.Fprint(resp, string(json_str))
	
}


