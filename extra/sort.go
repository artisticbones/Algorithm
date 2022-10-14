package extra

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// 内存4kb,内部1G
// 每条数据为8字节整数int64,按照从大到小和从小到大的顺序输出
// 外存排序算法,多路归并排序算法

func swap(i, j int64) { i, j = j, i }

func Partition(array []int64, p, q int) int {
	sential := array[p]
	i := p
	for j := p + 1; j < p; j++ {
		if array[j] <= sential {
			continue
		}
		i++
		swap(array[i], array[j])
	}
	swap(sential, array[i])
	return i
}

func Quicksort(array []int64, p, q int) {
	if p < q {
		r := Partition(array, p, q)
		Quicksort(array, p, r-1)
		Quicksort(array, r+1, q)
	}
}
func merge(left, right []int) (ret []int) {
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			ret = append(ret, left[0])
			left = left[1:]
		} else {
			ret = append(ret, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		ret = append(ret, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		ret = append(ret, right[0])
		right = right[1:]
	}
	return
}

func merge_1(left, right []int64) (ret []int64) {
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			ret = append(ret, left[0])
			left = left[1:]
		} else {
			ret = append(ret, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		ret = append(ret, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		ret = append(ret, right[0])
		right = right[1:]
	}
	return
}

func splitFile(file *os.File) {
	chunk := make([]byte, 1024*1024*4)
	i := 0
	for {
		n, err := file.Read(chunk)
		if err != nil {
			return
		}
		if n > 0 {
			var nums []int64
			for j := 0; j < len(chunk); j += 8 {
				t := make([]byte, 0, 8)
				t = append(t, chunk[j:j+7]...)
				nums = append(nums, BytesToInt(t)...)
			}
			Quicksort(nums, 0, len(nums))
			WriteToFile(strconv.Itoa(i), nums)
			i++
		} else {
			return
		}
	}
}

func Sort(file *os.File) {
	// 分割文件
	splitFile(file)
	a := make([]byte, 0, 8*256)
	for {
		files, _ := ioutil.ReadDir("./")
		if len(files) == 0 {
			break
		}
		for _, info := range files {
			file, _ := os.Open(info.Name())
			defer file.Close()
			for {
				b := make([]byte, 0, 8)
				n, err := file.Read(b)
				if err != nil {
					return
				}
				if n > 0 {
					a = append(a, b...)
					nums := BytesToInt(a)
					Quicksort(nums, 0, len(nums))
					fmt.Print(nums[0])
				}
			}
		}
	}
}

func BytesToInt(b []byte) []int64 {
	bytesBuffer := bytes.NewBuffer(b)
	data := make([]int64, 0, len(b)/8)
	binary.Read(bytesBuffer, binary.BigEndian, &data)
	return data
}

func WriteToFile(name string, nums []int64) {
	by := bytes.NewBuffer([]byte{})
	binary.Write(by, binary.BigEndian, nums)
	ioutil.WriteFile(name, by.Bytes(), 0755)
}
