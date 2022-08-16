package Controllers

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"TREgitim/Repositories"
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type RegisteringUser struct {
	Username string
	Mail     string
	Password string
}

func Register(c *gin.Context) {
	var user Models.User
	var reguser RegisteringUser
	c.BindJSON(&reguser)

	Config.DB.First(&user, "mail=?", reguser.Mail)
	if user.Mail == reguser.Mail {
		c.JSON(400, "Bu mail adresi kullanılıyor.")
	}
	if user.Mail != reguser.Mail {
		password := fmt.Sprintf("%x", sha256.Sum256([]byte(reguser.Password)))
		user.Password = password
		user.UserName = reguser.Username
		user.Mail = reguser.Mail
		user.RegisterDate = time.Now()
		Repositories.NewUser(user)
		c.JSON(202, "Kayıt başarılı")
	}
}
