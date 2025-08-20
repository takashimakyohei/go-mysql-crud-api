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