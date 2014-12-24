

package m2a


import(
    "os"
    "fmt"
    "io/ioutil"
    "encoding/json"
	"crypto/tls"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


var Cfg = DefaultConfig()

type Config struct {
    AuthSecret string `json:"auth_secret"`
	DbEngine string `json:"db_engine"`
    DbConnect string `json:"db_connect"`
    MailServer string `json:"mail_server"`
    HttpListen string `json:"http_listen"`

	Tls *tls.Config
}

func DefaultConfig()  *Config {
	return &Config{
		AuthSecret: "secret-123456-989654333^&*)0-dsa@here",
		DbEngine: "mysql",
		DbConnect:     "root:secret@localhost/my_mail_server",
		MailServer:     "example.com",
		HttpListen:      "0.0.0.0:8888",
	}
}


func Init() (*Config, *sql.DB) {

	// Load config
	Cfg := new(Config)
    file, e := ioutil.ReadFile("config.json")
    if e != nil {
        fmt.Printf("Config File Error: %v\n", e)
		//fmt.Printf("create one with -w \n")
        os.Exit(1)
    }
    json.Unmarshal(file, &Cfg)

	Cfg.Tls = new(tls.Config)
	Cfg.Tls.ServerName = Cfg.MailServer
	Cfg.Tls.InsecureSkipVerify = true

	InitDb(Cfg)

	return Cfg, Db
}

