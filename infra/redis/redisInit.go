package redis

import (
	"bookReadingNote/infra/tool/file/xmlFile"
	"bookReadingNote/infra/utils"
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

var (
	rdS  = new(RDSClient)
	once sync.Once
)

type RDSCle interface {
	setClient(rds *redis.Client)
	GetClient() *redis.Client
}

type RDSClient struct {
	rds *redis.Client
}

/*
	set redis client
*/
func (rC *RDSClient) setClient(rds *redis.Client) {
	// 只允许初始化时设置
	once.Do(func() {
		rC.rds = rds
	})
}

/*
	get redis client
*/
func (rC *RDSClient) GetClient() *redis.Client {
	return rC.rds
}

// connect hook
func connCallback(ctx context.Context, conn *redis.Conn) error {
	// 打印出该信息的时候表明redis连接成功
	ctxDT, ok := ctx.Deadline()
	fmt.Printf("Redis connected at: %v, Context deadline: %v, %v\n", utils.GetCurTimeUtc(), ctxDT, ok)
	return nil
}

func InitRedis(cfg xmlFile.XmlFile) {
	if rdCfg, ok := cfg.(*xmlFile.RDConfig); ok {
		rds := redis.NewClient(&redis.Options{
			Addr:        rdCfg.RdAddr,
			Password:    rdCfg.Pass,
			DB:          rdCfg.DbNum,
			PoolSize:    rdCfg.PoolSize,
			IdleTimeout: rdCfg.IdleTimeout * time.Second,
			OnConnect:   connCallback,
		})
		rdS.setClient(rds)
	} else {
		panic(errors.New("Redis init failed, cfg parameter error "))
	}
}

/*
	get already init struct RDSClient
*/
func GetRdS() *RDSClient {
	return rdS
}
