package httpser

import (
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"kaishan/core/handlers/conf"
	"kaishan/core/handlers/log"
	"kaishan/router"
	"net/http"
	"time"
)

// InitHttpSer 初始化 http server
func InitHttpSer(quit chan bool){

	gin.SetMode(gin.ReleaseMode)
	if conf.Viper.GetString("env") == "dev" {
		gin.SetMode(gin.DebugMode)
	}

	r := newRouter()
	srv := &http.Server{
		Addr:    conf.Viper.GetString("https.port"),
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Err.Errorf("httpser init err, listen: %s", err)
		}
	}()

	<-quit
	log.Info.Infof("httpser Shutdown ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Err.Errorf("httpser shutdown err : %s", err)
	}
	log.Info.Infof("httpser exiting ...")
}

// newRouter 路由配置
func newRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(sessionMiddleware())
	r.Use(corsMiddleware())

	// 路由
	api := r.Group("/api")
	{
		router.InitAPIRouter(api)
	}

	// 网页
	web := r.Group("/w")
	{
		router.InitWebRouter(web)
	}

	return r
}

// sessionMiddleware 初始化 session 和 cookie
func sessionMiddleware() gin.HandlerFunc {
	store := cookie.NewStore([]byte(conf.Viper.GetString("https.secret")))
	//Also set Secure: true if using SSL, you should though
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 7 * 86400, Path: "/"})
	return sessions.Sessions("gin-session", store)
}

// corsMiddleware 跨域
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,X-Token,X-UserId")
		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		// 处理请求
		c.Next()
	}
}