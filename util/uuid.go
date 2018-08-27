package util

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	hostname, _ = os.Hostname()
)

// GenUploadUUID tsStr_1 + md5(tsStr+hostname+randint)[:16] + tsStr_2
func GenUploadUUID() string {
	tsStr := fmt.Sprintf("%x", time.Now().UnixNano())
	tmpStr := MD5([]byte(fmt.Sprintf("%s%s%d", tsStr, hostname, rand.Intn(1000))))
	return tsStr[0:8] + tmpStr[:16] + tsStr[8:16]
}
