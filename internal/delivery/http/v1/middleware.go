package v1

import (
	"github.com/AZRV17/Skylang/pkg/auth"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

//func (h *Handler) userIdentity() func(next http.Handler) http.Handler {
//	return func(next http.Handler) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			manager, err := auth.NewManager(h.config.Auth.SecretKey)
//
//			token, err := r.Cookie("accessToken")
//			if err != nil {
//				tokenString, err := manager.NewJWT("1", "admin", time.Duration(time.Hour))
//				if err != nil {
//					w.WriteHeader(http.StatusInternalServerError)
//					c.JSON(w, r, map[string]interface{}{
//						"message": err.Error(),
//					})
//				}
//
//				log.Println("New token")
//
//				http.SetCookie(w, &http.Cookie{
//					Name:    "accessToken",
//					Value:   tokenString,
//					Expires: time.Now().Add(time.Hour),
//				})
//
//				token, _ = r.Cookie("accessToken")
//			}
//
//			log.Println(token)
//
//			if err != nil {
//				w.WriteHeader(http.StatusInternalServerError)
//				c.JSON(w, r, map[string]interface{}{
//					"message": err.Error(),
//				})
//
//				return
//			}
//
//			claims, _, _, err := manager.Parse(token.Value)
//			if err != nil {
//				w.WriteHeader(http.StatusUnauthorized)
//				c.JSON(w, r, map[string]interface{}{
//					"message": "unauthorized",
//				})
//
//				return
//			}
//
//			log.Println(claims)
//
//			next.ServeHTTP(w, r)
//		})
//	}
//}

func (h *Handler) userIdentity() func(c *gin.Context) {
	return func(c *gin.Context) {
		manager, err := auth.NewManager(h.config.Auth.SecretKey)
		token, err := c.Cookie("accessToken")
		if err != nil {
			tokenString, err := manager.NewJWT("1", "admin", time.Duration(time.Hour))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
			}

			log.Println("New token")

			http.SetCookie(c.Writer, &http.Cookie{
				Name:    "accessToken",
				Value:   tokenString,
				Expires: time.Now().Add(time.Hour),
			})

			token, err = c.Cookie("accessToken")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
			}
		}

		log.Println(token)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})

			return
		}

		claims, _, _, err := manager.Parse(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})

			return
		}

		log.Println(claims)

		c.Next()
	}
}
