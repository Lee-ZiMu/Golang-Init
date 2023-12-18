/**
 * @Description: websocket
 * @Author Lee
 * @Date 2023/12/18 16:41
 **/

package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var up = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(c *gin.Context) {
	// 升级成websocket
	conn, err := up.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("up.Upgrade error:", err)
	}
	// websocket读写逻辑
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("conn.ReadMessage error:", err)
			return
		}
		log.Println("receive:", string(p))

		conn.WriteMessage(messageType, []byte("hello, i am server"))
	}
}
