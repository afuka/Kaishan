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

var Srv *http.Server

// InitHttpSer 初始化 http server
func InitHttpSer(){

	gin.SetMode(gin.ReleaseMode)
	if conf.Viper.GetString("env") == "dev" {
		gin.SetMode(gin.DebugMode)
	}

	r := newRouter()
	Srv = &http.Server{
		Addr:    conf.Viper.GetString("https.port"),
		Handler: r,
	}

	go func() {
		// service connections
		if err := Srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("httpser init err, listen: " + err.Error())
		}
	}()

}

// Close 关闭
func Close()  {
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()
	if err := Srv.Shutdown(ctx); err != nil {
		log.Error("httpser shutdown err :" + err.Error())
	}
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