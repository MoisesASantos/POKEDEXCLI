package internal

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt	time.Time
	val 		[]byte
}

type cache struct {
	map_result map[string]cacheEntry
	my_mutex	*sync.Mutex
}

func (r cache) Add(key_intro string, data []byte) {

	r.my_mutex.Lock()
	defer r.my_mutex.Unlock()
	
	if r.map_result == nil {
		r.map_result = make(map[string]cacheEntry)
	}
	var entry_to_add cacheEntry

	entry_to_add.createdAt = time.Now() 
	entry_to_add.val = data 
  	r.map_result[key_intro] = entry_to_add
}

func (r cache) Get(key_intro string) ([]byte, bool) {
  
	r.my_mutex.Lock()
	defer r.my_mutex.Unlock()

	result, ok := r.map_result[key_intro]
	return result.val, ok
}


func (r cache) remove_item(interval time.Duration) {

	r.my_mutex.Lock()
	defer r.my_mutex.Unlock()

	limite := time.Now().Add(-interval)
	for chave, entrada := range r.map_result {
		if entrada.createdAt.Before(limite) {
			delete(r.map_result, chave)
		}
	}
}

func (r cache) reapLoop(interval time.Duration) {
  
	//cria um ticker a cada intervalo de tempo, o ticker tem um canal que é enviado um sinal por uma go routine interna cada vez que passa o intervalo de tempo difinido, por cada elemento ou sinal no canal do ticker significa um aviso que passou o intervalo de tempo, sendo assim entra em remove_item, bloquea o map_result de ser acessado por outras go routine, as outras go routine ficam em espera, para ler e alterar o mapa, é calculado o limete de tempo pelo qual o cache se tornou antigo, subtraindo o intervalo do nosso tempo atual, depois fizemos um loop em map_result e o que tiver sido criado antes desse limite é iliminado
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		r.remove_item(interval)
	}
}

func NewCache(interval time.Duration) *cache {
	
	newCache := cache{
		map_result: make(map[string]cacheEntry),
		my_mutex:   new(sync.Mutex),
	}
	go newCache.reapLoop(interval)
	
	return &newCache
}


