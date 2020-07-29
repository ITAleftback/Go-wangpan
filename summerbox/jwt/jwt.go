package jwt

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Jwt struct {
}

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

func NewHeader() Header {
	return Header{
		Alg: "HS256",
		Typ: "JWT",
	}
}

type Payload struct {
	Iss      string `json:"iss"`
	Exp      string `json:"exp"`
	Iat      string `json:"iat"`
	Username string `json:"username"`
	Uid      uint
}

func Create(username string, id uint) string {
	header := NewHeader()
	payload := Payload{
		Iss:      "redrock",
		Exp:      strconv.FormatInt(time.Now().Add(3*time.Hour).Unix(), 10),
		Iat:      strconv.FormatInt(time.Now().Unix(), 10),
		Username: username,
		Uid:      id,
	}

	h, _ := json.Marshal(header)
	p, _ := json.Marshal(payload)
	headerBase64 := base64.StdEncoding.EncodeToString(h)
	payloadBase64 := base64.StdEncoding.EncodeToString(p)
	str1 := strings.Join([]string{headerBase64, payloadBase64}, ".")

	key := "redrock"
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(str1))
	s := mac.Sum(nil)

	signature := base64.StdEncoding.EncodeToString(s)
	token := str1 + "." + signature
	return token
}

func CheckToken(token string) ( err error) {
	arr := strings.Split(token, ".")
	if len(arr) != 3 {
		err = errors.New("token error")
		return
	}
	_, err = base64.StdEncoding.DecodeString(arr[0])
	if err != nil {
		err = errors.New("token error")
		return
	}
	pay, err := base64.StdEncoding.DecodeString(arr[1])
	if err != nil {
		err = errors.New("token error")
		return
	}
	sign, err := base64.StdEncoding.DecodeString(arr[2])
	if err != nil {
		err = errors.New("token error")
		return
	}

	str1 := arr[0] + "." + arr[1]
	key := []byte("redrock")
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(str1))
	s := mac.Sum(nil)
	fmt.Println(sign)
	fmt.Println(s)
	if res := bytes.Compare(sign, s); res != 0 {
		fmt.Println("test")
		err = errors.New("token error")
		return
	}

	var payload Payload
	json.Unmarshal(pay,&payload)

	return
}
