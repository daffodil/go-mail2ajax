
//mcore is "main config and objects.. Essentials such as Config, Db Connection
package mcore


import(
    "os"
    "fmt"
    "io/ioutil"
    "encoding/json"
	"crypto/tls"
)


var Config *ConfigStruct;

type ConfigStruct struct {
    Auth string `json:"secret"`
    Db DbUser `json:"db"`
    Tls *tls.Config
    MailServer string `json:"mail_server"`
    WwwPort string `json:"www_port"`
}

type DbUser struct {
    Server string `json:"server"`
    User string `json:"user"`
    Password string `json:"password"`
    Database string `json:"database"`
}

func LoadConfig() *ConfigStruct {

	Config := new(ConfigStruct)
    file, e := ioutil.ReadFile("config.json")
    if e != nil {
        fmt.Printf("Config File error: %v\n", e)
        os.Exit(1)
    }
    json.Unmarshal(file, &Config)

	Config.Tls = new(tls.Config)
	Config.Tls.ServerName = Config.MailServer
	Config.Tls.InsecureSkipVerify = true
	return Config
}
