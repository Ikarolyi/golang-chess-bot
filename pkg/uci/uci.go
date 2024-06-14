package uci

import (
	"bufio"
	"fmt"
	"log"
	"log/syslog"
	"os"
	"strings"

	"github.com/ikaroly/gobot/pkg/bitboard"
	"github.com/ikaroly/gobot/pkg/evaluate"
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
    event = strings.TrimSpace(event)
    switch strings.TrimSpace(strings.Split(event, " ")[0]){
      case "quit":
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
        // println(len(search.GetBoardMoves(e.position)))
        println(GetBestMove(e.position))
      case "debug":
        println("Position FEN: ", e.position.ExportFEN())
        println("True combined board", e.position.BoardCombined.GetTrueCombined())
        println("Total ", len(e.position.Pieces))
        println("Eval depth0: ", evaluate.Evaluate(e.position))
      default: 
        log.Println("Unknown command: ", event)

    }
  }
}

func identify() {
  println("id name " + EngineName)
  println("id author " + EngineAuthor)
}

func GetBestMove(b game.Board) string{
  best_move, eval_score := search.MasterSearch(b, 4)
  println("info score cp", eval_score * int(b.SideToMove))

  string_move := bitboard.Decode(best_move.From) + bitboard.Decode(best_move.To)
  return fmt.Sprint("bestmove ", string_move)
}