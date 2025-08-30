package logger

import (
	"google.golang.org/grpc/codes"
)

type LogConf struct {
	ServiceName string `json:"server_name" mapstructure:"server_name" json:"server_name" `
	Mode        string `json:"mode" json:"mode" `
	Path        string `json:"path" json:"path" `
	Level       string `json:"level" json:"level" `
	Compress    bool   `json:"compress" json:"compress" `
	KeepDays    int    `json:"keep_days" mapstructure:"keep_days" json:"keep_days" `
}

type Decider func(methodName string, err error) bool
type ErrorToCode func(err error) codes.Code
