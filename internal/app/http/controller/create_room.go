package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jbohme/crud/configs/logger"
	"github.com/jbohme/crud/internal/app/model"
	"github.com/jbohme/crud/internal/app/model/service"
	"net/http"
)

type GameMessage struct {
	Action   string    `json:"action"`   // ex: "move", "start"
	Player   string    `json:"player"`   // ex: "X" ou "O"
	Position int       `json:"position"` // Posição no tabuleiro (1-9)
	Board    [9]string `json:"board"`
	Winner   string    `json:"winner,omitempty"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (uc *userControllerInterface) JoinRandomRoom(c *gin.Context) {
	// Tentando fazer o upgrade para WebSocket
	conn, errConn := upgrader.Upgrade(c.Writer, c.Request, nil)
	if errConn != nil {
		// Se a conexão WebSocket não for possível, respondemos com HTTP antes do upgrade
		//logger.Error(fmt.Sprintf("Error upgrading to WebSocket: %v", errConn))
		// Finaliza a requisição e não tenta mais escrever no HTTP
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade WebSocket connection"})
		return
	}

	// Verificar o token de autenticação do usuário
	user, errToken := model.VerifyToken(c.Request.Header.Get("Authorization"))
	if errToken != nil {
		// Se a autenticação falhar, retorna via HTTP antes do upgrade
		c.JSON(errToken.Code, errToken)
		return
	}
	logger.Info(fmt.Sprintf("User authenticated: %#v", user))

	// Criar o objeto do jogador com a conexão WebSocket
	playerDomain := service.PlayerDomain{
		Id:       user.GetID(),
		NickName: user.GetNickName(),
		QtyWins:  user.GetWins(),
		Conn:     conn, // A conexão WebSocket
	}

	// Chamar o serviço para tentar juntar o jogador a uma sala
	domainResult, errDomain := service.JoinRandomRoomServices(playerDomain)
	if errDomain != nil {
		// Se falhar ao juntar o jogador à sala, responde via HTTP (não mais via WebSocket)
		c.JSON(errDomain.Code, errDomain)
		return
	}

	// Enviar uma mensagem via WebSocket indicando que o jogador entrou na sala
	err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Você entrou na sala: %s", domainResult.GetID())))
	if err != nil {
		// Se houver erro ao enviar a mensagem pelo WebSocket
		//logger.Error(fmt.Sprintf("Error sending WebSocket message: %v", err))
		return
	}

	// Após o upgrade, não devemos mais responder com HTTP.
	// A comunicação agora será via WebSocket.
}
