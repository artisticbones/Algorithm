package extra

import (
	"encoding/binary"
	"fmt"
	"os"
)

func CountingSort(in []int64, k int) (out []int64) {
	counting := make([]int64, k+1)
	out = make([]int64, len(in))
	for i := range counting {
		counting[i] = 0
	}
	for i := range in {
		counting[in[i]]++
	}
	for i := 1; i < len(counting); i++ {
		counting[i] += counting[i-1]
	}
	for i := len(in) - 1; i >= 0; i-- {
		out[counting[in[i]]-1] = in[i]
		counting[in[i]]--
	}
	return
}

func calculateBucketCountWithFileSize(fileSize int, memorySize int) int {
	return fileSize / memorySize
}

func sort(file *os.File) {
	// 1.应当有的桶数量
	bucketCount := calculateBucketCountWithFileSize(1024*1024*1024, 4*1024)
	// 2.每个桶的数值范围
	begin, end := 0, bucketCount-1
	bytesRead := 0
	for i := 0; i < bucketCount; i++ {
		bucket := make([]int64, 0, 512)
		file, err := os.Open(file.Name())
		if err != nil {
			panic(err)
		}
		j := 0
		for {
			b := make([]byte, 8)
			file.Seek(int64(j*8), 0)
			n, err := file.Read(b)
			bytesRead += n
			if n > 0 {
				// 大端还是小端，发送和接收端统一即可
				n := int64(binary.BigEndian.Uint64(b))
				if n >= int64(begin) && n <= int64(end) {
					bucket = append(bucket, n)
				}
			}
			if err != nil || (bytesRead >= 4*1024) {
				// 3。达到内存容载最大值
				fmt.Println(CountingSort(bucket, end))
				begin += bucketCount
				end += bucketCount
				bytesRead = 0
				break
			}
			j++
		}
	}
}
