package save

import (
	resp "RestAPITest/internal/lib/api/response"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
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

const aliasLength = 5 // також можна перенести в config

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
			log.Error("failed to decode request body", slog.Error) // пишемо помилку в log

			render.JSON(w, r, resp.Error("failed to decode request")) // повертаємо json-відповідь клієнту

			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil { // створюємо валідатора, який має провалідувати структуру
			validateErr := err.(validator.ValidationErrors) // якщо буде помилка, то він поверне помилку даного типу
			log.Error("invalid request", slog.Error)        // залогуємо дану помилку

			//render.JSON(w, r, resp.Error("invalid request"))
			render.JSON(w, r, resp.ValidationError(validateErr)) // формуємо запит з повідомленням про помилку

			return
		}

		alias := req.Alias
		if alias == "" {
			alias = random.NewRandomString(aliasLength)
		}

	}
}
