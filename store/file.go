package store

import (
	"io/ioutil"
	"os"
)

func IsExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateFile(fileName string) (err error) {
	ok, err := IsExists(fileName)
	if err != nil {
		return
	}
	if ok {
		return nil
	}
	f, err := os.Create(fileName)
	defer f.Close()
	return
}

func CreateFileTrunc(fileName string) (err error) {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC, 0600)
	//_, err = f.Write([]byte{})
	defer f.Close()
	return
}

func WriteFile(fileName string, data []byte) (err error) {
	f, err := os.OpenFile(fileName, os.O_RDWR, 0600)
	defer f.Close()
	if err == nil {
		_, err = f.Write(data)
	}
	return
}

func ReadFile(fileName string) (data []byte, err error) {
	f, err := os.OpenFile(fileName, os.O_RDONLY, 0600)
	defer f.Close()
	if err == nil {
		data, err = ioutil.ReadAll(f)
		return
	}
	return nil, err
}
