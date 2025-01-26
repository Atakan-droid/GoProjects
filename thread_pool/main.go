package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"thread_pool/work"
)

func getFromAPI() error {
	const urlString = "https://jsonplaceholder.typicode.com/posts/1"
	resp, err := http.Get(urlString)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("API response received", resp.Status)
	return nil
}

func main() {
	wp, err := work.NewPool(5, 5)
	if err != nil {
		log.Fatalf("error creating pool: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wp.Start(ctx)

	for range 20 {

		task := work.NewTask(getFromAPI, func(err error) {
			fmt.Println("error occurred", err)
		})

		wp.AddTaskNonBlocking(task)
	}

	for complated := range wp.TaskCompleted() {
		fmt.Println("task completed", complated)
		if wp.IsPoolDone() {
			wp.Stop()
			break
		}
	}
}
