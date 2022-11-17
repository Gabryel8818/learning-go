package checkmail

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func Checkmail(email string) {
	mail := checkmail.ValidateFormat(email)
	if mail == nil {
		fmt.Println("Email valido")
	} else {
		fmt.Println("Email invalido")
	}
}
