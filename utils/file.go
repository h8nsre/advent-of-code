package utils

import (
	"bufio"
	"github.com/thanos-io/thanos/pkg/runutil"
	"os"
)

func ScanLine(filename string, f func(string) error) error {
	file, err := os.Open(filename)

	if err != nil {
		return err
	}

	defer runutil.CloseWithErrCapture(&err, file, "close file")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if err := f(scanner.Text()); err != nil {
			return err
		}
	}
	return scanner.Err()
}
