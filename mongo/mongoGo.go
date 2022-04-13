package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var (
	db         *mongo.Database
	collection *mongo.Collection
	client     *mongo.Client
)

type LogRecord struct {
	JobName string `bson:"jobName"`
	Command string `bson:"command"`
	Err     string `bson:"err"`
	Content string `bson:"content"`
}

// DB 連線
func initEngine() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 建立連接
	c, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalln(err)
	}

	client = c

	// 檢查連結
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// 選擇數據庫
	db = client.Database("my_db")

	// 選擇表
	collection = db.Collection("my_collection")
}

// insert
func insertOne() {
	lr := LogRecord{}
	iResult, err := collection.InsertOne(context.TODO(), lr)
	// ollection.InsertMany(context.TODO(), lr) // 多筆
	if err != nil {
		log.Fatalln(err)
		return
	}

	id := iResult.InsertedID.(primitive.ObjectID)
	fmt.Println("增加ID", id.Hex())
}

// find
func Find() {
	// 搜尋條件
	filter := bson.M{"JobName": "test"}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatalln(err)
		return
	}

	// 延遲關閉游標
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	//搜尋結果數據化
	for cursor.Next(context.TODO()) {
		var lr *LogRecord
		if cursor.Decode(&lr) != nil {
			log.Fatal(err)
			return
		}

		fmt.Println(lr)
	}

	// 搜尋結果數據化 另一種方法
	// var results []LogRecord
	// err = cursor.All(context.TODO(), &results)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _,result := range results{
	// 	fmt.Println(result)
	// }
}

// update
func updateMongo() {
	// 搜尋條件
	filter := bson.M{"JobName": "test"}
	// 修改Data
	update := bson.M{"Command": "byModel", "Content": "model"}

	uResult, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(uResult)

}

// delete
func deleteMongo() {
	// 搜尋條件
	filter := bson.M{"JobName": "test"}

	uResult, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(uResult)
}
