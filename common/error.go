package common

import "net/http"

const INVALID_REQUEST = "INVALID_REQUEST"
const DB_ERROR = "DB_ERROR"

type AppError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`  // message error dùng để show cho client
	RootError  error  `json:"_"`        // ignore không return về root error
	Log        string `json:"logError"` // message error dùng để debug (show ở dev - hide ở prod)
	Key        string `json:"keyError"` // có thể sử dụng cho mục đích đa ngôn ngữ
	// (ví dụ: keyError = USER_NOT_FOUND , thì mỗi ngôn ngữ sẽ dựa vào key này để return cho ng dùng các message tương ứng với ngôn ngữ của ng dùng)
}

func NewErrorResponse(statusCode int, rootError error, message, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootError:  rootError,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

// AppError struct sẽ có kiểu "error" bởi vì nó có chứa method Error()
func (e *AppError) Error() string {
	return e.RootError.Error()
}

/* Truy root error
* Lý do: Do có thể các lỗi được wrap qua nhiều lớp
* (giả sử: transport error wrap bussines error -> business error wrap store error -> store error wrap DB error (đây là root error))
* Sử dụng kỹ thuật type-assertion (tạm gọi là ép kiểu) -> kiểm tra lỗi hiện tại có dạng AppError không
* -> Nếu ko phải dạng AppError => nó là root error
* -> Nếu vẫn đúng là dạng AppError (ok = true) => không phải root error => thực hiện tiếp "err.GetRootError()" (đệ quy)
 */
func (e *AppError) GetRootError() error {
	if err, ok := e.RootError.(*AppError); ok {
		return err.GetRootError()
	}
	return e.RootError
}

func InvalidRequest(err error, message string) *AppError {
	msg := "Invalid request"
	if message != "" {
		msg = message
	}
	return NewErrorResponse(
		http.StatusBadRequest,
		err,
		msg,
		err.Error(),
		INVALID_REQUEST,
	)
}

func BadRequest(err error, message string) *AppError {
	msg := "Bad request"
	if message != "" {
		msg = message
	}
	return NewErrorResponse(
		http.StatusBadRequest,
		err,
		msg,
		err.Error(),
		"BAD_REQUEST",
	)
}

func InternalServerError(err error) *AppError {
	return NewErrorResponse(
		http.StatusBadRequest,
		err,
		"Internal server error",
		err.Error(),
		"INTERNAL_SERVER_ERROR",
	)
}

func ErrDB(entity string, err error) *AppError {
	return NewErrorResponse(
		http.StatusInternalServerError,
		err,
		"Something went wrong with DB",
		err.Error(),
		DB_ERROR,
	)
}
