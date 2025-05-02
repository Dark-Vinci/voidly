package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/dark-vinci/stripchat/beetle/utils"
	"github.com/dark-vinci/stripchat/beetle/utils/models"
)

type Tokens struct {
	AccessToken       string `json:"access_token"`
	AccessTokenExpiry string `json:"access_token_expiry"`
}

var claimsID = "user-id"
var claimsExpiry = "exp"
var claimsCreatedAt = "created_at"

func (m *Middleware) CreateToken(c *gin.Context, userID string) (*Tokens, error) {
	accessToken := jwtGo.New(jwtGo.GetSigningMethod(m.jwt.SigningAlgorithm))
	accessClaims := accessToken.Claims.(jwtGo.MapClaims)

	accessExpire := time.Now().Add(time.Hour * time.Duration(m.env.JWTExpiration))

	accessClaims[claimsID] = userID
	accessClaims[claimsExpiry] = accessExpire.Unix()
	accessClaims[claimsCreatedAt] = m.jwt.TimeFunc().Unix()

	accessTokenString, err := m.signedString(accessToken)
	if err != nil {
		return nil, err
	}

	c.SetCookie(
		userID,
		accessTokenString,
		int(time.Now().Add(time.Hour*time.Duration(m.env.JWTExpiration)).Unix()-time.Now().Unix()),
		"/",
		m.jwt.CookieDomain,
		m.jwt.SecureCookie,
		m.jwt.CookieHTTPOnly,
	)

	return &Tokens{
		AccessToken:       accessTokenString,
		AccessTokenExpiry: accessExpire.String(),
	}, err
}

func (m *Middleware) signedString(token *jwtGo.Token) (string, error) {
	return token.SignedString(m.jwt.Key)
}

func (m *Middleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := utils.GetRequestID(c.Request.Context())
		bearerToken := c.GetHeader("Authorization")

		rID, _ := uuid.Parse(requestID)

		if len(bearerToken) == 0 {
			utils.ErrorResponse(c, http.StatusUnauthorized, utils.ErrorData{
				ID:      rID,
				Details: "No authorization token provided",
				Status:  http.StatusUnauthorized,
			})
			return
		}

		if !strings.HasPrefix(bearerToken, "Bearer ") {
			utils.ErrorResponse(c, http.StatusUnauthorized, utils.ErrorData{
				ID:      rID,
				Handler: packageName,
				Details: utils.ErrInvalidTokenHeaderFormat.Error(),
			})
			return
		}

		token := strings.TrimPrefix(bearerToken, "Bearer ")

		userID, err := m.parseToken(token)
		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, utils.ErrorData{
				ID:      rID,
				Handler: packageName,
				Details: err.Error(),
			})
		}

		parseUserID, _ := uuid.Parse(userID)

		if parseUserID == uuid.Nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, utils.ErrorData{
				ID:      rID,
				Handler: packageName,
				Details: "Invalid session or unauthorized user",
			})
			return
		}

		ctx := models.CTX{
			Context: c.Request.Context(),
		}

		user, err := m.app.GetUserByID(ctx, parseUserID)
		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, utils.ErrorData{
				ID:      rID,
				Handler: packageName,
				Details: err.Error(),
			})
		}

		ctx.RequestID = rID
		ctx.UserID = parseUserID
		ctx.Email = user.Email

		c.Set(utils.CTX, ctx)
	}
}

func (m *Middleware) parseToken(tokenStr string) (string, error) {
	token, err := jwtGo.Parse(tokenStr, func(token *jwtGo.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtGo.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.env.JWTKey), nil
	})

	if token == nil {
		m.logger.Error().Str("token", tokenStr).Msg("unable to parse token - token is most likely not valid")
		return "", utils.ErrInvalidToken
	}

	if claims, ok := token.Claims.(jwtGo.MapClaims); ok && token.Valid {
		userID := claims[claimsID].(string)

		return userID, nil
	}

	return "", err
}

func (m *Middleware) Default() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (m *Middleware) GenerateCredentials() string {
	return ""
}
