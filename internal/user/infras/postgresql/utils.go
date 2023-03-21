package postgresql

import "golang.org/x/crypto/bcrypt"

func hashPassword(password string) (hashedPassword string) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
