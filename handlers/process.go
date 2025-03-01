package handlers

import (
	"bufio"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

func HandlerYtDescription(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusAccepted)
	fmt.Fprintf(ctx, "Hello, %s!\n", ctx.UserValue("id"))
}

func HandlerStream(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)
	if id == "" {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(`{"error": "ID é obrigatório"}`)
		return
	}

	// Configura o stream (chunked response)
	ctx.SetContentType("text/event-stream")
	ctx.Response.Header.Set("Cache-Control", "no-cache")
	ctx.Response.Header.Set("Connection", "keep-alive")
	ctx.Response.SetBodyStreamWriter(func(w *bufio.Writer) {
		defer w.Flush()

		for i := 0; i < 10; i++ {
			_, err := w.Write([]byte(fmt.Sprintf("chunk %d\n", i)))
			time.Sleep(1 * time.Second)
			if err != nil {
				// handling error
				return
			}

			w.Flush()
		}
	})
}
