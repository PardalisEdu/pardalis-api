package utils

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	tests := []struct {
		name       string
		data       interface{}
		wantStatus int
		wantErr    bool
	}{
		{
			name:       "Valid data",
			data:       map[string]string{"key": "value"},
			wantStatus: http.StatusOK,
			wantErr:    false,
		},
		{
			name:       "Nil data",
			data:       nil,
			wantStatus: http.StatusOK,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			err := WriteJSON(w, tt.wantStatus, tt.data)

			if (err != nil) != tt.wantErr {
				t.Errorf("WriteJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if w.Code != tt.wantStatus {
				t.Errorf("WriteJSON() status = %v, want %v", w.Code, tt.wantStatus)
			}

			if !tt.wantErr {
				contentType := w.Header().Get("Content-Type")
				if contentType != "application/json" {
					t.Errorf("WriteJSON() content-type = %v, want application/json", contentType)
				}
			}
		})
	}
}

func TestParseJSON(t *testing.T) {
	type testStruct struct {
		Key string `json:"key"`
	}

	tests := []struct {
		name    string
		payload string
		want    testStruct
		wantErr bool
	}{
		{
			name:    "Valid JSON",
			payload: `{"key":"value"}`,
			want:    testStruct{Key: "value"},
			wantErr: false,
		},
		{
			name:    "Invalid JSON",
			payload: `{"key":}`,
			wantErr: true,
		},
		{
			name:    "Empty body",
			payload: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got testStruct
			body := bytes.NewBufferString(tt.payload)
			req := httptest.NewRequest("POST", "/", body)

			err := ParseJSON(req, &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && got != tt.want {
				t.Errorf("ParseJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTokenFromRequest(t *testing.T) {
	tests := []struct {
		name        string
		headerToken string
		queryToken  string
		want        string
	}{
		{
			name:        "Header token",
			headerToken: "Bearer token123",
			queryToken:  "",
			want:        "Bearer token123",
		},
		{
			name:        "Query token",
			headerToken: "",
			queryToken:  "token123",
			want:        "token123",
		},
		{
			name:        "Both tokens - header priority",
			headerToken: "Bearer token123",
			queryToken:  "token456",
			want:        "Bearer token123",
		},
		{
			name:        "No token",
			headerToken: "",
			queryToken:  "",
			want:        "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			if tt.headerToken != "" {
				req.Header.Set("Authorization", tt.headerToken)
			}
			if tt.queryToken != "" {
				q := req.URL.Query()
				q.Add("token", tt.queryToken)
				req.URL.RawQuery = q.Encode()
			}

			if got := GetTokenFromRequest(req); got != tt.want {
				t.Errorf("GetTokenFromRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
