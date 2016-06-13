package main

import "code.google.com/p/go-sqlite/go1/sqlite3"
import "fmt"

type cache struct {
	ip       string
	hostname string
}

const (
	TABLE_LOCATION = "/tmp/cache-muls.db"
	TABLE   = "cache"
	colHost = "host"
	colIP   = "ip"
)

// func cacheInit() *sqlite3.Conn {
// 	sql := `CREATE TABLE cache(
// 				ip VARCHAR(255) PRIMARY KEY NOT NULL,
// 				host VARCHAR(255) NOT NULL
// 		    );`
// 	conn, err := sqlite3.Open(TABLE_LOCATION)
// 	if err != nil {
// 		return
// 	}
// 	err := conn.Exec(sql)
// }

func insert(conn *sqlite3.Conn, c cache) (err error) {
	args := sqlite3.NamedArgs{"$host": c.hostname, "$ip": c.ip}
	err = conn.Exec("INSERT INTO cache VALUES($ip, $host)", args)
	return err
}

// func selects(conn *sqlite3.Conn, c cache) (err error) {
// 	sql := "SELECT * FROM cache"
// }

func main() {
	// delCache := make(chan cache)
	// addCache := make(chan cache)
	c1, _ := sqlite3.Open("sqlite-test.db")
	err := c1.Exec(".schema cache")
	fmt.Println(err)
	// c := cache{"192.168.1.103", "cache103"}
	// err := insert(c1, c)
	// fmt.Println(err)
	// fmt.Println(c1)
    //
	// sql := "SELECT * FROM cache"
	// row := make(sqlite3.RowMap)
	// for s, err := c1.Query(sql); err == nil; err = s.Next() {
	// 	s.Scan(row)      // Assigns 1st column to rowid, the rest to row
	// 	fmt.Println(row) // Prints "1 map[a:1 b:demo c:<nil>]"
	// }
}

// create if not exist
func init() {
}
