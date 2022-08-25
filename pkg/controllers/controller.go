package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/karthik-Prathipati/josh-status-checker/pkg/model"
)

func PublishWebsites(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("error while reading data from requst body at Publish Websites: %v", err)
	}
	var NewWebsitesAdded []model.Website
	listOfWebsites := model.Websites
	// fmt.Println(body)
	json.Unmarshal(body, &NewWebsitesAdded)
	listOfWebsites = append(listOfWebsites, NewWebsitesAdded...)
	// fmt.Println(len(NewWebsitesAdded))
	model.Websites = listOfWebsites
	err2 := json.NewEncoder(w).Encode(model.Websites)

	if err2 != nil {
		fmt.Printf("error while encoding data at Publish websites:  %v\n", err)
	}

}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	listOfWebsites := model.Websites

	allWebsitesStatus := true

	for _, website := range listOfWebsites {
		if !website.Status {
			allWebsitesStatus = false
			break
		}
	}
	json.NewEncoder(w).Encode(allWebsitesStatus)
}

// type StatusChecking interface {
// 	Check(name string, list []model.Website, index int, wg *sync.WaitGroup)
// }

// type HttpChecker struct {
// }

// func (h HttpChecker) Check(url string, listOfWebsites []model.Website, index int, wg *sync.WaitGroup) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	status_code := resp.StatusCode

// 	if status_code != 200 {
// 		listOfWebsites[index].Status = false

// 		return
// 	}

// 	listOfWebsites[index].Status = true
// 	// model.Websites=listOfWebsites
// 	defer wg.Done()
// }

// func init() {
// 	listOfWebsites := model.Websites

// 	var length int = len(listOfWebsites)
// 	var wg sync.WaitGroup
// 	fmt.Println("hi currently in a new func")

// 	for {
// 		var currentStatus StatusChecking = HttpChecker{}
// 		wg.Add(length)
// 		for index, website := range listOfWebsites {
// 			go currentStatus.Check(website.Address, listOfWebsites, index, &wg)
// 		}

// 		model.Websites = listOfWebsites
// 		fmt.Println("changes added now")
// 		time.Sleep(time.Second * 60)
// 	}

// }

func init() {
	listOfWebsites := model.Websites

	// var length int = len(listOfWebsites)
	wg := new(sync.WaitGroup)
	fmt.Println("hi currently in a new func")

	t := time.NewTicker(time.Minute)
	done := make(chan bool)

	// go func() {
	// 	var currentStatus StatusChecking = HttpChecker{}
	// 	wg.Add(len(listOfWebsites))
	// 	for index, website := range listOfWebsites {
	// 		go currentStatus.Check(website.Address, listOfWebsites, index, &wg)
	// 	}
	// 	defer func() {
	// 		done <- true
	// 	}()
	// }()
	go CheckStatusForAllWebsites(listOfWebsites, wg, done)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println(len(listOfWebsites))
				fmt.Println(len(model.Websites))
				model.Websites = listOfWebsites

			case <-t.C:
				listOfWebsites = model.Websites
				fmt.Println("currently in select statement")
				// wg.Add(1)
				go CheckStatusForAllWebsites(listOfWebsites, wg, done)
				// wg.Wait()
			}

			model.Websites = listOfWebsites
			fmt.Println("changes added now")
			// time.Sleep(time.Second * 60)
		}
	}()

}
func CheckStatusForAllWebsites(listOfWebsites []model.Website, wg *sync.WaitGroup, done chan bool) {
	fmt.Printf("currently in CheckForWebsites %v\n", len(listOfWebsites))
	var currentStatus StatusChecking = HttpChecker{}

	wg.Add(len(listOfWebsites))
	for index, website := range listOfWebsites {
		fmt.Println(index)
		go currentStatus.Check(website.Address, listOfWebsites, index, wg)
	}
	wg.Wait()
	done <- true
}

type StatusChecking interface {
	Check(name string, list []model.Website, index int, wg *sync.WaitGroup)
}

type HttpChecker struct {
}

func (h HttpChecker) Check(url string, listOfWebsites []model.Website, index int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("currently in check func")
	resp, err := http.Get(url)
	//fmt.Printf("resp for %v\n", err)
	if err != nil {
		listOfWebsites[index].Status = false
		fmt.Printf("false for %v\n", index)
		return
		// fmt.Println(err)
	}
	status_code := resp.StatusCode

	if status_code >= 500 {
		listOfWebsites[index].Status = false
		fmt.Printf("false for %v\n", index)
		return
	}

	listOfWebsites[index].Status = true
	model.Websites = listOfWebsites
}
