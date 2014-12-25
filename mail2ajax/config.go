

package mail2ajax


import(
    "os"
    "fmt"
	"flag"
    //"io/ioutil"
    //"encoding/json"
	"github.com/BurntSushi/toml"
	"crypto/tls"
	//"database/sql"
	//_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)


//var Config *ConfigStruct

type Config struct {

	Debug bool `toml:"debug" json:"debug" `

    AuthSecret string `toml:"auth_secret" json:"auth_secret" `

	DBEngine string `toml:"db_engine" json:"db_engine"`
    DBConnect string `toml:"db_connect" json:"db_connect"`

	HTTPListen string `toml:"http_listen" json:"http_listen"`
    IMAPAddress string `toml:"imap_adddress" json:"imap_adddress"`
	SMTPLogin string `toml:"smtp_login" json:"smtp_login"`

	Tls *tls.Config
	DB *sqlx.DB
}

/*
func WriteConfig(file_path string)  *Config {
	return &Config{
		AuthSecret: "secret-123456-989654333^&*)0-dsa@here",
		DBEngine: "mysql",
		DBConnect:     "root:secret@localhost/my_mail_server",
		HTTPListen:      "0.0.0.0:8888",
	}
}
*/

func Init() (*Config) {


	var config_file = flag.String("config", "config.toml", "Config file")

	// Load config
	cfg := new(Config)
    //file, e := ioutil.ReadFile(*config_file)
    //if e != nil {
    //    fmt.Printf("Config File Error: %v\n", e)
		//fmt.Printf("create one with -w \n")
     //   os.Exit(1)
    //}
	if _, err := toml.DecodeFile(*config_file, &cfg); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	cfg.Tls = new(tls.Config)
	cfg.Tls.ServerName = cfg.IMAPAddress
	cfg.Tls.InsecureSkipVerify = true

	InitDb(cfg)

	return cfg
}

