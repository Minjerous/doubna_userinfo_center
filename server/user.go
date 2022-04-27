package server

import (
	"doubna_userinfo/dao"
	"doubna_userinfo/model"
	"doubna_userinfo/tool"
	"strings"
)

func JudgePasswordCorrect(loginAccount, password string) (model.User, bool, error) {
	// 判断登录类型
	flag := tool.VerifyEmailFormat(strings.ToLower(loginAccount))
	if flag {
		//邮箱登录
		user, err := dao.QueryByEmail(loginAccount)
		//判断错误类型
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				return model.User{}, false, nil
			}
			return model.User{}, false, err
		}
		//判断密码
		if tool.Match(user.Password, password, user.Salt) {
			return user, true, nil
		} else {
			return model.User{}, false, nil
		}
	} else {
		//手机号登录
		user, err := dao.QueryByPhone(loginAccount)
		//判断错误类型
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				return model.User{}, false, nil
			}
			return model.User{}, false, err
		}
		//判断密码
		if tool.Match(user.Password, password, user.Salt) {
			return user, true, nil
		} else {
			return model.User{}, false, nil
		}
	}
}

