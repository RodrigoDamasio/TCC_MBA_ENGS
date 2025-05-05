package handler

import (
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Permitir conexões de qualquer origem (modificar conforme necessário para segurança)
		return true
	},
}

// LatestBlockHandler godoc
// @Summary Estabelece uma conexão WebSocket para receber o último bloco em tempo real
// @Description Conecta ao nó Ethereum via WebSocket e notifica o cliente sempre que um novo bloco for minerado
// @Tags Blockchain
// @Produce json
// @Success 101 {object} map[string]interface{} "Conexão WebSocket estabelecida"
// @Failure 500 {object} map[string]string "Erro ao conectar ao nó Ethereum ou WebSocket"
// @Router /ws/latest-block [get]
func LatestBlockHandler(c *gin.Context) {
	// Faz o upgrade para uma conexão WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Erro ao fazer upgrade para WebSocket: %v", err)
		return
	}
	defer conn.Close()

	// Conecta ao nó Ethereum via WebSocket
	client, err := ethclient.Dial("ws://localhost:8561")
	if err != nil {
		log.Printf("Erro ao conectar ao nó Ethereum via WebSocket: %v", err)
		conn.WriteMessage(websocket.TextMessage, []byte("Erro ao conectar ao nó Ethereum"))
		return
	}

	// Canal para receber novos blocos
	headers := make(chan *types.Header)

	// Subscrição para novos blocos
	sub, err := client.SubscribeNewHead(c, headers)
	if err != nil {
		log.Printf("Erro ao subscrever novos blocos: %v", err)
		conn.WriteMessage(websocket.TextMessage, []byte("Erro ao subscrever novos blocos"))
		return
	}
	defer sub.Unsubscribe()

	// Loop para enviar blocos em tempo real para o cliente WebSocket
	for {
		select {
		case err := <-sub.Err():
			log.Printf("Erro na subscrição: %v", err)
			conn.WriteMessage(websocket.TextMessage, []byte("Erro na subscrição ao Ethereum"))
			return
		case header := <-headers:
			// Escreve o número do bloco no WebSocket
			message := map[string]interface{}{
				"blockNumber": header.Number.String(),
				"blockHash":   header.Hash().Hex(),
			}
			if err := conn.WriteJSON(message); err != nil {
				log.Printf("Erro ao enviar mensagem para o WebSocket: %v", err)
				return
			}
		}
	}
}
