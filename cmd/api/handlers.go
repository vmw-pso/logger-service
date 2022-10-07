package main

import (
	"net/http"

	"github.com/vmw-pso/logger-service/data"

	"github.com/vmw-pso/toolkit"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (s *server) WriteLog() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestPayload JSONPayload

		_ = s.tools.ReadJSON(w, r, &requestPayload)

		event := data.LogEntry{
			Name: requestPayload.Name,
			Data: requestPayload.Data,
		}

		err := s.models.LogEntry.Insert(event)
		if err != nil {
			s.tools.ErrorJSON(w, err)
			return
		}

		resp := toolkit.JSONResponse{
			Error:   false,
			Message: "logged",
		}

		s.tools.WriteJSON(w, http.StatusAccepted, resp)
	}
}
