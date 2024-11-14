//package room
//
//import (
//	"errors"
//	"fmt"
//	"sync"
//)
//
//type Player struct {
//	ID   string
//	Name string
//}
//
//type Room struct {
//	ID      string
//	Players []Player
//	Status  string
//}
//
//var rooms = make(map[string]*Room)
//var mu sync.Mutex
//
//// Função para criar uma nova sala
//func createRoom(id string) (*Room, error) {
//	mu.Lock()
//	defer mu.Unlock()
//
//	if _, exists := rooms[id]; exists {
//		return nil, errors.New("sala já existe")
//	}
//
//	room := &Room{
//		ID:      id,
//		Players: []Player{},
//		Status:  "Aguardando",
//	}
//
//	rooms[id] = room
//	return room, nil
//}
//
//// Função para buscar uma sala pelo ID
//func findRoomById(id string) (*Room, error) {
//	mu.Lock()
//	defer mu.Unlock()
//
//	if room, exists := rooms[id]; exists {
//		return room, nil
//	}
//	return nil, errors.New("sala não encontrada")
//}
//
//// Função para buscar uma sala aleatória disponível
//func joinRandomRoom(player Player) (*Room, error) {
//	mu.Lock()
//	defer mu.Unlock()
//
//	for _, room := range rooms {
//		if room.Status == "Aguardando" && len(room.Players) < 2 {
//			room.Players = append(room.Players, player)
//			if len(room.Players) == 2 {
//				room.Status = "Cheia"
//			}
//			return room, nil
//		}
//	}
//	return nil, errors.New("nenhuma sala disponível")
//}
//
//// Exemplo de uso
//func main() {
//	player1 := Player{ID: "1", Name: "Jogador 1"}
//	player2 := Player{ID: "2", Name: "Jogador 2"}
//
//	room, _ := createRoom("1234")
//	fmt.Println("Sala criada:", room.ID)
//
//	room, _ = joinRandomRoom(player1)
//	fmt.Println("Player 1 entrou na sala aleatória:", room.ID)
//
//	room, _ = joinRandomRoom(player2)
//	fmt.Println("Player 2 entrou na sala aleatória:", room.ID)
//}
