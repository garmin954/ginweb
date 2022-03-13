package dao

import (
	"ginweb/model"
	"ginweb/tools"
	"log"
)

type MemberDao struct {
	*tools.Orm
}

func (mbd MemberDao) InsertCode(sms model.SmsCode) int64 {
	r, err := mbd.InsertOne(&sms)
	if err != nil {
		log.Fatal(err.Error())
	}
	return r
}
