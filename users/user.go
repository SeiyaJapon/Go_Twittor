package users

import (
	"fmt"
	"time"

	"github.com/SeiyaJapon/golang/go_twittor/models"
)

func SingUpUser() {
	user := new(models.User)

	user.AddUser(10, "Pepe", time.Now(), true)

	fmt.Println(user)
}
