package ali

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"

)

func AliOss(fileName string, sourceBuffer io.Reader) string {
	/*setting.AliOss.Endpoint = "Endpoint"
	setting.AliOss.AccessKeyId = "AccessKeyId"
	setting.AliOss.AccessKeySecret = "AccessKeySecret"
	setting.AliOss.BucketName = "BucketName"*/
	client, _ := oss.New("setting.AliOss.Endpoint", "setting.AliOss.AccessKeyId", "setting.AliOss.AccessKeySecret")
	bucket, _ := client.Bucket("setting.AliOss.BucketName")
	err := bucket.PutObject("setting.AliOss.Env"+fileName, sourceBuffer)
	if err != nil {
		panic(err)
	}
	var bf bytes.Buffer
	bf.WriteString("setting.AliOss.PrivateServerDomain")
	bf.WriteString("/")
	bf.WriteString("setting.AliOss.Env")
	bf.WriteString(fileName)
	return bf.String()
}
