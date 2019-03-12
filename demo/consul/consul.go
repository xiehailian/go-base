package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

var consulClient *api.Client

type Service struct {
	ID	        string
	IP          string
	Port 		int
	Name 		string
	CheckPath	string
}

func LocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

func InitConsul(dsn string) error {
	if dsn == "" {
		return  fmt.Errorf("consul address is nil")
	}
	config := api.DefaultConfig()
	config.Address = dsn
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}
	consulClient = client
	return nil
}


func NewService(id string, ip string, port int, name string, path string) *Service {
	return &Service{
		id,
		ip,
		port,
		name,
		path,
	}
}

func (s *Service) Register() {
	var tags []string
	service := &api.AgentServiceRegistration{
		ID:			s.ID,
		Name:		s.Name,
		Port:		s.Port,
		Address:	s.IP,
		Tags: 		tags,
		Check:  	&api.AgentServiceCheck{
			HTTP:	fmt.Sprintf("http://%s:%d%s", s.IP, s.Port, s.CheckPath),
			Interval: 	"10s",
			Timeout: 	"5s",
		},
	}
	if err := consulClient.Agent().ServiceRegister(service); err != nil {
        log.Fatal(err)
    }
}

func (s *Service) Deregister()  {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)

	<- quit
	if err := consulClient.Agent().ServiceDeregister(s.ID); err != nil {
		log.Println(-1, "consul deregist error, reason[%s]", err.Error())
	}
}


func RegistMockService() {
	service := NewService("mock", "127.0.0.1", 8888, "mock", "/health")
	service.Register()
	go service.Deregister()
}




func consulCheckHandler(w http.ResponseWriter, r * http.Request, p httprouter.Params)  {
	sendResponse(w, http.StatusOK, "ok")
}

func sendResponse(w http.ResponseWriter, sc int, resp string) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}


func RegisterHandler() *httprouter.Router{

	router := httprouter.New()

	router.GET("/health", consulCheckHandler)

	return router
}

func main()  {
	InitConsul("http://192.168.56.2:8500")
	RegistMockService()
	r := RegisterHandler()
	log.Fatal(http.ListenAndServe(":8888", r))
}


