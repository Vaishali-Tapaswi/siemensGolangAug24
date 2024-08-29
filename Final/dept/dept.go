package deptdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/magiconair/properties"
)

func getDb() (db *sql.DB) {
	// load this from properties file
	p := properties.MustLoadFile("dbcon.properties", properties.UTF8)
	if login, ok := p.Get("login"); ok {
		if password, ok := p.Get("password"); ok {
			if url, ok := p.Get("url"); ok {
				str := login + ":" + password + "@" + url
				fmt.Println(str)
				db, err := sql.Open("mysql", login+":"+password+"@"+url)
				if err != nil {
					panic("Problem with getting database connection")
				}
				return db
			}
		}
	}
	log.Fatal("Problem with config file")
	panic("Problem with getting database connection")

}
func getRecords() []Dept {
	db := getDb()
	defer db.Close()
	rows, _ := db.Query("select * from dept")
	var depts []Dept = make([]Dept, 0)
	for rows.Next() {
		newdept := Dept{}
		if err := rows.Scan(&newdept.Deptno, &newdept.Dname, &newdept.Loc); err != nil {
			panic(err)
		}
		depts = append(depts, newdept)

	}
	return depts
}

func insertRecord(dept Dept) (err error) {
	db := getDb()
	defer db.Close()
	query := fmt.Sprintf("insert into dept values(%d, '%s','%s')", dept.Deptno, dept.Dname, dept.Loc)
	_, err = db.Exec(query)
	if err != nil {
		// query execution problem
		log.Fatal("Error Executing query")
		return
	}
	return
}
func main() {
	dept := Dept{11, "HR", "Hyd"}
	fmt.Println(dept)
	insertRecord(dept)
	for _, v := range getRecords() {
		fmt.Println(v)
	}
}
