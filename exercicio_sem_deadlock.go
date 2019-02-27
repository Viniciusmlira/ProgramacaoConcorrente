//executar no play.golang.org

package main
import (
	"fmt"
	"sync"
	"time"
	"os"
)

var qte_fil = 5 //quantidade de filósofos
var count = 0

type garfo struct{ 
	sync.Mutex 
}

type filosofo struct{
	id int
	garfo_esq, garfo_dir *garfo
}

func (f filosofo) comer(){	
	f.garfo_esq.Lock()
	f.garfo_dir.Lock()

	fmt.Printf("Filosofo #%d esta comendo \n", f.id)
	time.Sleep(time.Second)

	f.garfo_dir.Unlock()
	f.garfo_esq.Unlock()

	fmt.Printf("Filosofo #%d terminou de comer \n", f.id)
	time.Sleep(time.Second)

	if count<qte_fil-1 {
		count++
	} else {
		os.Exit(0)
	}
}


func main() {
	
	garfos := make([]*garfo, qte_fil)
	for i := 0; i < qte_fil; i++{
		garfos[i] = new(garfo)
	}

	filosofos := make([]*filosofo, qte_fil)
	for i := 0; i < qte_fil; i++{
		filosofos[i] = &filosofo{
			id: i, garfo_esq: garfos[i], garfo_dir: garfos[(i+1)%qte_fil]}
		go filosofos[i].comer()
	}
	
	filosofos_comem := make(chan int, 1)
	<- filosofos_comem 
	// abre o canal
	
}


