package main

import(
	 "github.com/astaxie/beego/logs"
	 "sync"
)

func main(){
	log := logs.NewLogger(10000)
	log.SetLogger("file", `{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10, "perm": 0600}`)
	// log.Debug("this is a debug message")
	// log.Info("this is a info message")
	// log.Debug("this is a debug message")
	// log.Info("this is a info message")
	// log.Debug("this is a debug message")
	var wg sync.WaitGroup
	wg.Add(10000)
	for i := 0; i < 10000; i ++{
		go func(w *sync.WaitGroup){
			log.Info("this is a info message")
			w.Done()
		}(&wg)
	}
	wg.Wait()
	// log.Info("this is a info message")
	defer log.Close()
}
