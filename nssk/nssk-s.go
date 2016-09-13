package nssk

import (
	"fmt"
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
	secret := randKey(16)
	secrets[address] = secret
	AppendLog(fmt.Sprintln(address, "joined, secret: ", secret))
	return secret
}

// ConnectUser 连接用户
func ConnectUser(a, b, na string) string {
	data := map[string]interface{}{}
	data["Na"] = na
	data["B"] = b
	kab := randKey(16)
	data["Kab"] = kab
	cbs := map[string]interface{}{}
	cbs["Kab"] = kab
	cbs["A"] = a
	data["Cbs"] = Encrypt(cbs, secrets[b])
	AppendLog(fmt.Sprintln("(2)S->A:", data))
	AppendLog(fmt.Sprintln("(2)S->A: .Cbs=", cbs))
	return Encrypt(data, secrets[a])
}

func randKey(size int) string {
	// key := make([]byte, size)
	// if _, err := rand.Read(key); err != nil {
	// 	return ""
	// }
	// return string(key)
	kinds, result := [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		kind := rand.Intn(3)
		scope, base := kinds[kind][0], kinds[kind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}
