package others

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"math"
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

// 归并排序
func Merge(in1, in2 <-chan int64) <-chan int64 {
	out := make(chan int64, 512)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
	}()
	return out
}

// 读数据。第一个参数是读的来源对象，第二个参数是读取长度(-1全读)。输出是一个channel
func ReaderSource(reader io.Reader, chunkSize int) <-chan int64 {
	out := make(chan int64, 512)
	go func() {
		// 64位系统的int型大小是64，所以用一个64位buffer = byte(8)*8
		buffer := make([]byte, 8)
		// 读取长度的控制变量
		bytesRead := 0
		for {
			// n是读取的长度
			n, err := reader.Read(buffer)
			bytesRead += n
			// 可能最后读取4字节数据，nil=EOF，所以要先读取数据，再判断nil
			if n > 0 {
				// 大端还是小端，发送和接收端统一即可
				out <- int64(binary.BigEndian.Uint64(buffer))
			}
			if err != nil ||
				(chunkSize != -1 && bytesRead >= chunkSize) {
				break
			}
		}
		close(out)
	}()
	return out
}

// 写数据。第一个参数是写的目的对象，第二个参数是写的数据channel
func WriteSink(writer io.Writer, in <-chan int64) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}

// 搭建归并节点组，递归调用实现2路归并
func MergeN(inputs ...<-chan int64) <-chan int64 {
	if len(inputs) == 1 {
		return inputs[0]
	}
	m := len(inputs) / 2
	// merge inputs[0..m) and inputs [m..end)
	return Merge(
		MergeN(inputs[:m]...),
		MergeN(inputs[m:]...))
}

func InMemSort(in <-chan int64) <-chan int64 {
	out := make(chan int64, 512)
	go func() {
		// Read into memory
		a := []int64{}
		for v := range in {
			a = append(a, v)
		}
		// Sort
		Quicksort(a, 0, len(a))
		// Output
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

// pipeline的搭建及运行，单机上时，分块数(chunkCount)最好是cpu的核数
func CreatePipeline(filename string, fileSize, chunkCount int) <-chan int64 {
	chunkSize := fileSize / chunkCount
	// 初始化结点组
	sortResults := []<-chan int64{}
	for i := 0; i < chunkCount; i++ {
		// 打开文件
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		// 设置offset
		file.Seek(int64(i*chunkSize), 0)
		// 读取数据
		source := ReaderSource(bufio.NewReader(file), chunkSize)
		// 内部排序后，追加到结点组中
		sortResults = append(sortResults, InMemSort(source))
	}

	// 归并结点组
	return MergeN(sortResults...)
}

// func WriteToFile(p <-chan int64, filename string) {
// 	// 创建文件
// 	file, err := os.Create(filename)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
// 	// 写入文件
// 	writer := bufio.NewWriter(file)
// 	defer writer.Flush()
// 	WriteSink(writer, p)
// }

func PrintFile(filename string) {
	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 读取数据
	p := ReaderSource(file, -1)
	count := 0
	fmt.Println("Sorted:")
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}

func Sort(file *os.File) {
	// 单机版
	// filename_prefix := "external_sort"
	// p := CreatePipeline(filename_prefix+".in", 1024*1024*1024, 256) // 文件大小(1024*1024)，读取块数(256)

	// WriteToFile(p, filename_prefix+".out")
	// PrintFile(filename_prefix + ".out")
	// 分割文件
	splitFile(file)
	for {
		nums := make([]int64, 0)
		a := make([]byte, 0, 8*256)
		files, _ := ioutil.ReadDir("./")
		if len(files) == 0 {
			break
		}
		for i := 0; i < int(math.Log2(float64(len(files)))); i++ {
			tmp := len(files)
			for j := 0; j < tmp; j += 2 {
				aReadBytes := 0
				file, _ := os.Open(strconv.Itoa(j))
				defer file.Close()
				a := make([]byte, 8)
				file.Seek(int64(4*1024*1024*j), 0)
				n, err := file.Read(a)
				if err != nil {
					return
				}
				aReadBytes += n
			}
			tmp /= 2
		}
		for _, info := range files {
			file, _ := os.Open(info.Name())
			defer file.Close()
			b := make([]byte, 0, 8)
			n, err := file.Read(b)
			if err != nil {
				return
			}
			if n > 0 {
				a = append(a, b...)
				nums = BytesToInt(a)
			}
		}
		Quicksort(nums, 0, len(nums))
		fmt.Print(nums[0])
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
