package mongodb

import (
	"context"
	"fmt"
	"github.com/coderitx/webServerTools/config"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestMongoDb(t *testing.T) {
	conf := &config.MongoConf{
		Host:            "xx.xx.xx.xx",
		User:            "admin",
		Password:        "admin",
		MaxPoolSize:     1,
		MaxConnIdleTime: 5000,
		ConnectTimeout:  5000,
	}
	err := MongoInit(conf)
	if err != nil {
		fmt.Println("connect error:", err)
	}
	client := GetMongo("admin", "web")
	fmt.Println("cli", client)
	filer := bson.D{{"name", "golang"}}
	bytes, _ := client.FindOne(context.Background(), filer).DecodeBytes()
	fmt.Println(bytes.String())
}
