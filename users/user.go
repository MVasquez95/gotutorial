package users

import (
	"fmt"
	"time"

	"github.com/gotutorial/models"
)

func AltaUsuario() {
	u := new(models.User)
	u.AddUser(13, "Bolt", time.Now(), true)
	fmt.Println(u)
}
