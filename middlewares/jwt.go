package middware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/scliang-strive/webServerTools/common/response"
	"strconv"
	"github.com/gin-gonic/gin"
)


const (
	SingKey        = ""
	JWTExpiresTime = 10000
)

type CustomClaims struct {
	UserId int64  `json:"user_id"`
	Name   string `json:"user_name"`
	Phone  string `json:"phone"`
	jwt.StandardClaims
}

// JwtSign 定义一个 JWT验签 结构体
type JwtSign struct {
	SigningKey []byte
}

type JWT struct {
	SigningKey []byte
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}

		j := CreateMyJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}

		if claims.ExpiresAt-time.Now().Unix() < JWTExpiresTime {
			claims.ExpiresAt = time.Now().Unix() + JWTExpiresTime
			newToken, _ := j.CreateToken(*claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))

		}
		c.Set("claims", claims)
		c.Next()
	}
}

// CreateMyJWT 使用工厂创建一个 JWT 结构体
func CreateMyJWT() *JwtSign {
	return &JwtSign{
		[]byte(SingKey),
	}
}

// CreateToken 生成一个token
func (j *JwtSign) CreateToken(claims CustomClaims) (string, error) {
	// 生成jwt格式的header、claims 部分
	tokenPartA := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 继续添加秘钥值，生成最后一部分
	return tokenPartA.SignedString(j.SigningKey)
}

// ParseToken 解析Token
func (j *JwtSign) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if token == nil {
		return nil, errors.New("")
	}
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New(ve.Error())
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New(ve.Error())
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// 如果 TokenExpired ,只是过期（格式都正确），我们认为他是有效的，接下可以允许刷新操作
				token.Valid = true
				goto labelHere
			} else {
				return nil, errors.New(ve.Error())
			}
		}
	}
labelHere:
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("")
	}
}

// RefreshToken 更新token
func (j *JwtSign) RefreshToken(tokenString string, extraAddSeconds int64) (string, error) {

	if CustomClaims, err := j.ParseToken(tokenString); err == nil {
		CustomClaims.ExpiresAt = time.Now().Unix() + extraAddSeconds
		return j.CreateToken(*CustomClaims)
	} else {
		return "", err
	}
}

