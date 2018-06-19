package main

import (
	"log"
	"github.com/urfave/cli"
	"os"
	"sync"
	"github.com/shin1x1/en9sqs/enqueue"
	"github.com/shin1x1/en9sqs/worker"
)

type config struct {
	queueUrl        string
	region          string
	concurrentCount int
	messageCount    int
	message         string
}

func main() {
	app := cli.NewApp()

	app.Name = "en9sqs"
	app.Usage = "concurrent enqueue messages to Amazon SQS"
	app.ArgsUsage = "queue_url message"
	app.Version = "0.0.2"
	setUp(app)

	app.Action = func(context *cli.Context) error {
		if len(context.Args()) < 2 {
			cli.ShowAppHelpAndExit(context, 255)
		}

		conf := &config{
			queueUrl:        context.Args().Get(0),
			message:         context.Args().Get(1),
			region:          context.String("region"),
			concurrentCount: context.Int("concurrency"),
			messageCount:    context.Int("messages"),
		}

		loopChan := make(chan int)
		var wg sync.WaitGroup

		threshold := conf.messageCount / 10;
		for i := 1; i <= conf.concurrentCount; i++ {
			e := enqueue.NewSqsEnqueue(conf.queueUrl, conf.region)
			w := worker.NewWorker(&wg, i, e, threshold)

			wg.Add(1)
			go w.Run(conf.message, loopChan)
		}

		go func() {
			for i := 1; i <= conf.messageCount; i++ {
				loopChan <- i
			}
			close(loopChan)
		}()

		wg.Wait()
		log.Printf("Done(%d messages)", conf.messageCount)

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func setUp(app *cli.App) {
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "region, r",
			Usage: "Region",
			Value: "ap-northeast-1",
		},
		cli.StringFlag{
			Name:  "concurrency, c",
			Usage: "Number of concurrency workers",
			Value: "1",
		},
		cli.StringFlag{
			Name:  "messages, n",
			Usage: "Number of enqueueing messages",
			Value: "1",
		},
	}
}
