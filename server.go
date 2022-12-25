package main

import (
	"fmt"
	"net/http"

	"message/controller"
	router "message/router"
)

var (
	messageControler controller.MessageController = controller.NewMessageController()
	httpRouter       router.Router                = router.NewMuxRouter()
)

func main() {
	const port string = ":9000"
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running")
	})
	httpRouter.GET("/messages", messageControler.GetMessages)
	httpRouter.POST("/messages", messageControler.AddMessages)
	httpRouter.SERVE(port)
}
