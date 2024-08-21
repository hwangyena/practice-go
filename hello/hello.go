package hello

import (
	"errors"
	"fmt"
	"math/rand"
)

// 대문자로 시작하면 export 된다
func Hello(name string) (string, error){
	if name==""{
		return "", errors.New("Empty")
	}

	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprintf("Hi, %v!", name)
	return message, nil
}

// func Hellos(name []string) (map[string]string, error){
// 	messages := make(map[string]string)

// 	for _, name := 
// }

func randomFormat() string {
	formats := []string{
		"Hi %v",
		"Hello %v",
		"안녕 %v",
	}

	return formats[rand.Intn(len(formats))]
}