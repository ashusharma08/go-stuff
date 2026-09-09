package main

import (
	"bufio"
	"context"
	"os"
	"os/signal"
	"path"
	"sync"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	guard := make(chan struct{})
	go func() {
		defer func() {
			guard <- struct{}{}
		}()
		err := mainErr(ctx)
		if err != nil {
		}
	}()
	select {
	case <-ctx.Done():
	case <-guard:
		return
	}

	<-guard
}

func mainErr(pctx context.Context) error {

	ctx, cancel := context.WithCancel(pctx)
	defer cancel()
	disp := NewDispatcher()
	errorAggChan := make(chan *LogEntry, 100)
	errorAggregator := newErrorAggregator()
	go errorAggregator.process(ctx, errorAggChan)
	disp.register(errorAggChan)
	defer func() {
		errorAggregator.print()
	}()

	labelsAggChan := make(chan *LogEntry, 100)
	keywordAggregator := newKeywordAggregator([]string{"database", "connection", "open"})
	go keywordAggregator.process(ctx, labelsAggChan)
	disp.register(labelsAggChan)
	defer func() {
		keywordAggregator.print()
	}()

	dataChan := make(chan string, 10)

	directoryItems, err := os.ReadDir("./logfiles")
	if err != nil {
		return err
	}
	filesChan := make(chan string, 5)
	var writerWG sync.WaitGroup

	var wg sync.WaitGroup
	for range 10 {
		wg.Add(1)
		writerWG.Add(1)
		go func() {
			defer wg.Done()
			defer writerWG.Done()
			proc := newProcessor(dataChan, disp)
			go proc.process(ctx)

			for {
				select {
				case <-ctx.Done():
					return
				case path, ok := <-filesChan:
					if !ok {
						return
					}

					f, err := os.Open(path)
					if err != nil {
						cancel()
						return
					}
					defer f.Close()
					scanner := bufio.NewScanner(f)
					for scanner.Scan() {
						line := scanner.Text()
						dataChan <- line
					}
				}
			}
		}()
	}
	wg.Add(1)
	go func() {
		defer func() {
			close(filesChan)
			wg.Done()
		}()

		for _, item := range directoryItems {
			if item.IsDir() {
				continue
			}
			select {
			case <-ctx.Done():
				return
			case filesChan <- path.Join("logfiles", item.Name()):
			}
		}
	}()
	go func() {
		writerWG.Wait()
		close(dataChan)
	}()
	wg.Wait()
	return nil
}
