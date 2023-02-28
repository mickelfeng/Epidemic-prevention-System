package user_service

import (
	"context"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/query"
	myredis "github.com/pwh-pwh/Epidemic-prevention-System/dao/redis"
	"github.com/pwh-pwh/Epidemic-prevention-System/logic"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
	"github.com/pwh-pwh/Epidemic-prevention-System/service/menu_service"
	"log"
	"strings"
	"time"
)

func GetByUserName(username string) *models.SysUser {
	user := query.Use(mysql.DB).SysUser
	take, err := user.WithContext(context.Background()).Where(user.Username.Eq(username), user.Status.Eq(1)).Take()
	if err != nil {
		log.Printf("GetByUserName error:%v", err)
		return nil
	}
	redisClient := myredis.GetRedisClient()
	redisClient.Set(myredis.UserPre+username, take, 60*60*time.Second)
	return take
}

func GetUserAuthorityList(username string) []string {
	redisClient := myredis.GetRedisClient()
	sysUser := new(models.SysUser)
	err := redisClient.Get(myredis.UserPre + username).Scan(sysUser)
	if err != nil {
		log.Printf("GetUseAuthority redis get sysUser error:%v\n", err)
	}
	var authorityList []string
	result, err := redisClient.Get(myredis.GrantedAuthorityPre + sysUser.Username).Result()
	if err == nil {
		authorityList = strings.Split(result, ",")
	} else {
		sysRoles, err := logic.GetRoleByUserId(sysUser.ID)
		if err != nil {
			log.Printf("GetuserAuL error :%v\n", err)
			return nil
		}
		for _, r := range sysRoles {
			authorityList = append(authorityList, "ROLE_"+r.Code)
		}
		menuIds, err := logic.GetMenuIds(sysUser.ID)
		if err != nil {
			log.Printf("get menuids error:%v\n", err)
		}
		if len(menuIds) > 0 {
			sysMenus, err := menu_service.ListByIds(menuIds)
			if err != nil {
				log.Printf("menu listbyids error:%v\n", err)
			}
			for _, me := range sysMenus {
				authorityList = append(authorityList, me.Perms)
			}
		}
		authorityStr := strings.Join(authorityList, ",")
		_, err = redisClient.Set(myredis.GrantedAuthorityPre+sysUser.Username, authorityStr, time.Hour).Result()
		if err != nil {
			log.Printf("redisclient set authoritylist err:%v\n", err)
		}
	}
	return authorityList
}