/*
 * Author: lihy lihy@zhiannet.com
 * Date: 2022-11-15 11:23:31
 * LastEditors: lihy lihy@zhiannet.com
 * Note: Need note condition
 */
package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	rndVerifyCode = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func SHA1(src ...string) string {
	var data []byte
	buf := bytes.NewBuffer(data)
	for _, v := range src {
		buf.WriteString(v)
	}
	h := sha1.New()
	h.Write(buf.Bytes())
	return hex.EncodeToString(h.Sum(nil))
}

// MD5 求值MD5码
func MD5(src string) string {
	h := md5.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}

func GenPassword(pwd string, salt string) string {
	return SHA1(pwd, salt)
}

func GenStoreSn() uint64 {
	source := rand.NewSource(time.Now().UnixNano())
	var max int64 = 999999999
	var min int64 = 100000000
	return uint64(rand.New(source).Int63n(max-min) + min)
}

func CheckPassword(pwd string, src string, salt string) bool {
	check := SHA1(pwd, salt)
	return strings.Compare(check, src) == 0
}

func GenVerifyCode() string {
	return fmt.Sprintf("%06v", rndVerifyCode.Int31n(1000000))
}

func VerifyPhone(phone string) bool {
	rule := `^1[3-9]\d{9}$`
	rgx := regexp.MustCompile(rule)
	return rgx.MatchString(phone)
}

func IsSameDay(t1 time.Time, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day()
}

func TrimJsonArray(s string) string {
	if len(s) >= 2 {
		if s[0] == '[' || s[len(s)-1] == ']' {
			s = s[1 : len(s)-1]
		}
	}
	return s
}

// type T []string|int64

func IsStrEmpty(strs ...string) bool {
	for _, str := range strs {
		if len(str) == 0 {
			return true
		}
	}
	return false
}

func GetMaxLenStr(strs ...string) string {
	m := 0
	i := 0
	for j, str := range strs {
		l := len(str)
		if l > m {
			m = l
			i = j
		}
	}
	return strs[i]
}

func GetMinLenStr(strs ...string) string {
	m := int(^uint(0) >> 1)
	i := 0
	for j, str := range strs {
		l := len(str)
		if l < m {
			m = l
			i = j
		}
	}
	return strs[i]
}

func CheckPasswordLevel(password string) string {
	level := 2
	length := len(password)
	var hasNum, hasSmall, hasBig bool
	if matched, _ := regexp.MatchString("[0-9]", password); matched {
		hasNum = true
	}
	if matched, _ := regexp.MatchString("[a-z]", password); matched {
		hasSmall = true
	}
	if matched, _ := regexp.MatchString("[A-Z]", password); matched {
		hasBig = true
	}
	if hasNum && hasSmall && hasBig {
		level = 3
	}
	if matched, _ := regexp.MatchString("[`~!@#$%^&*()_\\-+=<>?:\"{}|,.\\/;'\\\\[\\]·~！@#￥%……&*（）——\\-+={}|《》？：“”【】、；‘'，。、]", password); level == 3 && matched {
		level = 4
	}
	return fmt.Sprintf("%d-%d", level, length)
}
func CheckUserPasswordStrength(level string, policyLevel, policyLong uint8) bool {
	// level - length
	ls := strings.Split(level, "-")
	l1, _ := strconv.Atoi(ls[0])
	l2, _ := strconv.Atoi(ls[1])
	if policyLevel > uint8(l1) || (policyLevel >= 2 && policyLong > uint8(l2)) {
		return true
	}
	return false
}
