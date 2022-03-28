package account

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	"xiaohuazhu/internal/config"
	"xiaohuazhu/internal/model"
	"xiaohuazhu/internal/util/auth"
	"xiaohuazhu/internal/util/result"
)

// SignUp ...
func (s *Service) SignUp(ctx *gin.Context) {
	logrus.Infof("[account|SignUp] 开始注册")
	var param = model.AccountDTO{}
	if err := ctx.ShouldBindJSON(&param); err != nil {
		result.Fail(ctx, "参数错误")
		return
	}

	// 1. 判断用户是否存在
	user, err := s.accountDao.GetByUsername(param.Username, true)
	if err != nil {
		result.ServerError(ctx)
		return
	}
	if user != nil {
		logrus.Warnf("[account|SignUp] 该用户名已被注册, username: [%s]", param.Username)
		result.Fail(ctx, "该用户名已被注册")
		return
	}

	// 2. 密码加密 salt
	hash, err := bcrypt.GenerateFromPassword([]byte(param.Password+config.AllConfig.Application.Auth.PasswordSalt), bcrypt.DefaultCost)
	if err != nil {
		result.ServerError(ctx)
		return
	}
	var po = model.Account{
		Username: param.Username,
		Password: string(hash),
		// 给一个默认icon
		Icon: "image/test1.jpg",
	}
	// 保存
	if err := s.accountDao.Add(&po); err != nil {
		logrus.Errorf("[account|SignUp] 注册异常, err: [%+v]", err)
		result.Fail(ctx, err.Error())
		return
	}
	logrus.Infof("用户: [%s] 注册成功", param.Username)
	result.Success(ctx)
}

// SignIn in
func (s *Service) SignIn(ctx *gin.Context) {
	logrus.Infof("[account|SignIn] 用户登录")
	var param = model.AccountDTO{}
	if err := ctx.ShouldBindJSON(&param); err != nil {
		result.Fail(ctx, "参数错误")
		return
	}
	// 1. 用户有没有 && 用户密码是否正确
	user, err := s.accountDao.GetByUsername(param.Username, false)
	if err != nil {
		logrus.Errorf("[account|SignIn] 登录异常, err: [%+v]", err)
		result.ServerError(ctx)
		return
	}
	if user == nil {
		logrus.Errorf("[account|SignIn] 没有该用户 %s", param.Username)
		result.Fail(ctx, "用户名或密码错误")
		return
	}
	// 密码验证
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(param.Password+config.AllConfig.Application.Auth.PasswordSalt))
	if err != nil {
		logrus.Errorf("[account|SignIn] 没有该用户 %s", param.Username)
		result.Fail(ctx, "用户名或密码错误")
		return
	}

	// 2. 返回 jwt，client每次请求需要携带进行鉴权
	token, err := auth.GenerateToken(&model.AccountDTO{
		Id:       user.ID,
		Username: user.Username,
		Icon:     user.Icon,
	})
	if err != nil {
		logrus.Errorf("[account|SignIn] 生成 jwt 异常, err: [%+v]", err)
		result.ServerError(ctx)
		return
	}
	logrus.Infof("[account|SignIn] 用户登录成功: [%s]", param.Username)
	result.Ok(ctx, token)
}