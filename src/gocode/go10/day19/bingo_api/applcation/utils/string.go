package utils

import (
	"bingo_api/applcation/constants"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

func Random(n int, chars string) string {
	if n <= 0 {
		return ""
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, n, n)
	l := len(chars)
	for i := 0; i < n; i++ {
		bytes[i] = chars[r.Intn(l)]
	}
	return string(bytes)
}

func RandomLetters(n int) string {
	/**
	生成指定长度的字符串(字母)
	*/
	return Random(n, constants.LETTERS)
}

func RandomNumeric(n int) string {
	/**
	生成指定长度的字符串(数字)
	*/
	return Random(n, constants.NUMBERS)
}

func RandomLettersNumeric(n int) string {
	/**
	生成指定长度的字符串(数字+字母)
	*/
	return Random(n, constants.LETTERS_NUMBERIC)
}

func RandomAscii(n int) string {
	/**
	生成指定长度的字符串(字母+数字+特殊字符)
	*/
	return Random(n, constants.ASCII)
}

/*
 加密密码
*/

func MakeHashPassword(RawPassword string) (HashPassword string, err error) {
	pwd := []byte(RawPassword)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return
	}
	HashPassword = string(hash)
	return
}

/**
验证密码
*/

func CheckPassword(HashPassword string, RawPassword string) bool {
	ByteHash := []byte(HashPassword)
	BytePwd := []byte(RawPassword)
	err := bcrypt.CompareHashAndPassword(ByteHash, BytePwd)
	if err != nil {
		return false
	}
	return true
}
