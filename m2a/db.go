

package m2a


import(
	"os"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	// ?? how do psql and interbase
)

// Global Db Connection
var Db *sql.DB

func InitDb(config *Config) {
	fmt.Printf("Conf: ",  config.DbEngine)
	var err error
	Db, err = sql.Open(config.DbEngine, config.DbConnect)
	//Db, err = sql.Open("mysql", config.Db.User+":"+config.Db.Password+"@/"+config.Db.Database)
	if err != nil {
		fmt.Printf("Db Login Failed: ", err,"=", config.DbEngine)
		os.Exit(1)
	}
}



