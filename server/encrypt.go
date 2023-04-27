package server

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"strconv"
	"strings"
	"time"

	"github.com/ddzwd/JumpServerSSHClient/instance"
)

func int_to_bytestring(i uint64) []byte {
	var result []byte
	for i > 0 {
		result = append(result, byte(i&0xff))
		i >>= 8
	}
	for len(result) < 8 {
		result = append(result, 0)
	}

	// 反转result
	for i := 0; i < len(result)/2; i++ {
		j := len(result) - i - 1
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func NewMFA(secret string) string {
	// 通过当前时间,计算出MFA秘钥
	return MFAHmacSha1(secret, uint64(time.Now().Unix()))
}

func MFAHmacSha1(secret string, time_value uint64) string {
	// base32解码
	decoded_bytes, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		instance.Logger.Fatalf("base32 decode error: %v", err)
	}
	instance.Logger.Debugf("decode_value: %v", decoded_bytes)
	instance.Logger.Debugf("now: %v", time_value)
	// 获取当前时间戳的字节
	now_btyes := int_to_bytestring(time_value / 30) // 保证30秒内有效
	instance.Logger.Debugf("now []bytes: %v", now_btyes)
	h := hmac.New(sha1.New, decoded_bytes)
	h.Write(now_btyes)
	mac_value := h.Sum(nil)
	instance.Logger.Debugf("mac_value: %v", mac_value)
	offset := mac_value[len(mac_value)-1] & 0xf // 取后8位
	instance.Logger.Debugf("offset: %v", offset)

	code := (uint64(mac_value[offset])&0x7f)<<24 | (uint64(mac_value[offset+1])&0xff)<<16 | (uint64(mac_value[offset+2])&0xff)<<8 | (uint64(mac_value[offset+3]) & 0xff)

	instance.Logger.Debugf("code: %v", code)
	str_code := strconv.FormatUint(code%1000000, 10)
	instance.Logger.Debugf("str_code: %v", str_code)

	left_len := 6 - len(str_code)
	if left_len <= 0 {
		return str_code
	}
	return strings.Repeat("0", left_len) + str_code

}
