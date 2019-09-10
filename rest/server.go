package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nymtech/directory-server/metrics"
	"github.com/nymtech/directory-server/pki"
)

// Config defines the values passed into the REST Service
type Config struct {
	Addr    string
	Metrics metrics.Service
	Port    int
	PKI     pki.Service
}

// Server gives us a place to store values for our REST API
type Server struct {
	controllers []controller
	port        int
	router      *gin.Engine
	pki         pki.Service
}

// New returns a new REST API server
func New(cfg *Config) *Server {
	var controllers []controller
	pkiCfg := &pki.Config{}
	metricsCfg := &metrics.Config{}

	controllers = append(controllers, pki.New(pkiCfg))
	controllers = append(controllers, metrics.New(metricsCfg))

	s := &Server{
		controllers: controllers,
		port:        cfg.Port,
		pki:         cfg.PKI,
	}
	s.router = s.makeRouter(controllers...)

	return s
}

// Run the REST server.
func (srv *Server) Run() {
	srv.router.Run()
}

// Node is a mixnode or coconode
type Node struct {
	PubSigningKey   string
	PubTransportKey string
	Up              bool
}
