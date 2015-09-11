

package mail2ajax


import(
    "os"
    "fmt"
	"flag"
    "io/ioutil"

	"gopkg.in/yaml.v2"
	"crypto/tls"

	"github.com/jmoiron/sqlx"
)


//var Config *ConfigStruct

type Config struct {

	Debug bool `yaml:"debug" json:"debug" `

    AuthSecret string `yaml:"auth_secret" json:"auth_secret" `

	DBEngine string `yaml:"db_engine" json:"db_engine"`
    DBConnect string `yaml:"db_connect" json:"db_connect"`

	HTTPListen string `yaml:"http_listen" json:"http_listen"`
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

func NewConfig() *Config {
	cfg := new(Config)
	cfg.HTTPListen = "8080"
	return cfg
}

func Init() (*Config, error) {


	var config_file = flag.String("config", "config.yaml", "Config file")

	// Load config
	cfg := NewConfig()
    contents, e := ioutil.ReadFile(*config_file)
    if e != nil {
        fmt.Printf("Config File Error: %v\n", e)
		fmt.Printf("create one with -w \n")
        os.Exit(1)
		return nil, e
    }
	if err := yaml.Unmarshal(contents, &cfg); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	cfg.Tls = new(tls.Config)
	cfg.Tls.ServerName = cfg.IMAPAddress
	cfg.Tls.InsecureSkipVerify = true

	InitDb(cfg)

	return cfg, nil
}

