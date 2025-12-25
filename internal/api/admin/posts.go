package admin

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type createPostRequest struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	Published bool   `json:"published"`
}

func CreatePost(db *sql.DB, adminToken string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// only allow POST
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// admin auth check
		auth := r.Header.Get("Authorization")
		if adminToken == "" || auth != "Bearer "+adminToken {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		// parse request body
		var req createPostRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		if strings.TrimSpace(req.Title) == "" || strings.TrimSpace(req.Content) == "" {
			http.Error(w, "title and content required", http.StatusBadRequest)
			return
		}

		// generate slug
		slug := strings.ToLower(req.Title)
		slug = strings.TrimSpace(slug)
		slug = strings.ReplaceAll(slug, " ", "-")

		// insert post
		_, err := db.Exec(`
			INSERT INTO posts (title, slug, content, published, created_at)
			VALUES (?, ?, ?, ?, ?)
		`,
			req.Title,
			slug,
			req.Content,
			boolToInt(req.Published),
			time.Now(),
		)

		if err != nil {
			http.Error(w, "failed to insert post", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
