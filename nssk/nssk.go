package nssk

import (
	"log"
	"math/rand"
	"time"
)

var users = []string{}
var secrets = map[string]string{}

// ListUser 列出用户列表
func ListUser() []string {
	return users
}

// JoinUser 加入新的用户
func JoinUser(address string) string {
	for _, user := range users {
		if user == address {
			return secrets[user]
		}
	}
	users = append(users, address)
	secret := randStr(32)
	secrets[address] = secret
	log.Println(secret)
	return secret
}
func randStr(size int) string {
	kinds, result := [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		kind := rand.Intn(3)
		scope, base := kinds[kind][0], kinds[kind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}
