package uci

import "fmt"

import "log"
import "log/syslog"

type Engine struct {
  log_file *syslog.Writer
}

type Server interface {
  Init()
	Read() string
  Listen()
}

func (e Engine) Init() {
	e.log_file, _ = syslog.New(syslog.LOG_SYSLOG, "GOBOT")

	log.SetOutput(e.log_file)

	log.Println("Engine server started")
  
}

func (e Engine) Read() string {
	var event string
  fmt.Scanf("%s", &event)

  return event
}

func (e Engine) Listen() {
  for {
    var event = e.Read()
    if(event == "Quit") {
      log.Println("Quit")
      break
    }else if(event != ""){
      log.Println(event)
    }
  }
}