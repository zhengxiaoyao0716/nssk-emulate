package nssk

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/zhengxiaoyao0716/util/requests"
)

var bKeys = map[string]string{}

// ReplyCreateConnect 应答创建连接请求
func ReplyCreateConnect(cbs, kbs string) (string, error) {
	// (3)A->B
	_content, err := Decrypt(cbs, kbs)
	if err != nil {
		return "", err
	}
	content := _content.(map[string]interface{})
	AppendLog(fmt.Sprintln("(3)A->B:", content))

	a := content["A"].(string)
	bKeys[a] = content["Kab"].(string)
	return a, nil
}

// VerifyConnect 发起验证连接请求
func VerifyConnect(a, b string) (string, error) {
	// (4)B->A
	nb := rand.Intn(65535)
	kab, ok := bKeys[a]
	if !ok {
		return "", errors.New("未找到ab间的通讯密钥")
	}
	delete(bKeys, a)
	cab := Encrypt(map[string]interface{}{"Nb": strconv.Itoa(nb)}, kab)
	data := map[string]interface{}{"Address": b, "Cab": cab}
	resp, err := requests.Post(
		a+"/api/a/connect/reply-verify",
		data,
	)
	if err != nil {
		return "", err
	}
	AppendLog(fmt.Sprintln("(4)B->A:", cab))

	// (5)A->B
	nbm1, err := Decrypt(resp.JSON()["body"].(string), kab)
	if err != nil {
		return "", err
	}
	if nbm1.(string) != strconv.Itoa(nb-1) {
		return "", errors.New("Nb错误")
	}
	AppendLog(fmt.Sprintln("(5)A->B:", nbm1))

	return kab, nil
}
