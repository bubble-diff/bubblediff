package main

import (
	"errors"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"github.com/bubble-diff/bubblediff/config"
	"github.com/bubble-diff/bubblediff/models"
	"github.com/bubble-diff/bubblediff/oauth"
)

var (
	JwtAuthMws *jwt.GinJWTMiddleware
	CorsMws    gin.HandlerFunc
)

func InitMiddlewares() (err error) {
	conf := config.Get()

	// jwt鉴权中间件初始化
	JwtAuthMws, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm: "bubblediff",
		// 对称算法，用于签名
		SigningAlgorithm: "HS256",
		// 对称加密初始密钥
		Key: []byte(conf.ClientSecret),
		// jwt过期时间
		Timeout: time.Hour,
		// jwt最大延长有效时间
		MaxRefresh: time.Hour,
		// 生成jwt时，我们需要定义jwt的payload，为了能从jwt得知本次请求者，
		// 需要在payload声明该用户id。注意，claims必须存在jwt.IdentityKey这个key。
		PayloadFunc: func(user interface{}) jwt.MapClaims {
			if u, ok := user.(*models.User); ok {
				return jwt.MapClaims{jwt.IdentityKey: u.ID}
			}
			return jwt.MapClaims{}
		},
		// 授权函数，若为有效请求，应返回该用户身份id，后续作为Authorizator入参。
		Authenticator: oauth.GitHubOAuth,
		// 验证身份是否有效
		Authorizator: func(identity interface{}, c *gin.Context) bool {
			// 这是mapclaims的坑，由于json将所有数字当成浮点数，即使写入的值为int。
			// id := int(identity.(float64))
			// todo: 查询mongodb，判断该id是否存在，若不存在返回false
			return true
		},
		// 告知jwt token应该从请求的哪里获取
		TokenLookup: "header: Authorization",
		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",
		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	if err != nil {
		return errors.New("JWT Error:" + err.Error())
	}

	// CORS跨域资源共享中间件
	CorsMws = func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", conf.WebAddress)
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}

	return nil
}
