package sms

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "gopkg.in/goracle.v2" //
)

//DBInfo for ..
type DBInfo struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DsnURL   string `json:"dsnurl"`
}

// New for create dbinfo
func New(aAlias string) *DBInfo {
	return getDBInfo(aAlias)
}

// Close dbinfo
func (dsn *DBInfo) Close() {
	dsn = nil
}

// GetDBInfo for ..
func (dsn *DBInfo) GetDBInfo() *DBInfo {
	return dsn
}

// ExecuteStoreProcedure to execute storeprocedure
func (dsn *DBInfo) ExecuteStoreProcedure(aSQL string, args ...interface{}) bool {
	var myReturn bool

	// initial
	myReturn = false

	//dsn := GetDBInfo(aAlias)
	if dsn != nil {
		var connStr = dsn.User + "/" + dsn.Password + "@" + dsn.DsnURL
		db, err := sql.Open("goracle", connStr)
		defer db.Close()
		if err != nil {
			log.Fatal(err)
			return false
		}
		//log.Println(aSQL)

		result, err := db.Exec(aSQL, args...)
		if err != nil {
			log.Fatal(err)
			return false
		}
		_ = result // not used
		myReturn = true
	}

	return myReturn
}

// SelectSQL to select statement db
func (dsn *DBInfo) SelectSQL(aSQL string) (*sql.Rows, error) {
	var myReturn *sql.Rows

	if dsn != nil {
		var connStr = dsn.User + "/" + dsn.Password + "@" + dsn.DsnURL

		db, err := sql.Open("goracle", connStr)
		defer db.Close()
		if err != nil {
			//log.Fatal(err)
			return nil, err
		}
		//log.Println(aSQL)

		rows, err := db.Query(aSQL)
		//defer rows.Close()
		if err != nil {
			//fmt.Println("Error running query")
			//fmt.Println(err)
			return nil, err
		}
		myReturn = rows
	}

	return myReturn, nil
}

// ExecuteSQL to execute statement db
func (dsn *DBInfo) ExecuteSQL(aSQL string) (int64, error) {
	var myReturn int64

	if dsn != nil {
		var connStr = dsn.User + "/" + dsn.Password + "@" + dsn.DsnURL

		db, err := sql.Open("goracle", connStr)
		defer db.Close()
		if err != nil {
			log.Fatal(err)
			return 0, err
		}

		//log.Println(aSQL)
		rows, err := db.Exec(aSQL)
		if err != nil {
			//log.Fatal(err)
			return 0, err
		}

		// Get row affected
		affected, err := rows.RowsAffected()
		if err != nil {
			//log.Fatal(err)
			return 0, err
		}

		myReturn = affected
	}

	return myReturn, nil
}

func getDBInfo(alias string) *DBInfo {
	var myReturn *DBInfo

	myReturn, err := getUsernameAndPwd(alias)
	if err != nil {

	}

	return myReturn
}

func getUsernameAndPwd(alias string) (*DBInfo, error) {
	var myReturn DBInfo
	var aSQL string
	//var key int

	var dbusername = "devl"
	var dbuserpass = "developer"
	var dbname = "172.22.204.25:1521/PED"
	var connStr = dbusername + "/" + dbuserpass + "@" + dbname

	db, err := sql.Open("goracle", connStr)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	aSQL = " SELECT A.BASEID, NVL(C.DATASOURCE, ' '), A.USERNAME, A.PASSWORD, A.DATABASE" +
		" FROM SMS_DATABASE A, SMS_DATABASE_CONFIG B, SMS_DATABASE_DATASOURCEURL C" +
		" WHERE A.ALIAS = '" + alias + "'" +
		" AND A.BASEID = B.BASEID" +
		" AND A.DATABASE = C.DATABASE" +
		" AND B.BASECONF = 1"

	//log.Println(aSQL)

	rows, err := db.Query(aSQL)
	defer rows.Close()
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return nil, err
	}

	var baseid int
	var datasource string
	var username string
	var password string
	var database string
	//
	if rows.Next() {
		rows.Scan(&baseid, &datasource, &username, &password, &database)
	}

	myReturn.User = decodeString(username, baseid)
	myReturn.Password = decodeString(password, baseid)
	myReturn.DsnURL = datasource

	return &myReturn, nil
}

func decodeString(str string, key int) string {
	var myReturn string

	for i := 0; i < len(str)/3; i++ {
		tmp := str[i+(i*3)-i : (i*3)+3]
		i64, err := strconv.ParseInt(tmp, 10, 64)
		if err == nil {
			code := (rune)(int(i64)+256-key) % 256
			myReturn = myReturn + string(code)
		}
	}

	return myReturn
}
