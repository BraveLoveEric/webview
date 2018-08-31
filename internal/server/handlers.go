package server

import "net/http"

func (s *Server) index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Vary", "Accept-Encoding")

		w.Write([]byte(`Index`))
	}
}

func (s *Server) notFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Vary", "Accept-Encoding")
		w.WriteHeader(http.StatusNotFound)

		w.Write([]byte(`Not Found`))
	}
}

// getHero fetches the data for a hero
func (s *Server) getHero() http.HandlerFunc {
	type getHeroResponse struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		resp := []byte(`hero data`)
		writeResp(w, resp)
	}
}
