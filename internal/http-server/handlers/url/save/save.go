package save

import (
	resp "RestAPITest/internal/lib/api/response"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"
	"net/http"
)

type Request struct {
	URL   string `json:"url" validate:"required,url"` // опис запитів які будуть надсилатися
	Alias string `json:"alias,omitempty"`
}

type Response struct { // відповідь яку ми повертаємо
	resp.Response
	Alias string `json:"alias,omitempty"`
}

type URLSaver interface { // URLSaver інтерфейс storage
	SaveURL(urlToSave string, alias string) (int64, error)
}

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc { // конструктор для handler
	return func(w http.ResponseWriter, r *http.Request) { // можемо передати певні параметри які будуть вказані в обробнику
		const op = "handlers.url.save.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", slog.Error(err)) // пишемо помилку в log

			render.JSON(w, r, resp.Error("failed to decode request")) // повертаємо json-відповідь клієнту

			return
		}

		log.Info("request body decoded", slog.Any("request", req))
	}
}
