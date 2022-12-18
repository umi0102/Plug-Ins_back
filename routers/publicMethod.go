package routers

import (
	"Plug-Ins/databases/mysql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"unsafe"
)

func Byte2Str(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

func Str2Byte(s string) []byte {
	tmp1 := (*[2]uintptr)(unsafe.Pointer(&s))
	tmp2 := [3]uintptr{tmp1[0], tmp1[1], tmp1[1]}
	return *(*[]byte)(unsafe.Pointer(&tmp2))
}

func Str2Map(jsonData string) (result map[string]interface{}, err error) {
	err = json.Unmarshal([]byte(jsonData), &result)
	return result, err
}

func Map2Str(mapData map[string]interface{}) (result string, err error) {
	resultByte, errError := json.Marshal(mapData)
	result = string(resultByte)
	err = errError
	return result, err
}

// GetImage 获取图片的Base64
func GetImage(path string) (baseImg string, err error) {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("2222222222222222222222")
	}
	mimeType := http.DetectContentType(file)
	switch mimeType {
	case "image/jpeg":
		baseImg = "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(file)
	case "image/png":
		baseImg = "data:image/png;base64," + base64.StdEncoding.EncodeToString(file)
	}
	return
}

// UploadImage  上传图片
func UploadImage(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("files")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "400", "message": err.Error()})
		return
	}

	fileExt := filepath.Ext(fileHeader.Filename)
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" {
		get, _ := ctx.Get("phone")

		fileDir := "./public/upload/images/usericon/"

		// fileDb := fmt.Sprintf("public/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
		err = os.MkdirAll(fileDir, 0777)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": "400", "message": err.Error()})
			return
		}

		fileName := fmt.Sprintf("%s%s", get, fileExt)
		filePathStr := filepath.Join(fileDir, fileName)
		err1 := ctx.SaveUploadedFile(fileHeader, filePathStr)
		if err1 != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": "400", "message": err.Error()})
			return
		}

		imgDir := fmt.Sprintf("%s%s%s", fileDir, get, fileExt)
		mysql.InsUpdDelMysql(fmt.Sprintf(`UPDATE userinfos SET userinfo_usericon = "%s" WHERE userinfo_phone="%s"`, imgDir, get))
		ctx.JSON(200, gin.H{
			"status":   "200",
			"filename": fileHeader.Filename,
		})
	}

}

// RandCreator 生产随机数
func RandCreator(l int) string {
	str := "0123456789abcdefghigklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+"
	strList := []byte(str)

	var result []byte
	i := 0

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i < l {
		newStr := strList[r.Intn(len(strList))]
		result = append(result, newStr)
		i = i + 1
	}
	return string(result)
}
