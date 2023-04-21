package gotools

import (
    "sync"
    "time"
)

type item struct {
    value      interface{}
    expiration int64
}

type MemStore struct {
    store map[string]item
    mutex sync.Mutex
}

// NewMemStore cria uma nova instância do MemStore
func NewMemStore() *MemStore {
    return &MemStore{store: make(map[string]item)}
}

// Set armazena um valor com a chave especificada e expiração após um período de tempo
func (s *MemStore) Set(key string, value interface{}, duration time.Duration) {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    expTime := time.Now().Add(duration).UnixNano()
    s.store[key] = item{value: value, expiration: expTime}
}

// Get recupera o valor armazenado pela chave especificada. Se o valor expirou ou não existe, retorna nil
func (s *MemStore) Get(key string) interface{} {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    if val, ok := s.store[key]; ok {
        if time.Now().UnixNano() < val.expiration {
            return val.value
        }
        delete(s.store, key)
    }
    return nil
}
