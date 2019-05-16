package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"sync"
)

type Test struct {
	Name string
	Age  int32
}

func Write(ob interface{}) error {
	if reflect.ValueOf(ob).Kind() == reflect.Struct {
		filename := reflect.TypeOf(ob).Name() + ".json"
		data, err := json.Marshal(ob)
		if err != nil {
			return err
		}
		fmt.Println("Saving...: ", ob)
		var m sync.Mutex
		go func(filename string, data []byte, m *sync.Mutex) {
			//m.Lock()
			f, err := os.Create("./" + filename)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()
			_, err = f.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Saved")
			//m.Unlock()
		}(filename, data, &m)
		return nil
	}
	return errors.New("Unsupported type in struct")
}

func Delete(ob interface{}) error {
	if reflect.ValueOf(ob).Kind() == reflect.Struct {
		filename := reflect.TypeOf(ob).Name() + ".json"
		err := os.Remove("./" + filename)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("Unsupported type in struct")
}

func Load(ob interface{}) error {
	if reflect.ValueOf(ob).Kind() == reflect.Ptr {
		typeName := reflect.ValueOf(ob).Elem().Type().Name()
		filename := typeName + ".json"
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}
		err = json.Unmarshal(data, ob)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("Unsupported type in struct")
}
