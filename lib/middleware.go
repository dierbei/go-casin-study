package lib

import (
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

func CheckLogin() gin.HandlerFunc{
	return func(context *gin.Context) {
	 	  if context.Request.Header.Get("token")==""{
				context.AbortWithStatusJSON(400,gin.H{"message":"token required"})
		  }else{
		  		 context.Set("user_name",context.Request.Header.Get("token"))
		  		 context.Next()
		  }
	}
}
func RBAC() gin.HandlerFunc  {
	e:= casbin.NewEnforcer("resources/model.conf","resources/p.csv")
	return func(context *gin.Context) {
		user,_:=context.Get("user_name")
		if !e.Enforce(user,context.Request.RequestURI,context.Request.Method){
			context.AbortWithStatusJSON(403,gin.H{"message":"forbidden"})
		}else{
			context.Next()
		}
	}
}
func Middlewares() (fs []gin.HandlerFunc)  {
	 fs=append(fs,CheckLogin(),RBAC())
	 return
}