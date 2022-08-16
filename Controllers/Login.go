package Controllers

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"strconv"
)

type MailPW struct {
	Mail     string
	Password string
}

var store = sessions.NewCookieStore([]byte("sessioncontrol"))

func Login(c *gin.Context) {

	var user MailPW
	var userdb Models.User

	c.BindJSON(&user)
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(user.Password)))
	Config.DB.First(&userdb, "mail = ?", user.Mail)

	var sesm = user.Mail

	if userdb.Mail != user.Mail {
		c.JSON(400, "Sistemde kullanıcı bulunamadı.")
		c.Redirect(301, "/Login")
	}
	if userdb.Mail == user.Mail && userdb.Password != password {
		c.JSON(400, "Hatalı şifre girdiniz.")
		c.Redirect(301, "/Login")
	}
	if userdb.Mail == user.Mail && userdb.Password == password {

		session, _ := store.Get(c.Request, "sessioncontrol")
		session.Values["sessionmail"] = sesm
		session.Save(c.Request, c.Writer)
		var userr Models.User
		var childmentee Models.Childmentee
		Config.DB.Where("mail = ?", sesm).First(&userr)
		Config.DB.Where("user_id = ?", userr.ID).First(&childmentee)
		if childmentee.Active != true && childmentee.What != true {
			var mentee Models.Mentee
			Config.DB.Where("id = ?", childmentee.MenteeID).First(&mentee)
			mentee.MenteeCount += 1
			childmentee.Active = true
			childmentee.What = true
			Config.DB.Save(&mentee)
			Config.DB.Save(&childmentee)
		}
		c.JSON(202, userdb)
		//main := "/Company/"
		//url := fmt.Sprintf("%s%d", main, userdb.ID)
		//c.Redirect(301, url)

	}

}

func Logout(c *gin.Context) {
	session, err := store.Get(c.Request, "sessioncontrol")
	if err != nil {
		fmt.Println(err)
	}
	session.Values["sessionmail"] = ""
	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)
	//deneme := session.Values["sessionmail"]
	//fmt.Println(deneme)
	c.Redirect(200, "/Login")
}

func Logoutt(c *gin.Context) {
	number, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var user Models.User
	Config.DB.Where("id = ?", uint(number)).First(&user)
	user.IsDeleted = false
	Config.DB.Save(&user)
	c.JSON(200, "Çıkış başarılı.")
}
