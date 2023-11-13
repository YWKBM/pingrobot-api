package pingrobot

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"
)

type WebServiceInfo struct {
	ID        int    `json:"id"`
	UserEmail string `json:"user_email"`
	Name      string `json:"name"`
	Link      string `json:"link"`
	Port      int    `json:"port"`
	Status    string `json:"status"`
}

type Result struct {
	ID           int
	UserEmail    string
	URL          string
	StatusCode   int
	ResponseTime time.Duration
	Error        error
}

type Pool struct {
	db           *sql.DB
	workersCount int
	tasks        chan WebServiceInfo
	results      chan Result
	workers      []*Worker
	webServices  []WebServiceInfo
	wg           sync.WaitGroup
}

func NewPool(db *sql.DB, workersCount int, tasks chan WebServiceInfo, results chan Result) *Pool {
	return &Pool{
		db:           db,
		workersCount: workersCount,
		tasks:        tasks,
		results:      results,
		webServices:  make([]WebServiceInfo, 128),
		wg:           sync.WaitGroup{},
	}
}

func (p *Pool) RunBackground() {
	for i := 1; i <= p.workersCount; i++ {
		worker := newWorker(i, p.tasks, time.Second)
		p.workers = append(p.workers, worker)
		go worker.StartBackground(&p.wg, p.results)
	}
}

func (p *Pool) getAllWebServiceInfo() {
	rows, err := p.db.Query("SELECT * FROM web_services")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var webService WebServiceInfo

		err := rows.Scan(&webService.ID, &webService.UserEmail, &webService.Name, &webService.Link, &webService.Port, &webService.Status)
		p.webServices = append(p.webServices, webService)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (p *Pool) generateTasks() {
	time.Sleep(time.Minute)
	p.getAllWebServiceInfo()
	for _, webService := range p.webServices {
		p.tasks <- webService
	}
}

func (p *Pool) processResults() {
	go func() {
		for result := range p.results {
			fmt.Println(result)
			if result.Error != nil {
				//TODO: send email
				p.db.Query("UPDATE web_services SET status = $1 WHERE ID = $2", "ERROR", result.ID)
			}

			p.db.Query("UPDATE web_services SET status = $1 WHERE ID = $2", "SUCCESS", result.ID)
		}
	}()
}

func (p *Pool) Stop() {
	for _, worker := range p.workers {
		worker.Stop()
	}
}
