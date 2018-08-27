package mpupload

import (
	"math"

	"github.com/moxiaomomo/filestore-cloud/util"
)

// 上传初始化返回的数据
type UploadInitInof struct {
	FileSize    int64
	UploadID    string
	ChunkSize   int64
	ChunkCount  int64
	ChunkIDFrom int64
	UploadHost  string
}

// 初始化分块上传
func InitiateMultipartUpload(filesize int64) UploadInitInof {
	upInfo := UploadInitInof{
		FileSize:    filesize,
		UploadID:    util.GenUploadUUID(),
		ChunkIDFrom: 1,
		ChunkSize:   5242880,
		ChunkCount:  1,
		UploadHost:  "http://upload.test.com",
	}

	if filesize > 5242880 {
		upInfo.ChunkCount = int64(math.Ceil(float64(filesize / 5242880)))
	}

	return upInfo
}

// 上传分块
func UploadPart() {

}

// 取消上传
func CancelUploadPart() {

}

// 完成上传，通知合并
func CompleteUploadPart() {

}

// 获取已上传的分块信息
func GetUploadedInfo() {

}
