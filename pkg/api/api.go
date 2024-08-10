package api

import (
	"github.com/go-chi/chi/v5/middleware"
	"go-user-comments-postgres/pkg/db/models"
	"net/http"
)

// start api with the pgdb and return a chi router
func StartAPI(pgdb *pg.DB) *chi.Mux {
	//get the router
	r := chi.NewRouter()
	//add middleware
	//in this case we will store our DB to use it later
	r.Use(middleware.Logger, middleware.WithValue("DB", pgdb))

	//routes for our service
	r.Route("/comments", func(r chi.Router) {
		r.Post("/", createComment)
		r.Get("/", getComments)
		r.Get("/{commentID}", getCommentByID)
		r.Put("/{commentID}", updateCommentByID)
		r.Delete("/{commentID}", deleteCommentByID)
	})

	//test route to make sure everything works
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("up and running"))
	})

	return r
}

// -- Requests and Responses
type CommentRequest struct {
	Comment string `json:"comment"`
	UserID  int64  `json:"user_id"`
}

type CommentResponse struct {
	Success bool            `json:"success"`
	Error   string          `json:"error"`
	Comment *models.Comment `json:"comment"`
}
