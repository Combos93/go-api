package apiserver

import (
	"testing"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func TestAPIServer_configureLogger(t *testing.T) {
	type fields struct {
		config *Config
		logger *logrus.Logger
		router *mux.Router
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "All envs are right",
			fields: fields{
				config: NewConfig(map[string]string{
					"BIND_ADDR": ":8080",
					"LOG_LEVEL": "debug",
				}),
				logger: logrus.New(),
				router: mux.NewRouter(),
			},
			wantErr: false,
		},
		{
			name: "Log level is incorrect",
			fields: fields{
				config: NewConfig(map[string]string{
					"BIND_ADDR": ":8080",
					"LOG_LEVEL": "hahaha",
				}),
				logger: logrus.New(),
				router: mux.NewRouter(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &APIServer{
				config: tt.fields.config,
				logger: tt.fields.logger,
				router: tt.fields.router,
			}
			if err := s.configureLogger(); (err != nil) != tt.wantErr {
				t.Errorf("APIServer.configureLogger() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
