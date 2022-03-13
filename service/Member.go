package service

import (
	"fmt"
	"ginweb/dao"
	"ginweb/model"
	"ginweb/tools"
	"math/rand"
	"time"
)

type MemberService struct {
}

func (mbs *MemberService) SendSmsCode(phone string) bool {
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	smsCode := model.SmsCode{Phone: phone, Code: code, BizId: "3216165132", CreateAt: time.Now().Unix()}
	memberDao := dao.MemberDao{tools.DbEngin}
	result := memberDao.InsertCode(smsCode)
	return result > 0
}
