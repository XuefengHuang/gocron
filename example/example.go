package main

import (
	"encoding/json"
	"fmt"
	"github.com/XuefengHuang/gocron"
)

func task() {
	fmt.Println("I am runnning task.")
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func main() {
	// Do jobs with params
	gocron.Every(1).Second().Do(taskWithParams, 1, "hello")

	// Do jobs on specific weekday
	gocron.Every(1).Monday().Do(task)
	gocron.Every(1).Thursday().Do(task)

	// function At() take a string like 'hour:min'
	gocron.Every(1).Day().At("10:30").Do(task)
	gocron.Every(1).Monday().At("18:30").Do(task)

	// remove, clear and next_run
	_, time := gocron.NextRun()
	fmt.Println(time)

	jobs := gocron.GetJobs()
	fmt.Println(jobs[2].Id)
	b, err := json.Marshal(jobs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
	gocron.Start()
	gocron.Clear()

	//gocron.RemoveById(1)
	//gocron.RemoveById(2)
	//gocron.RemoveById(3)
	//gocron.RemoveById(4)

	jobs = gocron.GetJobs()
	b, err = json.Marshal(jobs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
	// also , you can create a your new scheduler,
	// to run two scheduler concurrently
	/*
		s := gocron.NewScheduler()
		s.Every(3).Seconds().Do(task)
		<-s.Start()
	*/
	gocron.Stop()
	stopped := make(chan bool, 1)
	// function Start start all the pending jobs
	<-stopped
}
