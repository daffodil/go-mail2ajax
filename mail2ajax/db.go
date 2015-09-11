

package mail2ajax


import(
	"os"
	"fmt"

	"github.com/jmoiron/sqlx"
	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
	// ?? how do psql and interbase

)

func InitDb(config *Config) {
	fmt.Printf("Conf: ",  config.DBEngine)
	var err error
	config.DB, err = sqlx.Connect(config.DBEngine, config.DBConnect)
	//Db, err = sql.Open("mysql", config.Db.User+":"+config.Db.Password+"@/"+config.Db.Database)
	if err != nil {
		fmt.Printf("Db Login Failed: ", err,"=", config.DBEngine)
		os.Exit(1)
	}
}



