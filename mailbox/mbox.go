package mailbox


import(

	//"fmt"
	"net/mail"
	"code.google.com/p/go-imap/go1/imap"

	"net/http"
	"github.com/gorilla/mux"

	"github.com/daffodil/go-mail2ajax/mcore"
)

//===============================================
type MbLogin  struct{
	Email string
	Password string
	Active uint
}


//== Check and get credentials of email
func GetImapClient(resp http.ResponseWriter, req *http.Request)(*imap.Client) {

	//= Get Email address and validate
	vars := mux.Vars(req)

	addr, addr_err := mail.ParseAddress(vars["address"])
	if addr_err != nil {
		SendErrorPayload("Invalid email address", resp)
		return nil
	}

	// Get email and Password and active from DB and validate
	sql := "select email, password, active from virtual_users where email=?"
	mb := new(MbLogin)
	err := mcore.Db.QueryRow(sql, addr.Address).Scan(&mb.Email, &mb.Password, &mb.Active)
	if err != nil {
		SendErrorPayload("Mailbox not found", resp)
		return nil
	}

	//= Connect to Server
	client, conn_err := imap.DialTLS(mcore.Config.MailServer, mcore.Config.Tls)
	if conn_err != nil {
		SendErrorPayload( conn_err.Error(), resp )
		return nil
	}
	_, login_err := client.Login( mb.Email, mb.Password )
	if login_err != nil {
		SendErrorPayload("IMAP login error", resp)
		return nil
	}
	return client

}


func GetUIDs(mbox string, client *imap.Client) ([]uint32, error ) {
	
	uids := make([]uint32, 0)
	
	cmd, err := client.Select(mbox, true)
	if err != nil {
		return uids, err
	}
	
	//== Get UIDS of all messages
	cmd, err = imap.Wait( client.UIDSearch("", "ALL") )
	if err != nil {
		return uids, err
	}
	
	for idx := range cmd.Data {
		for _, uid := range cmd.Data[idx].SearchResults() {
			uids = append(uids, uid)
		}
	}
	return uids, nil
	
}
func GetLastUIDs(alluids []uint32) *imap.SeqSet {
	//payload.Uids, err = GetUIDs("INBOX", client)
	
	//= Calc last few messages
	lenny := len(alluids)
	last := lenny - 50  // ################
	if  last < 0 {
		last = 0
	}
	//= Make List of messages uids
	uidlist, _ := imap.NewSeqSet("")
	for _, uid := range alluids[last:lenny] {
		uidlist.AddNum(uid)
	}
	return uidlist
}
