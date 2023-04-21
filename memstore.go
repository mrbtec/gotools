package gotools

import (
    "sync"
    "time"
    "runtime"
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

// ClearMemory libera a memória alocada pelo programa para o garbage collector
func (s *MemStore)  ClearMemory() {
    // Chama a função GC() do pacote runtime para forçar a execução do garbage collector
    runtime.GC()

    // Aloca um slice vazio de bytes para forçar a liberação de qualquer memória não utilizada
    var dummySlice []byte
    for i := 0; i < 10; i++ {
        dummySlice = append(dummySlice, make([]byte, 1000000)...)
    }
    dummySlice = nil
}

// SetMemoryLimit define um limite máximo para a quantidade de memória que o programa pode alocar.
func (s *MemStore) SetMemoryLimit(limit uint64) {
    memLimit := &limit
    if *memLimit > 0 {
        runtime.MemProfileRate = 0
        go func() {
            for {
                var memStats runtime.MemStats
                runtime.ReadMemStats(&memStats)
               	if memStats.Alloc > *memLimit {
                    runtime.GC()
                }
                time.Sleep(time.Second)
            }
        }()
    }
}

// GetMemoryUsage retorna a quantidade de memória atualmente em uso pelo programa em bytes.
func (s *MemStore) GetMemoryUsage() uint64 {
    var memStats runtime.MemStats
    runtime.ReadMemStats(&memStats)
    return memStats.Alloc
}


