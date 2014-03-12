package ttgDB

import(
	"database/sql"
	_"github.com/mattn/go-sqlite3"
)


func CheckErr(err error){
	if(err != nil){
		panic(err)
	}
}
func SetupDB() *sql.DB{
	db,err := sql.Open("sqlite3","./temp.db")
	if(err != nil){
		panic(err)
	}
	InitDB(db)
	return db
}

func InitDB(db *sql.DB){
	row, err := db.Query("select name from sqlite_master where type='table' and table_name='users';")
	CheckErr(err)
	defer row.Close()
	row.Next()
	var name string
	row.Scan(&name)
	if(name == ""){
		_, err := db.Exec("create table users ( id integer primary key autoincrement, username varchar(32), email varchar(320), passphrase varchar(200));")
		CheckErr(err)
		
	}
}



