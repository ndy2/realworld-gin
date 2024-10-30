package auth

import "golang.org/x/crypto/bcrypt"

type Logic struct {
	repo Repo
}

// NewLogic 는 새로운 Logic을 생성하고 반환합니다.
func NewLogic(repo Repo) Logic {
	return Logic{repo: repo}
}

// Login 는 사용자를 인증하고 토큰을 반환합니다.
func (l Logic) Login(email string, password string) (string, error) {
	// 사용자가 존재하는지 확인합니다.
	user, err := l.repo.FindUserByEmail(email)
	if err != nil {
		return "", err
	}
	// 비밀번호가 일치하는지 확인합니다.
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	// TODO - 토큰을 생성하고 반환합니다.
	// token, err := generateToken(user)
	//if err != nil {
	//	return "", err
	//}
	return "token", nil
}
