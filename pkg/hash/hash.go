package hash

import (
	"crypto/md5"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"server/pkg/logger"
	"strings"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/3/1 16:32
 * @Desc:
 */

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(password string) string {
	// GenerateFromPassword 的第二个参数是 cost 值。建议大于 12，数值越大耗费时间越长
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	logger.LogIf(err)

	return string(bytes)
}

// BcryptCheck 对比明文密码和数据库的哈希值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// BcryptIsHashed 判断字符串是否是哈希过的数据
func BcryptIsHashed(str string) bool {
	// bcrypt 加密后的长度等于 60
	return len(str) == 60
}

func BcryptMd5Hash(password string, salt string) string {
	w := md5.New()
	io.WriteString(w, password+salt)
	md5str2 := fmt.Sprintf("%x", w.Sum(nil))
	return md5str2
}

func BcryptIsMd5Hash(str string) bool {
	// bcrypt 加密后的长度等于 60
	return len(str) == 32
}

func BcryptCheckMd5(password string, hash string, salt string) bool {
	tmpPass := BcryptMd5Hash(password, salt)
	return strings.EqualFold(tmpPass, hash)
}
