package router

import (
	"any/api"
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

	var dbUser string = os.Getenv("USER")
	var dbPassword string = os.Getenv("PASSWORD")
	var dbName string = os.Getenv("NAME")
	var dbString string = dbUser+":"+dbPassword+"@/"+dbName+"?parseTime=true"
	var secret string = os.Getenv("secret")

func HomePage(c *gin.Context) {
	c.HTML(200, "home.html", gin.H{})
}

func LoginPage(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func RedirectLoginPage(c *gin.Context) {
	c.Header("HX-Redirect", "/loginPage")
}
func Login(c *gin.Context) {
	ctx := context.Background()
	db, err := sql.Open("mysql", dbString)
	if err != nil {
		c.String(500, "db err")
	}
	defer db.Close()
	queries := api.New(db)
	userName := c.PostForm("userName")
	hasher := md5.New()
	hasher.Write([]byte(c.PostForm("password")))
	hash := hex.EncodeToString(hasher.Sum(nil))
	result, err := queries.SelectUserById(ctx, userName)
	if err != nil {
		c.HTML(200, "error.html", gin.H{
			"errCode": 400,
			"err":     "Mighi be wrong",
		})
	}
	if result.Password == hash {
		sendToken(c, userName)
		c.Header("HX-Redirect", "/myProfile")
	} else {
		c.String(200, "wrong Password")
	}
}

func CreateUserPage(c *gin.Context) {
	c.HTML(200, "createUser.html", gin.H{})
}
func RedirectCreateUser(c *gin.Context) {
	c.Header("HX-Redirect", "createUserPage")
}
func CreateUser(c *gin.Context) {
	ctx := context.Background()
	db, err := sql.Open("mysql", dbString)
	if err != nil {
		c.String(500, "db err")
	}
	defer db.Close()
	queries := api.New(db)
	hasher := md5.New()
	hasher.Write([]byte(c.PostForm("password")))
	hash := hex.EncodeToString(hasher.Sum(nil))
	var temp api.CreateUserParams
	temp.Username = c.PostForm("username")
	temp.Color = c.PostForm("color")
	temp.Profession = c.PostForm("profession")
	temp.Bio = c.PostForm("bio")
	temp.Password = hash
	result, err := queries.CreateUser(ctx, temp)
	if err != nil {
		c.HTML(200, "error.html", gin.H{
			"errCode": 500,
			"err":     "somtehing went wrong with the database",
		})
	}
	if result != nil {
		sendToken(c, temp.Username)
		// send it back to the profile page of the user
		c.Header("HX-Redirect", "/myProfile")
	}
}

func sendToken(c *gin.Context, userName string) {
	//create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userName": userName,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	//hash to token with the secret
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.HTML(200, "error.html", gin.H{
			"errCode": 502,
			"err":     "The Cookie could not be set",
		})
	}

	//send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Token", tokenString, 3600*24*30, "", "", false, true)
}

func Middleware(c *gin.Context) {
	//get the cookie
	tokenString, err := c.Cookie("Token")
	// if you send response with an code of other than 200, the html wont be renderd.
	if err != nil {
		c.HTML(200, "error.html", gin.H{ // you send the login page here
			"errCode": 503,
			"err":     "You Need to login First",
		})
		return
	}
	//parse and validate the cookie the reaturn value is the secret
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) { return []byte(secret), nil })
	if err != nil {
		c.HTML(200, "error.html", gin.H{
			"errCode": 504,
			"err":     "Invalid Cookie Please Login Again.",
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if token.Valid && claims["userName"] != nil && ok {
		c.Set("userName", claims["userName"])
		c.Next()
	} else {
		c.HTML(200, "error.html", gin.H{
			"errCode": 500,
			"err":     "invalid cookie",
		})
	}
}

func CreateBlogPage(c *gin.Context) {
	c.HTML(200, "createBlog.html", gin.H{})
}
func CreateBlog(c *gin.Context) {
	ctx := context.Background()
	db, err := sql.Open("mysql", dbString)
	if err != nil {
		c.String(500, "db err")
	}
	defer db.Close()
	queries := api.New(db)
	user, _ := c.Get("userName")
	var temp api.CreateBlogParams
	temp.Title = c.PostForm("title")
	temp.Blogtext = c.PostForm("text")
	temp.Username = user.(string)
	result, err := queries.CreateBlog(ctx, temp)
	fmt.Println(result)
	if err != nil {
		c.HTML(200, "error.html", gin.H{
			"errCode": 500,
			"err":     "Something went Wrong with the database",
		})
	}
	c.Header("HX-Redirect", "/myProfile")
}

// deleting the cookie
func Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	//setting the age of the cookie as -1
	c.SetCookie("Token", "", -1, "", "", false, true)
	// c.String(200, "loged Out")
}

func LogoutSendHome(c *gin.Context) {
	Logout(c)
	c.Header("HX-Redirect", "/")
}

func UpdatedProfile(c *gin.Context) {
	ctx := context.Background()
	db, err := sql.Open("mysql", dbString)
	if err != nil {
		c.String(500, "db err")
	}
	defer db.Close()
	queries := api.New(db)
	var user api.UpdateUserParams
	user.Bio = c.PostForm("bio")
	user.Color = c.PostForm("color")
	user.Profession = c.PostForm("profession")
	user.Username = c.PostForm("username")
	user.Username_2 = c.PostForm("username2")
	res, err := queries.UpdateUser(ctx, user)
	if err != nil {
		c.HTML(200, "error.html", gin.H{
			"errCode": 500,
			"err":     "Something is wrong with the database",
		})
		return
	}
	if user.Username != user.Username_2 {
		Logout(c)
		sendToken(c, user.Username)
	}
	if res != nil {
		c.HTML(200, "myProfile.html", gin.H{})
	}
}

func UpdateBlog(c *gin.Context) {
	ctx := context.Background()
	db, err := sql.Open("mysql", dbString)
	if err != nil {
		c.String(500, "db err")
	}
	defer db.Close()
	queries := api.New(db)
	id := c.Param("id")
	intId, err := strconv.ParseInt(id, 10, 31)
	if err != nil {
		c.HTML(200, "error.html", gin.H{
			"errCode": 507,
			"err":     "wrond param value",
		})
	}
	var a api.UpdateBlogParams
	a.Blogid = int32(intId)
	a.Blogtext = c.PostForm("blogText")
	a.Title = c.PostForm("title")
	res, err := queries.UpdateBlog(ctx, a)
	if err != nil {
		c.HTML(200, "error.html", gin.H{
			"errCode": 507,
			"err":     "Someting webt wrong with  the  database.",
		})
	} else {
		c.HTML(200, "myProfile.html", gin.H{
			"res": res,
		})
	}
}

func UpdateBlogPage(c *gin.Context) {
	ctx := context.Background()
	db, err := sql.Open("mysql", dbString)
	if err != nil {
		c.String(500, "db err")
	}
	defer db.Close()
	queries := api.New(db)
	id := c.Param("id")
	intId, err := strconv.ParseInt(id, 10, 31)
	if err != nil {
		c.HTML(200, "error.html", gin.H{
			"errCode": 507,
			"err":     "wrond param value",
		})
	}
	res, err := queries.SelectBlogByUserName(ctx, int32(intId))
	if err != nil {
		c.HTML(200, "error.html", gin.H{
			"errCode": 503,
			"err":     "Something went wrong with the database",
		})
	} else {
		c.HTML(200, "updateBlog.html", gin.H{
			"id":    res.Blogid,
			"title": res.Title,
			"text":  res.Blogtext,
		})
	}
}

func DeleteBlog(c *gin.Context) {
	ctx := context.Background()
	db, err := sql.Open("mysql", dbString)
	if err != nil {
		c.String(500, "db err")
	}
	defer db.Close()
	queries := api.New(db)
	id := c.Param("id")
	intId, err := strconv.ParseInt(id, 10, 31)
	if err != nil {
		c.HTML(200, "error.html", gin.H{
			"errCode": 507,
			"err":     "wrond param value",
		})
	}
	e := queries.DeleteBlog(ctx, int32(intId))
	if e != nil {
		c.HTML(200, "error.html", gin.H{
			"errCode": 500,
			"err":     "Something is wrong with the database",
		})
	} else {
		c.HTML(200, "myProfile.html", gin.H{})
	}
}

func UpdateProfile(c *gin.Context) {
	ctx := context.Background()
	db, err := sql.Open("mysql", dbString)
	if err != nil {
		c.String(500, "db err")
	}
	defer db.Close()
	queries := api.New(db)
	user, _ := c.Get("userName")
	res, err := queries.SelectUserById(ctx, user.(string))
	if err != nil {
		c.HTML(200, "error.html", gin.H{
			"errCode": 500,
			"err":     "something went Wron with the database",
		})
	}
	c.HTML(200, "updateProfile.html", gin.H{
		"userName":   res.Username,
		"bio":        res.Bio,
		"color":      res.Color,
		"profession": res.Profession,
	})
}

func Blog(c *gin.Context) {
	ctx := context.Background()
	db, err := sql.Open("mysql", dbString)
	if err != nil {
		c.String(500, "db err")
	}
	defer db.Close()
	queries := api.New(db)
	id := c.Param("id")
	intId, err := strconv.ParseInt(id, 10, 31)
	if err != nil {
		c.HTML(200, "error.html", gin.H{
			"errCode": 507,
			"err":     "wrond param value",
		})
	}
	result, err := queries.SelectBlogByUserName(ctx, int32(intId))
	fmt.Println(result)
	res := []api.Blog{
		result,
	}
	if err != nil {
		c.HTML(200, "error.html", gin.H{
			"errCode": 502,
			"err":     "Something went wrong with the database",
		})
	}
	c.HTML(200, "blog.html", gin.H{
		"blog": res,
	})
}

func MyProfile(c *gin.Context) {
	ctx := context.Background()
	db, err := sql.Open("mysql", dbString)
	if err != nil {
		c.String(500, "db err")
	}
	defer db.Close()
	queries := api.New(db)
	user, _ := c.Get("userName")
	prof, err := queries.SelectUserById(ctx, user.(string))
	if err != nil {
		c.HTML(200, "error.html", gin.H{
			"errCode": 500,
			"err":     "Something went wrong with the database",
		})
	}
	res := []api.User{
		prof,
	}
	c.HTML(200, "profile.html", gin.H{
		"user": res,
		"mine": "true",
	})
	blogs, err := queries.SelectBlogsByUserName(ctx, user.(string))
	if err != nil {
		c.HTML(200, "error.html", gin.H{
			"errCode": 500,
			"err":     "Something is wrong with the Database",
		})
	}
	if blogs != nil {
		c.HTML(200, "myBlog.html", gin.H{
			"arr": blogs,
		})
	}
}

func Profile(c *gin.Context) {
	ctx := context.Background()
	db, err := sql.Open("mysql", dbString)
	if err != nil {
		c.String(500, "db err")
	}
	defer db.Close()
	queries := api.New(db)
	username := c.Param("user")
	result, err := queries.SelectUserById(ctx, username)
	if err != nil {
		c.HTML(200, "error.html", gin.H{
			"errCode": 500,
			"err":     "Someting went wrong with the database",
		})
	}
	res := []api.User{
		result,
	}
	c.HTML(200, "profile.html", gin.H{
		"user": res,
	})
}

func GeneralBlogs(c *gin.Context) {
	ctx := context.Background()
	db, err := sql.Open("mysql", dbString)
	if err != nil {
		c.String(500, "db err")
	}
	defer db.Close()
	queries := api.New(db)
	result, err := queries.SelectBlogs(ctx)
	fmt.Println(result)
	if err != nil {
		c.HTML(200, "error.html", gin.H{
			"errCode": 500,
			"err":     "Someting went wrong with the database.",
		})
	}
	if result != nil {
		c.HTML(200, "generalBlogs.html", gin.H{
			"arr": result,
		})
	}
}
