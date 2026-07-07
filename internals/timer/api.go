package timer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (t *Timer) StartServer(addr string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var id, callbackURL string
		var fireInSec int

		contentType := r.Header.Get("Content-Type")
		if strings.HasPrefix(contentType, "application/json") {
			var req struct {
				ID          string `json:"id"`
				TimerTime   int    `json:"timer_time"`
				CallbackURL string `json:"callback_url"`
			}
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
				return
			}
			id = req.ID
			fireInSec = req.TimerTime
			callbackURL = req.CallbackURL
		} else {
			id = r.FormValue("id")
			fireInSecStr := r.FormValue("timer_time")
			callbackURL = r.FormValue("callback_url")

			if id == "" || fireInSecStr == "" || callbackURL == "" {
				http.Error(w, "Missing required parameters: id, timer_time, or callback_url", http.StatusBadRequest)
				return
			}

			var err error
			fireInSec, err = strconv.Atoi(fireInSecStr)
			if err != nil {
				http.Error(w, "Invalid timer_time, must be an integer", http.StatusBadRequest)
				return
			}
		}

		if id == "" || callbackURL == "" || fireInSec <= 0 {
			http.Error(w, "Invalid parameters: id, callback_url and a positive timer_time are required", http.StatusBadRequest)
			return
		}

		task := &TimerTask{
			ID:          id,
			FireAt:      time.Now().Add(time.Duration(fireInSec) * time.Second),
			CallBackURL: callbackURL,
		}

		heapID := t.AddTask(task)
		log.Printf("Routed task %s to heap %d (expires in %ds)", task.ID, heapID, fireInSec)

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "success: routed to heap %d\n", heapID)
	})

	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Printf("Listening for HTTP API requests on %s...", addr)
	return server.ListenAndServe()
}
