package worker

import (
	"sync"
	"github.com/shin1x1/en9sqs/enqueue"
	"log"
)

type Worker struct {
	wg           *sync.WaitGroup
	no           int
	enq          *enqueue.SqsEnqueue
	logThreshold int
}

func NewWorker(wg *sync.WaitGroup, no int, enq *enqueue.SqsEnqueue, log int) *Worker {
	return &Worker{
		wg:           wg,
		no:           no,
		enq:          enq,
		logThreshold: log,
	}
}

func (w *Worker) Run(mess string, ch chan int) {
	defer w.wg.Done()

	for i := range ch {
		if err := w.enq.Enqueue(mess); err != nil {
			log.Fatalln(err)
		}

		if w.logThreshold > 0 && i%w.logThreshold == 0 {
			log.Printf("message %d done(worker:%d)", i, w.no)
		}
	}
}
