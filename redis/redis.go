package redis

import (
	"github.com/garyburd/redigo/redis"
	"net/http"
	"time"
	"log"
	"fmt"
	"io"
)

//链接池大小
var MAX_POOL_SIZE=20
var redisPoll chan redis.Conn

/**
添加redis链接池
 */
func putRedis(conn redis.Conn)  {
	//基于函数和接口间互不信任原则
	if redisPoll==nil {
		//初始化连接池
		redisPoll=make(chan redis.Conn,MAX_POOL_SIZE)
	}
	//链接池超过最大限制
	if len(redisPoll) >= MAX_POOL_SIZE {
		conn.Close()
		return
	}
	//输出链接到链接池中
	redisPoll <- conn
}

//初始化redis
func InitRedis(network, address string) redis.Conn {
	//缓冲机制，相当于消息队列
	if len(redisPoll) == 0{
		//如果长度为0，则初始化长度为MAX_POOL_SIZE的channel通道
		redisPoll=make(chan redis.Conn,MAX_POOL_SIZE)
		go func() {
			for i:=0;i < MAX_POOL_SIZE;i++{
				c,err:=redis.Dial(network,address)
				if err !=nil{
					panic(err)
				}
				putRedis(c)
			}
		}()
	}
	return <-redisPoll
}

func RedisServer(w http.ResponseWriter,r *http.Request)  {
	startTime:=time.Now()
	c:=InitRedis("tcp","127.0.0.1:6379")
	dbkey:="netgame:info"
	//向redis中写入数据
	if ok,err:=redis.Bool(c.Do("LPUSH",dbkey,"yanetao"));ok{
		//从redis中读取数据
		val,_:=c.Do("LPOP",dbkey)
		valMsg:=fmt.Sprintf("val:%s",val)
		io.WriteString(w,valMsg)
	}else{
		log.Print(err)
	}
	msg:=fmt.Sprintf("用时:%s",time.Now().Sub(startTime))
	io.WriteString(w,msg+"\n\n");
}