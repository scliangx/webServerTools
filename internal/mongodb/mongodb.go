package mongodb

import (
	"context"
	"fmt"
	"github.com/scliangx/webServerTools/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var MClient *mongo.Client

func MongoInit(cfg *config.MongoConf) (e error) {
	// 连接uri
	uri := fmt.Sprintf("mongodb://%s:%s@%s/?connect=direct", cfg.User, cfg.Password, cfg.Host)
	// 构建mongo连接可选属性配置
	opt := new(options.ClientOptions)
	// 设置最大连接的数量
	opt = opt.SetMaxPoolSize(uint64(cfg.MaxPoolSize))
	// 设置连接超时时间
	opt = opt.SetConnectTimeout(time.Duration(cfg.ConnectTimeout) * time.Second)
	// 设置连接的空闲时间 毫秒
	opt = opt.SetMaxConnIdleTime(time.Duration(cfg.MaxConnIdleTime))
	// 开启驱动
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri), opt)
	if err != nil {
		return err
	}
	// 注意，在这一步才开始正式连接mongo
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	e = client.Ping(ctx, nil)
	if e != nil {
		return e
	}
	MClient = client
	return nil
}

func GetMongo(databaseName, collectionName string) *mongo.Collection {
	return MClient.Database(databaseName).Collection(collectionName)
}
