package response

type Response struct {
	Status string `json:"status"`          // Буде вказано Error або OK
	Error  string `json:"error,omitempty"` // omitempty параметр який можна вказати в страктегові json, що вказує, коли параметр в json буде пустий це поле буде відсутнє
}

const (
	StatusOk    = "OK"
	StatusError = "Error"
)

func OK() Response {
	return Response{
		Status: StatusOk,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}
