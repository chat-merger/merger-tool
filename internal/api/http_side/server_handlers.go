package http_side

import (
	"chatmerger/internal/domain/model"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

func (s *HttpSideServer) createClientHandler(w http.ResponseWriter, r *http.Request) {
	// парсирнг тела как json структуры
	var input model.CreateClient
	var name = r.PostFormValue("name")
	if name == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		input = model.CreateClient{Name: name}
	}
	// создание клиента через юзкейс
	err := s.CreateClient(input)
	if err != nil {
		log.Println(err)
		log.Printf("err = s.CreateClient(input) err: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = s.executeTemplWithClientsTable(w)
	if err != nil {
		log.Printf("execute templ  with clients: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *HttpSideServer) getClientsHandler(w http.ResponseWriter, r *http.Request) {
	err := s.executeTemplWithClientsTable(w)
	if err != nil {
		log.Printf("execute templ  with clients: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *HttpSideServer) deleteClientHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := s.DeleteClients([]model.ID{model.NewID(idStr)})
	if err != nil {
		log.Printf("failed delete clients: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *HttpSideServer) index(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile("web/index.html")
	if err != nil {
		log.Printf("failed read index.html file\n")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(file)
}

func (s *HttpSideServer) executeTemplWithClientsTable(wr io.Writer) error {
	var tmpl = template.Must(template.ParseFiles("web/clients_table.html"))

	clients, err := s.ClientsList()
	if err != nil {
		return fmt.Errorf("get clients list: %s", err)
	}
	conns, err := s.ConnectedClientsList()
	if err != nil {
		return fmt.Errorf("get connected list: %s", err)
	}
	var connectedNames []string
	for _, conn := range conns {
		connectedNames = append(connectedNames, conn.Name)
	}
	var resp = map[string]any{
		"Clients":     clients,
		"Connections": connectedNames,
	}
	return tmpl.Execute(wr, resp)
}
