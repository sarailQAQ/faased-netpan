package utils

import (
	"cloud-disk/define"
	"context"
	"crypto/md5"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	uuid "github.com/satori/go.uuid"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"path"
	"time"
)

// AnalyzeToken 解析token
func AnalyzeToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("登录已失效")
	}
	return uc, err
}

// GenerateToken 生成token
func GenerateToken(id int, identity string, userName string, second int) (error, string) {
	claim := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(second)).Unix(),
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return err, ""
	}
	return nil, signedString
}

// Md5ToString 生成md5
func Md5ToString(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// GetUUID 获取UUID
func GetUUID() string {
	return uuid.NewV4().String()
}

func UploadFileToMinio(r *http.Request) (string, error) {
	c := context.Background()
	client, err := minio.New(define.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(define.AccessKeyID, define.SecretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln("创建Minio链接出现异常", err)
		return "", err
	}
	//校验存储桶是否存在
	exists, err := client.BucketExists(c, define.BucketName)
	if err != nil {
		log.Fatalln("查询存储桶状态异常", err)
		return "", err
	}
	//存储桶不存在进行创建
	if !exists {
		err = client.MakeBucket(c, define.BucketName, minio.MakeBucketOptions{Region: define.BucketLocation, ObjectLocking: false})
		if err != nil {
			log.Fatalln("创建存储桶异常", err)
			return "", err
		}
		//设置存储桶为公读写模式
		err := client.SetBucketPolicy(c, define.BucketName, fmt.Sprintf(define.BucketPolicy, define.BucketName, define.BucketName))
		if err != nil {
			log.Fatalln("修改存储桶权限异常", err)
			return "", err
		}
	}
	//进行文件上传
	formFile, header, err := r.FormFile("file")
	fileName := GetUUID() + path.Ext(header.Filename)
	_, err = client.PutObject(c, define.BucketName, fileName, formFile, header.Size, minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}
	return "http://" + define.Endpoint + "/" + define.BucketName + "/" + fileName, nil
}

// RandCode
//生成随机吗
func RandCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < define.CodeLength; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}

// MailSendCode 验证码发送
func MailSendCode(mail, code string) error {
	e := email.NewEmail()
	e.From = "水牛云盘 <your email>"
	e.To = []string{mail}
	e.Subject = "验证码发送测试"
	e.HTML = []byte(fmt.Sprintf("<pre style=\"font-family:Helvetica,arial,sans-serif;font-size:13px;color:#747474;text-align:left;line-height:18px\">欢迎使用水牛云盘，您的验证码为：<span style=\"font-size:block\">%s</span></pre>", code))
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "your-email", define.EmailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// FormatErrorLog 统一异常返回方法
func FormatErrorLog(err error) string {
	return fmt.Sprintf("服务发生异常-%s----请联系管理员", err.Error())
}

func GetFileType(ext string) string {
	for _, video := range define.Videos {
		if ext == video {
			return "video"
		}
	}
	for _, image := range define.Image {
		if ext == image {
			return "image"
		}
	}
	for _, doc := range define.Doc {
		if ext == doc {
			return "doc"
		}
	}
	return "other"
}
