

package mailadmin

import(

	"github.com/gorilla/mux"
	//"code.google.com/p/go-imap/go1/imap"

	"github.com/daffodil/go-mail2ajax/mail2ajax"
)


var config *mail2ajax.Config


func Configure(cfg *mail2ajax.Config, router *mux.Router){
	config = cfg
	router.HandleFunc("/ajax/domains", DomainsAjaxHandler)
	//router.HandleFunc("/ajax/{domain}/forwardings", ForwardingsHandler)
	router.HandleFunc("/ajax/domain/{domain}", DomainAjaxHandler)
	//router.HandleFunc("/admin/mailbox/{domain}/{mailbox}/create", CreateMailboxHandler)
	//mux.Get("/rpc/mailbox/summary", mailajax.SummaryHandler)

	//router.HandleFunc("/ajax/mailbox/{address}/folders", FoldersHandler)

	//router.HandleFunc("/ajax/mailbox/{address}/folder/{folder}/message/{uid}", MessageHandler)
	//mux.Post("/rpc/mailbox/mb_id/{mb_id:[0-9]+}", mailadmin.MailBoxPostHandler)
	//mux.Get("/rpc/mailboxes", mailadmin.MailBoxesHandler)
}
