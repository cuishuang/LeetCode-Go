package util

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// LoadSolutionsDir define
func LoadSolutionsDir() ([]int, int) {
	files, err := ioutil.ReadDir("../leetcode/")
	if err != nil {
		fmt.Println(err)
	}
	solutionIds := []int{}
	for _, f := range files {
		if f.Name()[4] == '.' {
			tmp, err := strconv.Atoi(f.Name()[:4])
			if err != nil {
				fmt.Println(err)
			}
			solutionIds = append(solutionIds, tmp)
		}
	}
	sort.Ints(solutionIds)
	fmt.Printf("读取了 %v 道题的题解，当前目录下有 %v 个文件(可能包含 .DS_Store)，目录中有 %v 道题在尝试中\n", len(solutionIds), len(files), len(files)-len(solutionIds))
	return solutionIds, len(files) - len(solutionIds)
}

// LoadChapterFourDir define
func LoadChapterFourDir() []string {
	files, err := ioutil.ReadDir("../website/content/ChapterFour/")
	if err != nil {
		fmt.Println(err)
	}
	solutions, solutionIds, solutionsMap := []string{}, []int{}, map[int]string{}
	for _, f := range files {
		if f.Name()[4] == '.' {
			tmp, err := strconv.Atoi(f.Name()[:4])
			if err != nil {
				fmt.Println(err)
			}
			solutionIds = append(solutionIds, tmp)
			// len(f.Name())-3 = 文件名去掉 .md 后缀
			solutionsMap[tmp] = f.Name()[:len(f.Name())-3]
		}
	}
	sort.Ints(solutionIds)
	fmt.Printf("读取了第四章的 %v 道题的题解\n", len(solutionIds))
	for _, v := range solutionIds {
		if name, ok := solutionsMap[v]; ok {
			solutions = append(solutions, name)
		}
	}
	return solutions
}

// WriteFile define
func WriteFile(fileName string, content []byte) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("write file successful")
}

// LoadFile define
func LoadFile(filePath string) ([]byte, error) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader, output := bufio.NewReader(f), []byte{}
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return output, nil
			}
			return nil, err
		}
		output = append(output, line...)
		output = append(output, []byte("\n")...)
	}
}

// DestoryDir define
func DestoryDir(path string) {
	filepath.Walk(path, func(path string, fi os.FileInfo, err error) error {
		if nil == fi {
			return err
		}
		if !fi.IsDir() {
			return nil
		}
		name := fi.Name()
		if strings.Contains(name, "temp") {
			fmt.Println("temp file name:", path)
			err := os.RemoveAll(path)
			if err != nil {
				fmt.Println("delet dir error:", err)
			}
		}
		return nil
	})
}

// CopyFile define
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

// BinarySearch define
func BinarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}