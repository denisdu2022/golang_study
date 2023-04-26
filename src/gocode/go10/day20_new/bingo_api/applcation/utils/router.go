package utils

import "github.com/gin-gonic/gin"

func Register(r *gin.RouterGroup, httpMethods []string, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {

	//路由注册函数{一次性给同一个视图绑定多个不同的HTTP请求方法}
	//定义routes
	var routes gin.IRoutes

	//循环
	for _, httpMethod := range httpMethods {
		//让定义的routes等于循环后的每一个的请求方式和路径还有handlers
		routes = r.Handle(httpMethod, relativePath, handlers...)
	}

	//然后返回这个循环后的routes
	return routes
}
