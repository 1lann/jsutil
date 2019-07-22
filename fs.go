//+build js

package jsutil

import (
	"github.com/gopherjs/gopherjs/js"
)

var fs *js.Object

func init() {
	fs = js.Global.Call("require", "fs")
}

func ReadFile(path string) ([]byte, error) {
	type resultPair struct {
		data []byte
		err  error
	}

	resultChan := make(chan *resultPair)

	fs.Call("readFile", path, func(err *js.Error, data string) {
		go func() {
			if err.Object == nil {
				resultChan <- &resultPair{
					data: []byte(data),
					err:  nil,
				}
			} else {
				resultChan <- &resultPair{
					data: nil,
					err:  err,
				}
			}
		}()
	})

	result := <-resultChan

	return result.data, result.err
}
