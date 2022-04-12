package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var tableName = "test"
var db *sql.DB

func main() {
	initDB()
	createDB(tableName)
	createTable()
	insert("aaa")
	// db.Ping() // Ping() 這裡才開始建立連線
}

func initDB() {
	dbConnect, err := sql.Open(
		"mysql",
		"root:root@tcp(127.0.0.1:3306)/",
		// "root:root@tcp(127.0.0.1:3306)/dbName" // 指定連接資料庫
	)

	if err != nil {
		log.Fatalln(err)
	}

	db = dbConnect // 用全域變數接

	db.SetMaxOpenConns(10)    // 設定最大DB連線數, 若<=0 則無上限
	db.SetConnMaxIdleTime(10) // 設置最大idle閒置連線數
}

// 建立資料庫
func createDB(dbName string) {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName + ";")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(err)
}

// 刪除資料庫
func deleteDB(dbName string) {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + dbName + ";")
	if err != nil {
		log.Fatalln(err)
	}
}

// 建立Table
func createTable() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS `test`.`table`(`name` VARCHAR(10))")
	if err != nil {
		log.Fatalln(err)
	}

	// 使用 USE 來指定資料庫
	// db.Exec("USE `test`")
	// _, err := db.Exec("CREATE TABLE IF NOT EXISTS `table`(`name` VARCHA(10))")
}

// 更改欄位
func alterTable() {
	_, err := db.Exec("ALTER TABLE `test`.`table` ADD `id` INT AUTO_INCREMENT PRIMARY KEY;")
	if err != nil {
		log.Fatalln(err)
	}
}

// 刪除Table
func deleteTable() {
	_, err := db.Exec("DROP TABLE IF NOT EXISTS `test`.`table`")
	if err != nil {
		log.Fatalln(err)
	}
}

// Insert
func insert(v string) {
	rs, err := db.Exec("INSERT INTO `test`.`table`(`name`,`id`) VALUES (?, 1)", v)
	if err != nil {
		log.Fatalln(err)
	}

	rowCount, err := rs.RowsAffected()
	rowId, err := rs.LastInsertId() // 資料表中友Auto_Increment欄位才起作用,回傳剛剛新增的ID

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("新增 %d 筆資料, id = %d \n", rowCount, rowId)
}

// Query
func query() {
	// db.Query => 回傳符合結果資料的多筆資料
	// db.QueryRow => 回傳符合結果資料的一筆資料
	rows, err := db.Query("SELECT * FROM `test`.`table`")
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var (
			tName string
			tid   int
		)

		err = rows.Scan(&tName, &tid)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%q %d\n", tName, tid)
	}

	defer rows.Close()
}

// Update
func update(id int, name string) {
	_, err := db.Exec("UPDATE `test`.`table` SET `name` = ? WHERE `id` = ?;", name, id)
	if err != nil {
		log.Fatalln(err)
	}
}

// Delete
func delete(id int) {
	_, err := db.Exec("DELETE FROM `test`.`table` WHERE `id` = ?;", id)
	if err != nil {
		log.Fatalln(err)
	}
}
