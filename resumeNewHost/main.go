package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"resumeNewHost/DB"
)

func main() {
	DB.Construct()
	r := gin.Default()

	CONFIG := cors.DefaultConfig()
	CONFIG.AllowOrigins = []string{"*"}
	CONFIG.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	CONFIG.AllowHeaders = []string{"Content-Type", "Authorization"}
	r.Use(cors.New(CONFIG))
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.POST("/login", func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		var userInfo struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		errBind := ctx.Bind(&userInfo)
		if errBind != nil {
			log.Fatal(errBind)
		}
		var user DB.User
		result := DB.DB.Where("username=? AND password=?", userInfo.Username, userInfo.Password).First(&user)
		if result.RowsAffected > 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "登录成功",
				"data":    user.UUID,
			})
		}
	})
	r.POST("/insert", func(ctx *gin.Context) {
		var newUser struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Email    string `json:"email"`
			Status   string `json:"status"`
		}
		errBind := ctx.Bind(&newUser)
		if errBind != nil {
			log.Fatal(errBind)
		}
		log.Println(newUser)
		data := DB.User{Username: newUser.Username, Password: newUser.Password, Email: newUser.Email}
		DB.DB.Create(&data)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data":    data,
		})
	})
	r.POST("/resume/insert", func(ctx *gin.Context) {
		var newResume struct {
			ResumeDataID string `json:"resume"`
			Title        string `json:"title"`
			Username     string `json:"username"`
			Subtitle     string `json:"subtitle"`
			Skills       string `json:"skills"`
			School       string `json:"school"`
			Project      string `json:"project"`
			Competition  string `json:"competition"`
		}
		errBind := ctx.Bind(&newResume)
		if errBind != nil {
			log.Fatal(errBind)
		}
		var userResume DB.UserResume
		var user DB.User
		result := DB.DB.Where("resume_id= ?", newResume.ResumeDataID).Find(&userResume)
		DB.DB.Where("UUID= ?", newResume.ResumeDataID).First(&user)
		fmt.Println(result)

		if result.Error != nil {
			// 处理其他数据库错误
			panic("Database error: " + result.Error.Error())
		} else if result.RowsAffected == 0 {
			newUserResume := DB.UserResume{
				ResumeID: newResume.ResumeDataID,
				UserID:   user.UUID,
			}
			DB.DB.Create(&newUserResume)
			fmt.Println("insert successfully")
		} else {
			fmt.Println("update successfully")
		}
		data := DB.Resume{
			ResumeDataID: newResume.ResumeDataID,
			Title:        newResume.Title,
			Username:     newResume.Username,
			Subtitle:     newResume.Subtitle,
			Skills:       newResume.Skills,
			School:       newResume.School,
			Project:      newResume.Project,
			Competition:  newResume.Competition,
		}
		DB.DB.Create(&data)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data":    data,
		})
	})
	errRun := r.Run(":3000")
	if errRun != nil {
		return
	}
}
