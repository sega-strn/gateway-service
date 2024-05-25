package RESTHandlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	auth_proto "github.com/s21platform/auth-proto/auth-proto"
	grpc2 "github.com/s21platform/gateway-service/internal/repository/grpc"
	"google.golang.org/grpc/status"
	"net/http"
)

func AttachHandlers(r chi.Router, clients *grpc2.GrpcClients) {
	r.Route("/auth", func(authRouter chi.Router) {
		authRouter.Post("/login", GetAuth(clients))
	})
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetAuth(in Auth) func(w http.ResponseWriter, r *http.Request) {
	// Тут я уже должен получить интерфейс, который смогу заммокать. В частности достаточно просто замокать метод самого запроса, а все коннекшены - вынести отсюда
	//authService := auth.NewAuthServiceClient(cfg)
	return func(w http.ResponseWriter, r *http.Request) {
		data := &LoginData{}
		err := json.NewDecoder(r.Body).Decode(data)
		if err != nil || data.Username == "" || data.Password == "" {
			fmt.Println("some error for decode data")
			w.WriteHeader(http.StatusForbidden)
			return
		}
		request := auth_proto.LoginRequest{Username: data.Username, Password: data.Password}
		response, err := in.Login(context.Background(), &request)
		if err != nil {
			if statusError, ok := status.FromError(err); ok {
				http.Error(w, statusError.Message(), int(statusError.Code()))
				return
			}
		}
		http.SetCookie(w, &http.Cookie{
			Name:  "capy_token",
			Value: response.Jwt,
		})
		w.WriteHeader(200)
		return
	}
}
