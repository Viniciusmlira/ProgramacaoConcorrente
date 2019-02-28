package main
import(
	"fmt"
	"time"
	"sync"
)

type garfo struct{
	sync.Mutex
	Locked bool
}

type filosofo struct{
	id int
	garfo_esq, garfo_dir *garfo
}

func(f *filosofo) comer(){
	for{
		f.garfo_esq.Lock()
		if f.garfo_dir.Locked{
			f.garfo_esq.Unlock()
		}else{
			f.garfo_dir.Lock()
			f.garfo_dir.Locked = true
			fmt.Printf("Filosofo #%d esta comendo...\n", f.id)
			time.Sleep(time.Second*2)
			fmt.Printf("Filosofo #%d terminou de comer!\n", f.id)
			f.garfo_esq.Unlock()
			f.garfo_dir.Unlock()
			f.garfo_dir.Locked = false
			f.garfo_esq.Locked = false
			break
		}
	}
}

func main(){
	garfos := make([]*garfo, 5)
	filosofos := make([]*filosofo, 5)
	for i:= 0; i < 5; i++{
		garfos[i] = &garfo{
			Locked: false,
		}
	}
	
	for i := 0; i < 5; i++{
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

