package xhttp

import (
	"net"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const MB = 1 << (10 * 2)

var ErrBodySizeLimit = errors.New("body size too large")

// Config HTTP客户端相关配置
type Config struct {
	HTTPTimeout           time.Duration // HTTP请求超时时间
	DialTimeout           time.Duration // 拨号超时时间
	DialKeepAlive         time.Duration // 拨号保持连接时间
	MaxIdleConns          int           // 最大空闲连接数
	MaxIdleConnsPerHost   int           // 每个主机最大空闲连接数
	MaxConnsPerHost       int           // 每个主机最大连接数
	IdleConnTimeout       time.Duration // 空闲连接超时时间
	ResponseHeaderTimeout time.Duration // 读取响应头超时时间
	ExpectContinueTimeout time.Duration // 期望继续超时时间
	TLSHandshakeTimeout   time.Duration // TLS握手超时时间
	ForceAttemptHTTP2     bool          // 允许尝试启用HTTP/2
}

// GetDefaultConfig 获取默认HTTP客户端相关配置
func GetDefaultConfig() *Config {
	return &Config{
		HTTPTimeout:           20 * time.Second,
		DialTimeout:           15 * time.Second,
		DialKeepAlive:         30 * time.Second,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   100,
		MaxConnsPerHost:       100,
		IdleConnTimeout:       60 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 5 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ForceAttemptHTTP2:     true,
	}
}

// Client HTTP拓展客户端结构详情
type Client struct {
	*http.Client
}

// NewClient 新建HTTP拓展客户端
func NewClient(c *Config) *Client {
	return &Client{Client: NewHTTPClient(c)}
}

// NewHTTPClient 新建HTTP客户端
func NewHTTPClient(c *Config) *http.Client {
	if c == nil {
		c = GetDefaultConfig()
	}

	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   c.DialTimeout,
			KeepAlive: c.DialKeepAlive,
		}).DialContext,
		MaxIdleConns:          c.MaxIdleConns,
		MaxIdleConnsPerHost:   c.MaxIdleConnsPerHost,
		MaxConnsPerHost:       c.MaxConnsPerHost,
		IdleConnTimeout:       c.IdleConnTimeout,
		ResponseHeaderTimeout: c.ResponseHeaderTimeout,
		ExpectContinueTimeout: c.ExpectContinueTimeout,
		TLSHandshakeTimeout:   c.TLSHandshakeTimeout,
		ForceAttemptHTTP2:     c.ForceAttemptHTTP2,
	}

	client := &http.Client{
		Timeout:   c.HTTPTimeout,
		Transport: tr,
	}

	return client
}
