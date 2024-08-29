package main
 
import (
    "encoding/json"
    "fmt"
    "net/http"
)
 
func getCon(url, login, pass string) (db sql.Db){
	// load this from properties file 
	
	db, err := sql.Open("mysql", "admin:MyPassword@tcp(mydb.ctu244mmwtr1.us-east-1.rds.amazonaws.com:3306)/mydatabase1")
	if (err != nil){
		panic
	}
	// connection is established and database available 
}
func getRecords() []Dept {
	  db := getDB()
	  defer db.Close()
	  rows, _ := db.Query("select * from test.dept")
	  var depts []Dept = make([]Dept, 0)
	  for rows.Next() {
	    newdept := Dept{}
	    if err := rows.Scan(&newdept.Deptno, &newdept.Dname, &newdept.Loc); err != nil {
	      panic(err)->//log
	    }
	    depts = append(depts, newdept)
	    
	  }
	  return depts
	}
	
	func insertRecord(dept Dept) (err error) {
	  db := getDB()
	  defer db.Close()
	  query := fmt.Sprintf("insert into test.dept values(%d, '%s','%s')", dept.Deptno, dept.Dname, dept.Loc)
	  _, err := db.Exec(query)
	  if err != nil {
		// query execution problem 
		log ()
	    //fmt.Println("error connecting to db", err)
	    return 
	  }
	}
func main() {
 
    
    fmt.Println(db, err)
    defer db.Close()
    result, err := db.Exec("insert into dept values (2,'Fin', 'Hyd')")
    fmt.Println(result, err)
    rows, err := result.RowsAffected()
    fmt.Println(rows)
}
