package filehandling

import (
	"io/ioutil"
	"net/http"
	"os"
)

// 定义自定义异常类型
type userError string

func (e userError) Error() string {
	return e.Message()
}
func (e userError) Message() string {
	return string(e)
}

func HandleFile(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
	path := request.URL.Path[1:]
	file, err := os.Open(path)
	if err != nil {
		// panic(err)
		//http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		// panic(err)
		return err
	}

	writer.Write(bytes)
	return nil
}
