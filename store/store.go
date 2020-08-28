package store

import (
	"fmt"
	"sync"
)

//Store is the interface that wraps basic CRUD operation
type Store interface {
	get(key string) ([]byte, bool)
	post(key string, value []byte)
	del(key string)
}

//Data stores the data
type Data struct {
	sync.Mutex
	Data map[string][]byte
}

func (s *Data) post(key string, value []byte) {
	fmt.Println("post called")
	s.Lock()

	s.Data[key] = append(s.Data[key], value...)

	s.Unlock()

}

func (s *Data) get(key string) ([]byte, bool) {

	fmt.Println("get called")
	var data []byte

	if key == "" {
		var tmp string
		for key := range s.Data {
			tmp += key + ","
		}
		return []byte(tmp), true
	}

	
	if data, ok := s.Data[key]; ok {
		return data, ok
	}
	return data, false
}

func (s *Data) del(key string) {
	fmt.Println("del called")
	s.Lock()

	delete(s.Data, key)

	s.Unlock()
}

//Get Wraps get method to Store interface
func Get(s Store, key string) ([]byte, bool) {
	return s.get(key)
}

//Post Wraps get method to Store interface
func Post(s Store, key string, value []byte) {
	s.post(key, value)
	return
}

//Delete Wraps get method to Store interface
func Delete(s Store, key string) {
	s.del(key)
	return
}
