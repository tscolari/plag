package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"code.cloudfoundry.org/lager"
)

var wg sync.WaitGroup

func main() {
	logger := lager.NewLogger("test-app")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))

	for i := 0; i < 20; i++ {
		group := i % 5
		wg.Add(1)
		go logAndWait(logger, fmt.Sprintf("function-%d", group))
	}

	wg.Wait()
}

func logAndWait(logger lager.Logger, sessionName string) {
	randSleep := rand.Int() % 10
	time.Sleep(time.Duration(randSleep) * time.Second)

	logger = logger.Session(sessionName)
	logger.Info("start")
	defer logger.Info("end")
	defer wg.Done()

	logger.Info("start-to-sleep")
	randSleep = rand.Int() % 10
	time.Sleep(time.Duration(randSleep) * time.Second)
	logger.Info("finished-to-sleep")
}
