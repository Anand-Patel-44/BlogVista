package main

import (
	"any/router"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", router.Middleware, router.HomePage)
	r.GET("/loginPage", router.LoginPage)
	r.GET("/redirectLoginPage", router.RedirectLoginPage)
	r.POST("/login", router.Login)
	r.GET("/createUserPage", router.CreateUserPage)
	r.GET("/redirectCreateUser", router.RedirectCreateUser)
	r.POST("/createUser", router.CreateUser)
	r.GET("/createBlogPage", router.CreateBlogPage)
	r.POST("/blog", router.Middleware, router.CreateBlog)
	r.GET("/logout", router.Logout)
	r.GET("/generalBlogs", router.GeneralBlogs)
	r.GET("/profile/:user", router.Profile)
	r.GET("/myProfile", router.Middleware, router.MyProfile)
	r.GET("/blog/:id", router.Blog)
	r.GET("/updateProfile", router.Middleware, router.UpdateProfile)
	r.POST("/updatedProfile", router.UpdatedProfile)
	r.DELETE("/deleteBlog/:id", router.Middleware, router.DeleteBlog)
	r.GET("/updateBlogPage/:id", router.UpdateBlogPage)
	r.POST("/updateBlog/:id", router.UpdateBlog)
	r.GET("/logoutSendHome", router.LogoutSendHome)
	r.Run("localhost:8080")
}
