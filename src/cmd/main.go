package main

import (
    "go.uber.org/fx"
    "tictactoe/internal/datasource"
    "tictactoe/internal/web"
    "net/http"
)

func main() {
    fx.New(
        fx.Provide(
            datasource.NewRepository,
            web.NewHandler,
        ),
        fx.Invoke(func(h *web.Handler) {
            http.HandleFunc("POST /game/{current_game_UUID}", h.PlayerNewMove)
            go http.ListenAndServe(":8080", nil)
			fs := http.FileServer(http.Dir("./static"))
			http.Handle("/", fs)
        }),
    ).Run()
}
