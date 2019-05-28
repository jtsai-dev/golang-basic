package filehandling

import (
	"io/ioutil"
	"net/http"
	"os"
)

func HandleFile(writer http.ResponseWriter, request *http.Request) error {
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
