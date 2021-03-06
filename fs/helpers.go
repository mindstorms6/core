package fs

import (
	"io/ioutil"
	"os"
)

const publicFileMode os.FileMode = 0644
const publicDirMode os.FileMode = 0755
const privateFileMode os.FileMode = 0600
const privateDirMode os.FileMode = 0700

func Exists(path string) (bool, error) {
	if f, err := os.Open(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	} else {
		f.Close()
		return true, nil
	}
}

func ReadFile(path string) (string, error) {
	if content, err := ioutil.ReadFile(path); err != nil {
		return "", err
	} else {
		return string(content), nil
	}
}
