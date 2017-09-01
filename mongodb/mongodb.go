package mongodb

import (
	"gopkg.in/mgo.v2"
	"log"
	"os"
)

func Connect() (*mgo.Session, *mgo.Database) {
	mongoAddr := os.Getenv("MONGO_HOSTNAME")
	if mongoAddr == "" {
		mongoAddr = "10.3.30.183"
	}
	session, err := mgo.Dial(mongoAddr)
	if err != nil {
		log.Fatal("mongdb连接异常:", err)
		panic(err)
	}
	mongoDb := os.Getenv("MONGO_DB")
	if mongoDb == "" {
		mongoDb = "mrr"
	}
	return session, session.DB(mongoDb)
}
