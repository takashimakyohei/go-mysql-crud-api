package main

import "fmt"

// Speaker インターフェース - 話すことができるもの
type Speaker interface {
	Speak() string
}

// Dog 構造体
type Dog struct {
	Name string
}

// DogのSpeakメソッド実装
func (d Dog) Speak() string {
	return d.Name + "は「ワンワン」と鳴く"
}

// Cat 構造体
type Cat struct {
	Name string
}

// CatのSpeakメソッド実装
func (c Cat) Speak() string {
	return c.Name + "は「ニャーニャー」と鳴く"
}

// インターフェースを使用する関数
func MakeSpeak(s Speaker) {
	fmt.Println(s.Speak())
}
