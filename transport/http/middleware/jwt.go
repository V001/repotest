package middleware

import (
	"repotest/config"
	"repotest/service"
)

type JWTAuth struct {
	jwtKey []byte
	User   service.IUserService
}

func NewJWTAuth(cfg *config.Config, user service.IUserService) *JWTAuth {
	return &JWTAuth{jwtKey: []byte(cfg.JWTKey), User: user}
}

//
//func (m *JWTAuth) GenerateJWT(username string) (tokenString string, err error) {
//	expirationTime := time.Now().Add(1000 * time.Hour)
//	claims := &model.JWTClaim{
//		Name: username,
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: expirationTime.Unix(),
//		},
//	}
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	return token.SignedString(m.jwtKey)
//}
//func (m *JWTAuth) ValidateToken(signedToken string) (*model.JWTClaim, error) {
//	token, err := jwt.ParseWithClaims(
//		signedToken,
//		&model.JWTClaim{},
//		func(token *jwt.Token) (interface{}, error) {
//			return m.jwtKey, nil
//		},
//	)
//	if err != nil {
//		return nil, err
//	}
//	claims, ok := token.Claims.(*model.JWTClaim)
//	if !ok {
//		return nil, errors.New("couldn't parse claims")
//	}
//	if claims.ExpiresAt < time.Now().Local().Unix() {
//		return nil, errors.New("token expired")
//	}
//	return claims, nil
//}
//
//func (m *JWTAuth) ValidateAuth(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		token := extractToken(c.Request())
//		if token != "test" {
//			claims, err := m.ValidateToken(token)
//			if err != nil {
//				return echo.NewHTTPError(403, err.Error())
//			}
//
//			ctx := context.WithValue(c.Request().Context(), model.ContextUsername, claims.Name)
//			c.SetRequest(c.Request().WithContext(ctx))
//		}
//		return next(c)
//	}
//}
//
//func extractToken(r *http.Request) string {
//	bearToken := r.Header.Get("Authorization")
//	strArr := strings.Split(bearToken, " ")
//	if len(strArr) == 2 {
//		return strArr[1]
//	}
//	return ""
//}
//
//func (m *JWTAuth) ValidateActiveUser(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		token := extractToken(c.Request())
//		claims, err := m.ValidateToken(token)
//		if err != nil {
//			return echo.NewHTTPError(403, err.Error())
//		}
//		isVerify, err := m.User.IsVerified(c.Request().Context(), claims.Name)
//		if err != nil {
//			return err
//		}
//		if !isVerify {
//			return echo.NewHTTPError(http.StatusUnauthorized, errors.New("user is not verified"))
//		}
//		ctx := context.WithValue(c.Request().Context(), model.ContextUsername, claims.Name)
//		c.SetRequest(c.Request().WithContext(ctx))
//		return next(c)
//	}
//}
