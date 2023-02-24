package common

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/go-redis/redis"
	myredis "github.com/pwh-pwh/Epidemic-prevention-System/dao/redis"
	"log"
	"time"
)

const (
	// Standard width and height of a captcha image.
	StdWidth  = 200
	StdHeight = 50
)

//TODO 返回图片base64
func NewBase64Img() string {
	id := captcha.NewLen(6)
	fmt.Printf("id:%v\n", id)
	buffer := bytes.Buffer{}
	captcha.WriteImage(&buffer, id, StdWidth, StdHeight)
	data := buffer.Bytes()
	pre := "data:image/jpeg;base64,"
	dataStr := base64.StdEncoding.EncodeToString(data)
	res := pre + dataStr
	return res
}

func VerifyCaptcha(id, digits string) bool {
	return captcha.VerifyString(id, digits)
}

var redisCaptcha *RedisCaptcha = newRedisCaptcha()

func newRedisCaptcha() *RedisCaptcha {
	r := RedisCaptcha{
		myredis.GetRedisClient(),
	}
	/*reidsClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", "117.50.163.15", 6379),
	})
	r := RedisCaptcha{
		reidsClient,
	}*/
	captcha.SetCustomStore(r)
	return &r
}

type RedisCaptcha struct {
	client *redis.Client
}

func (r RedisCaptcha) Set(id string, digits []byte) {
	statusCmd := r.client.Set(myredis.CaptchaKey+":"+id, digits, time.Second*60*10)
	log.Printf("set bytes is:%v", digits)
	result, err := statusCmd.Result()
	if err != nil {
		log.Printf("rdiscaptcha set error:%v", err)
	} else {
		log.Println(result)
	}
}

//重写
func (r RedisCaptcha) Get(id string, clear bool) (digits []byte) {
	digits, err := r.client.Get(myredis.CaptchaKey + ":" + id).Bytes()
	if err != nil {
		log.Println(err)
	}
	if clear {
		r.client.Del(myredis.CaptchaKey + ":" + id)
	}
	return
}
