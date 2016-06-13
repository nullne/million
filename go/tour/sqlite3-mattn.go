package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"sync"
	// "os"
)

type cache struct {
	host string
	ip   string
}

const (
	DB_LOCATION = "/tmp/cache-ris.db"
	TABLE       = "cache"
	COL_HOST    = "host"
	COL_IP      = "ip"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	db, err := sql.Open("sqlite3", DB_LOCATION)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Query("select * from cache limit 1")
	if err != nil {
		sqlStmt := `
		CREATE TABLE cache(
			ip VARCHAR(255) PRIMARY KEY NOT NULL,
			host VARCHAR(255) UNIQUE NOT NULL
		);
		`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			log.Printf("%q: %s\n", err, sqlStmt)
			return
		}
	} else {
		fmt.Println("No error on init")
	}
}

// not transaction for some reason
func inserts(db *sql.DB, caches chan cache, wg *sync.WaitGroup) {
	for c := range caches {
		_, err := db.Exec("INSERT INTO "+TABLE+"  VALUES (?, ?)", c.ip, c.host)
		checkErr(err)
	}
	wg.Done()
}

func deletes(db *sql.DB, caches chan cache, wg *sync.WaitGroup) {
	for c := range caches {
		_, err := db.Exec("DELETE FROM "+TABLE+" WHERE ip = ? OR host = ?", c.ip, c.host)
		checkErr(err)
	}
	wg.Done()
}

func selects(db *sql.DB, host string) (string, error) {
	stmt, err := db.Prepare("SELECT * FROM cache WHERE host = ?")
	checkErr(err)
	defer stmt.Close()
	var ip string
	err = stmt.QueryRow(host).Scan(&ip)
	return ip, err
}

func main() {
	db, err := sql.Open("sqlite3", DB_LOCATION)
	checkErr(err)
	defer db.Close()
	ip, err := selects(db, "localhost-2")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(ip)
	}

	/*
		var wg sync.WaitGroup
		caches := make(chan cache, 10)
		wg.Add(1)
		caches <- cache{"192.168.1.1",""}
		caches <- cache{"","localhost-2"}
		caches <- cache{"192.168.1.3", "localhost-3"}
		close(caches)
		go deletes(db, caches, &wg)
		wg.Wait()
	*/
	/*
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}
		stmt, err := tx.Prepare("insert into foo(host, ip) values(?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		for i := 0; i < 100; i++ {
			_, err = stmt.Exec(i, fmt.Sprintf("こんにちわ世界%03d", i))
			if err != nil {
				log.Fatal(err)
			}
		}
		tx.Commit()

		rows, err := db.Query("select id, name from foo")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var id int
			var name string
			rows.Scan(&id, &name)
			fmt.Println(id, name)
		}

		stmt, err = db.Prepare("select name from foo where id = ?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		var name string
		err = stmt.QueryRow("3").Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)

		_, err = db.Exec("delete from foo")
		if err != nil {
			log.Fatal(err)
		}

		_, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
		if err != nil {
			log.Fatal(err)
		}

		rows, err = db.Query("select id, name from foo")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var id int
			var name string
			rows.Scan(&id, &name)
			fmt.Println(id, name)
		}
	*/
}
