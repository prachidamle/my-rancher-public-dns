package server

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

var (
	stmtIns *sql.Stmt
	stmtOut	*sql.Stmt
	
)
func init() {
	db, err := sql.Open("mysql", "root:@/rancher_public_dns")
	if err != nil {
	    panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. So use ping to test the connection
	err = db.Ping()
	if err != nil {
	    panic(err.Error()) // proper error handling instead of panic in your app
	}
//CREATE TABLE `rancher_public_dns`.`uuid_token` (
//  `id` INT NOT NULL AUTO_INCREMENT COMMENT '',
//  `uuid` VARCHAR(255) NOT NULL COMMENT '',
//  `token` VARCHAR(255) NOT NULL COMMENT '',
//  PRIMARY KEY (`id`)  COMMENT '');

}

func getConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@/rancher_public_dns")
	if err != nil {
	    return nil, err
	}

	// Open doesn't open a connection. So use ping to test the connection
	err = db.Ping()
	if err != nil {
	    return nil, err
	}
	
	return db, err
}


func insertUUIDTokenRecord(token string, uuid string) {
	db, err := getConnection()
	if err != nil {
	    panic(err.Error()) // proper error handling instead of panic in your app
	}
	// Prepare statement for inserting data
    stmtIns, err := db.Prepare("INSERT INTO uuid_token(uuid,token) VALUES( ?, ? )") // ? = placeholder
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
    
	res, err := stmtIns.Exec(token, uuid)
	if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
	id, err := res.LastInsertId()
	if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    fmt.Println(id)
}


func getUUID(token string) string {
	var uuid string 
	
	db, err := getConnection()
	if err != nil {
	    panic(err.Error()) // proper error handling instead of panic in your app
	}
	
	// Prepare statement for reading data
    stmtOut, err := db.Prepare("SELECT uuid FROM uuid_token WHERE token = ?")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer stmtOut.Close()

    // Query the uuid
    err = stmtOut.QueryRow(token).Scan(&uuid)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    fmt.Printf("The uuid is: %s", uuid)
    
    return uuid
}


