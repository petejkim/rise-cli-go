package signup

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/nitrous-io/rise-cli-go/client/users"
	"github.com/nitrous-io/rise-cli-go/pkg/readline"
	"github.com/nitrous-io/rise-cli-go/util"
)

func Signup(c *cli.Context) {
	var (
		err   error
		email string
	)

	for {
		email, err = readline.Read("Enter Email: ", true)
		util.ExitIfError(err)

		var password, passwordConf string

		readPw := func() {
			var err error
			password, err = readline.ReadSecurely("Enter Password: ", true)
			util.ExitIfError(err)

			passwordConf, err = readline.ReadSecurely("Confirm Password: ", true)
			util.ExitIfError(err)
		}

		readPw()
		for password != passwordConf {
			fmt.Println("Passwords do not match. Please re-enter password.")
			readPw()
		}

		appErr := users.Create(email, password)
		if appErr == nil {
			break
		}
		appErr.Handle()
		fmt.Println("There were errors in your input. Please try again.")
	}

	fmt.Println("Your account has been created. You will receive your confirmation code shortly.")

	for {
		confirmationCode, err := readline.Read("Enter Confirmation Code (Check your inbox!): ", true)
		util.ExitIfError(err)

		appErr := users.Confirm(email, confirmationCode)
		if appErr == nil {
			break
		}
		appErr.Handle()
	}

	fmt.Println("Thanks for confirming your email address! Your account is now active!")
}
