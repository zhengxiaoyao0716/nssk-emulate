package nssk

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/zhengxiaoyao0716/util/requests"
)

var aKeys = map[string]string{}

// CreateConnect 发起创建连接请求
func CreateConnect(s, a, b, kas string) error {
	// (1)A->S
	na := strconv.Itoa(rand.Intn(65535))
	data := map[string]interface{}{"From": a, "To": b, "Na": na}
	resp, err := requests.Post(
		s+"/api/s/user/connect",
		data,
	)
	if err != nil {
		return err
	}
	AppendLog(fmt.Sprintln("(1)A->S:", data))

	// (2)S->A
	_content, err := Decrypt(resp.JSON()["body"].(string), kas)
	if err != nil {
		return err
	}
	content := _content.(map[string]interface{})
	b = content["B"].(string)
	kab := content["Kab"].(string)
	AppendLog(fmt.Sprintln("(2)S->A:", content))

	// (3)A -> B
	if content["Na"].(string) != na {
		return errors.New("随机数Na校验错误，与服务器建立连接失败")
	}
	cbs := content["Cbs"].(string)
	data = map[string]interface{}{"Cbs": cbs}
	_, err = requests.Post(
		b+"/api/b/connect/reply-create",
		data,
	)
	if err != nil {
		return err
	}
	AppendLog(fmt.Sprintln("(3)A->B:", data))

	aKeys[b] = kab
	return nil
}

// ReplyVerifyConnect 应答验证连接请求
func ReplyVerifyConnect(a, b, cab string) (string, string, error) {
	// (4)B->A
	kab, ok := aKeys[b]
	delete(aKeys, b)
	_content, err := Decrypt(cab, kab)
	if err != nil {
		return "", "", err
	}
	content := _content.(map[string]interface{})
	AppendLog(fmt.Sprintln("(4)B->A:", content))

	// (5)A->B
	nb, err := strconv.Atoi(content["Nb"].(string))
	if err != nil {
		return "", "", errors.New("Nb不合法")
	}
	if !ok {
		return "", "", errors.New("未找到ab间的通讯密钥")
	}
	AppendLog(fmt.Sprintln("(5)A->B:", nb))
	return Encrypt(strconv.Itoa(nb-1), kab), kab, nil
}
