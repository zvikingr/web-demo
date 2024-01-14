package logic

import (
	"web-demo/dao"
	"web-demo/utils/log"
	"web-demo/utils/trace"
)

type UserRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func UpdateUser(ctx trace.Context, username, password string) error {
	log.Infof("%s||update user", ctx)

	if err := dao.UpdateUser(username, password); err != nil {
		log.Errorf("%s||update user failed||err=%v", ctx, err)
		return err
	}

	return nil
}
