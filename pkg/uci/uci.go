package uci

import (
	"bufio"
	"log"
	"log/syslog"
	"os"
	"strings"

	"github.com/ikaroly/gobot/pkg/game"
)

// UCI standard: https://www.wbec-ridderkerk.nl/html/UCIProtocol.html

const EngineName = "GOBOT"
const EngineAuthor = "ikarolyi"

type Engine struct {
  log_file *syslog.Writer
  position game.Board
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
  // var event string

  reader := bufio.NewReader(os.Stdin)
  event, _ := reader.ReadString('\n')
  // fmt.Fscanln(reader, &event)  

  return event
}

func (e Engine) Listen() {
  mainloop:
  for {
    var event = e.Read()
    switch strings.TrimSpace(strings.Split(event, " ")[0]){
      case "Quit":
        log.Println("Quit")
        break mainloop
      case "uci":
        identify() 
        println("uciok")
      case "isready":
        // TODO confiurations
        println("readyok")
      case "position":
        // 2024-05-03 17:49:45,072-->1:position startpos moves e2e4
        var _, fen, found = strings.Cut(event, " ")
        if found {
          e.position = game.NewPosition(fen)
        }
      default: 
        log.Println("Unknown command: ", event)
    }
    // if(event == "Quit") {
    //   log.Println("Quit")
    //   break
    // }else if(event != ""){
    //   log.Println(event)
    // }
  }
}

func identify() {
  println("id name " + EngineName)
  println("id author " + EngineAuthor)
}