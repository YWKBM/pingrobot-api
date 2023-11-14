package pingrobot

import (
	"database/sql"
	"os"
	"os/signal"
	"syscall"
)

func Run(db *sql.DB) {
	go func() {
		results := make(chan Result)
		tasks := make(chan *WebServiceInfo)

		pool := NewPool(db, 5, tasks, results)

		go pool.RunBackground()
		go pool.generateTasks()
		go pool.processResults()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

		<-quit
	}()
}
