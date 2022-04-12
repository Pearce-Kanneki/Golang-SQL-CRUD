package main

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type data struct {
	id   int    `json:"id"`
	name string `json:"name"`
}

var session *mgo.Session
var coll *mgo.Collection

func main() {

}

// DB 登入
func login() {
	cred := &mgo.Credential{
		Username: "root",
		Password: "root",
	}
	err := session.Login(cred)
	if err != nil {
		log.Fatalln(err)
	}
}

// 連接DB
func linkDB() {
	s, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Fatalln(err)
	}

	session = s
	db := session.DB("test")
	coll = db.C("login.info")
}

// index 索引
func index(keys ...string) {
	index := mgo.Index{
		Key:        keys, // 索引鍵
		Unique:     true, // 唯一索引
		DropDups:   true, // 存在資料後創建,則自動刪除重複資料
		Background: true, // 不長時間佔用寫鎖
	}
	err := coll.EnsureIndex(index)
	if err != nil {
		log.Fatalln(err)
	}
}

// insert
func insert() {
	err := coll.Insert(&data{id: 1, name: "abc"})
	if err != nil {
		log.Fatalln(err)
	}
}

// query list
func query() {
	var list []data
	err := coll.Find(bson.M{"id": 1}).All(&list)
	if err != nil {
		log.Fatalln(err)
	}
}

// query one
func find() {
	var tmp data
	err := coll.Find(bson.M{"name": "test"}).Sort("id").One(&tmp)
	if err != nil {
		log.Fatalln(err)
	}
}

// update
func update() {
	findM := bson.M{"id": 1}
	updateM := bson.M{"name": "abc123"}

	_, err := coll.Upsert(findM, updateM)
	if err != nil {
		log.Fatalln(err)
	}
}

// delete
func remove() {
	findM := bson.M{"id": 1}

	err := coll.Remove(findM)
	if err != nil {
		log.Fatalln(err)
	}
}
