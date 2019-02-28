package main

import (
    "fmt"
    "sync"
    "time"
)

type garfo struct{
	sync.Mutex
}

type filosofo struct{
	id int
	garfo_esq, garfo_dir *garfo
}

func(f *filosofo) comer(){
	f.garfo_esq.Lock()
	fmt.Printf("Filosofo #%d esta pegou o garfo esquerdo\n", f.id)
	time.Sleep(time.Second)
	f.garfo_dir.Lock()
	fmt.Printf("Filosofo #%d esta pegou o garfo direito\n", f.id)
	
	fmt.Printf("Filosofo #%d esta comendo...\n", f.id)
	
	time.Sleep(time.Second*2)
	
	fmt.Printf("Filosofo #%d terminou de comer!\n", f.id)
	
	f.garfo_esq.Unlock()
	f.garfo_dir.Unlock()
}

func main(){
	garfos := make([]*garfo, 5)
	filosofos := make([]*filosofo, 5)

	for i:= 0; i < 5; i++{
		garfos[i] = &garfo{}
	}

	for i:= 0; i < 5; i++{
		filosofos[i] = &filosofo{
			id: i,
			garfo_esq: garfos[i],
			garfo_dir: garfos[(i+1)%5],
		}
		go filosofos[i].comer()
	}

	jantar_filos := make(chan int)
	<- jantar_filos
}