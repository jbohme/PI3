package service

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/jbohme/crud/configs/rest_err"
	"sync"
)

type PlayerDomain struct {
	Id       string          `json:"id"`
	NickName string          `json:"nick_name"`
	QtyWins  uint            `json:"qty_wins"`
	Conn     *websocket.Conn // Conexão WebSocket
}

type Status uint

const (
	Available Status = iota
	Unavailable
	InGame
)

type GameMessage struct {
	Action   string    `json:"action"`   // "move" ou "start"
	Player   string    `json:"player"`   // "X" ou "O"
	Position int       `json:"position"` // Posição no tabuleiro (0-8)
	Board    [9]string `json:"board"`
	Winner   string    `json:"winner,omitempty"`
}

type RoomDomain struct {
	Id      string          `json:"id"`
	Players [2]PlayerDomain `json:"players"`
	Status  Status          `json:"status"`
	Board   [9]string       `json:"board"`
	Turn    string          `json:"turn"` // "X" ou "O"
	Mutex   sync.Mutex      // Mutex para controle de concorrência
}

var rooms = make(map[string]*RoomDomain)
var mu sync.Mutex

func (rd *RoomDomain) GetID() string {
	return rd.Id
}

func (rd *RoomDomain) GetFirstPlayer() PlayerDomain {
	return rd.Players[0]
}

func (rd *RoomDomain) GetSecondPlayer() PlayerDomain {
	return rd.Players[1]
}

func (rd *RoomDomain) GetStatus() Status {
	return rd.Status
}

func (rd *RoomDomain) SetFirstPlayer(player PlayerDomain) {
	rd.Players[0] = player
}

func (rd *RoomDomain) SetSecondPlayer(player PlayerDomain) {
	rd.Players[1] = player
}

func JoinRandomRoomServices(
	playerDomain PlayerDomain,
) (
	RoomDomain,
	*rest_err.RestErr) {

	room, err := JoinRandomRoom(playerDomain)
	if err != nil {
		return RoomDomain{}, err
	}

	if room.Players[1].Id != "" { // Sala cheia, iniciar jogo
		go room.StartGame()
	}

	return room, nil
}

func JoinRandomRoom(player PlayerDomain) (RoomDomain, *rest_err.RestErr) {
	mu.Lock()
	defer mu.Unlock()

	for _, room := range rooms {
		if room.GetStatus() == Available {
			if room.Players[0].Id == "" {
				room.SetFirstPlayer(player)
			} else if room.Players[1].Id == "" {
				room.SetSecondPlayer(player)
				room.Status = Unavailable
			}
			return *room, nil
		}
	}

	newRoom := &RoomDomain{
		Id:      generateRoomID(),
		Players: [2]PlayerDomain{player},
		Status:  Available,
		Turn:    "X",
	}
	rooms[newRoom.Id] = newRoom
	return *newRoom, nil
}

func generateRoomID() string {
	return fmt.Sprintf("room-%d", len(rooms)+1)
}

func (rd *RoomDomain) StartGame() {
	rd.Status = InGame
	rd.Board = [9]string{}

	for {
		currentPlayer := rd.GetCurrentPlayer()
		var msg GameMessage

		// Lê a jogada do jogador
		err := currentPlayer.Conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Erro na leitura da mensagem:", err)
			return
		}

		// Processa a jogada
		if msg.Action == "move" && rd.Board[msg.Position] == "" {
			rd.Board[msg.Position] = rd.Turn
			if rd.checkWin(rd.Turn) {
				rd.broadcast(GameMessage{
					Action: "end",
					Winner: rd.Turn,
					Board:  rd.Board,
				})
				return
			}

			// Alterna o turno
			rd.Turn = rd.switchTurn()
			rd.broadcast(GameMessage{
				Action:   "update",
				Player:   rd.Turn,
				Board:    rd.Board,
				Position: msg.Position,
			})
		}
	}
}

func (rd *RoomDomain) GetCurrentPlayer() *PlayerDomain {
	if rd.Turn == "X" {
		return &rd.Players[0]
	}
	return &rd.Players[1]
}

func (rd *RoomDomain) switchTurn() string {
	if rd.Turn == "X" {
		return "O"
	}
	return "X"
}

func (rd *RoomDomain) broadcast(msg GameMessage) {
	for _, player := range rd.Players {
		err := player.Conn.WriteJSON(msg)
		if err != nil {
			fmt.Println("Erro ao enviar mensagem:", err)
		}
	}
}

func (rd *RoomDomain) checkWin(player string) bool {
	winComb := [][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // Linhas
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // Colunas
		{0, 4, 8}, {2, 4, 6}, // Diagonais
	}

	for _, comb := range winComb {
		if rd.Board[comb[0]] == player && rd.Board[comb[1]] == player && rd.Board[comb[2]] == player {
			return true
		}
	}
	return false
}
