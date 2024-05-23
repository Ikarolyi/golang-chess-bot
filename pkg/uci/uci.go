package uci

import (
	"bufio"
	"log"
	"log/syslog"
	"os"
	"strings"

	"github.com/ikaroly/gobot/pkg/game"
	"github.com/ikaroly/gobot/pkg/search"
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
        // The position must be refreshed all the time <- "UCI is stateless"
        var _, fen, found = strings.Cut(event, " ")
        if found {
          var fen, after, found = strings.Cut(fen, " moves ")
          if found{
            var moves = strings.Split(after, " ")
            e.position = game.NewPosition(fen)

            e.position.MoveAll(moves)

            log.Println(len(e.position.Pieces))
          }
        }
      case "go":
        println(len(search.GetBoardMoves(e.position)))
        log.Println(e.position.ToString())
      case "debug":
        println(e.position.ToString())
        println("Total ", len(e.position.Pieces))
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