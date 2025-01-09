package session

import "sync"

type data struct {
	Data map[string]any
	mu   sync.RWMutex
}

func (d *data) Get(key string) any {
	return d.Data[key]
}

func (d *data) Set(key string, value any) {
	d.Data[key] = value
}

func (d *data) Delete(key string) {
	delete(d.Data, key)
}

func (d *data) Reset() {
	d.Data = make(map[string]any)
}
