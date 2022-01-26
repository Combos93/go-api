package apiserver

import (
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestAPIServer_configureLogger(t *testing.T) {
	type fields struct {
		config *Config
		logger *logrus.Logger
		router *http.ServeMux
	}
	tests := []struct {
		name    string
		fields  fields
		isError bool
	}{
		{
			name: "All envs are right",
			fields: fields{
				config: buildConfig("debug"),
				logger: logrus.New(),
				router: http.NewServeMux(),
			},
			isError: false,
		},
		{
			name: "Log level is incorrect",
			fields: fields{
				config: buildConfig("hahaha"),
				logger: logrus.New(),
				router: http.NewServeMux(),
			},
			isError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &APIServer{
				config: tt.fields.config,
				logger: tt.fields.logger,
				router: tt.fields.router,
			}
			if err := s.configureLogger(); (err != nil) != tt.isError {
				t.Errorf("APIServer.configureLogger() error = %v, isError %v", err, tt.isError)
			}
		})
	}
}

func buildConfig(logLevel string) *Config {
	return NewConfig(map[string]string{
		"BIND_ADDR": ":8080",
		"LOG_LEVEL": logLevel,
	})
}
