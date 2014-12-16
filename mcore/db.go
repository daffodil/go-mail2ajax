
//mcore is "main config and objects.. Essentials such as Config, Db Connection
package mcore


import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	// ?? how do psql and interbase
)

// Global Db Connection
var Db *sql.DB

func InitDb(config ConfigStruct) *sql.DB {
	var err error
	Db, err = sql.Open("mysql", config.Db.User+":"+config.Db.Password+"@/"+config.Db.Database)
	if err != nil {
		fmt.Printf("Db Login Failed: ")
	}
	//defer Db.Close()
	return Db
}



