package user

import "fmt"

type User struct {
	Name string
	Age  int
}

// GetName はユーザーの名前を返します
func (u User) GetName() string {
	return u.Name
}

// GetAge はユーザーの年齢を返します
func (u User) GetAge() string {
	return fmt.Sprintf("%d歳", u.Age)
}

// 値レシーバのメソッド
func (u User) HaveBirthday() {
	u.Age = u.Age + 1
}

// ポインタレシーバのメソッド
func (u *User) HaveBirthdayPointer() {
	u.Age = u.Age + 1
}
