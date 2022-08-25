package statuschecker

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"sync"
// 	"time"

// 	"github.com/karthik-Prathipati/josh-status-checker/pkg/model"
// )

// func init() {
// 	listOfWebsites := model.Websites

// 	// var length int = len(listOfWebsites)
// 	var wg sync.WaitGroup
// 	fmt.Println("hi currently in a new func")

// 	t := time.NewTicker(time.Minute)
// 	done := make(chan bool)

// 	// go func() {
// 	// 	var currentStatus StatusChecking = HttpChecker{}
// 	// 	wg.Add(len(listOfWebsites))
// 	// 	for index, website := range listOfWebsites {
// 	// 		go currentStatus.Check(website.Address, listOfWebsites, index, &wg)
// 	// 	}
// 	// 	defer func() {
// 	// 		done <- true
// 	// 	}()
// 	// }()

// 	go func() {
// 		for {
// 			select {
// 			case <-done:
// 				model.Websites = listOfWebsites

// 			case <-t.C:
// 				go CheckStatusForAllWebsites(listOfWebsites, &wg, done)
// 			}

// 			// model.Websites = listOfWebsites
// 			fmt.Println("changes added now")
// 			// time.Sleep(time.Second * 60)
// 		}
// 	}()

// }
// func CheckStatusForAllWebsites(listOfWebsites []model.Website, wg *sync.WaitGroup, done chan bool) {
// 	var currentStatus StatusChecking = HttpChecker{}
// 	wg.Add(len(listOfWebsites))
// 	for index, website := range listOfWebsites {
// 		go currentStatus.Check(website.Address, listOfWebsites, index, wg)
// 	}
// 	defer func() {
// 		done <- true
// 	}()
// }

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
