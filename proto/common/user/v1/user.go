package userv1

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (hashedPassword string) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func (u *UserORM) Authenticate(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))

	return err == nil
}
