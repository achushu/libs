package websocket

import (
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/achushu/libs/out"
	gorilla "github.com/gorilla/websocket"
)

type Manager struct {
	sockets map[string]*gorilla.Conn
	rng     *rand.Rand
	mu      sync.RWMutex
}

func NewManager() *Manager {
	m := &Manager{
		sockets: make(map[string]*gorilla.Conn),
		rng:     rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	return m
}

func (m *Manager) Add(ws *gorilla.Conn) string {
	id := m.generateConnID()
	m.mu.Lock()
	defer m.mu.Unlock()
	m.sockets[id] = ws

	return id
}

func (m *Manager) Get(id string) (*gorilla.Conn, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	ws, ok := m.sockets[id]
	return ws, ok
}

func (m *Manager) Send(msg []byte, id string) error {
	defer func() {
		if r := recover(); r != nil {
			out.Println("websocket/manager", "Panicked sending to conn ", id, ":\n", r)
		}
	}()

	if ws, ok := m.Get(id); ok {
		return ws.WriteMessage(gorilla.TextMessage, msg)
	}
	return NewConnectionError(ClientNotFound, id)
}

func (m *Manager) CloseConn(id string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if ws, ok := m.sockets[id]; ok {
		err := ws.Close()
		if err != nil {
			out.Println("error closing websocket: ", err)
		}
		delete(m.sockets, id)
		out.Println("removed connection ", id)
		return true
	}
	return false
}

func (m *Manager) Count() int {
	return len(m.sockets)
}

func (m *Manager) Info(id string) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if ws, ok := m.sockets[id]; ok {
		// TODO: Store info in a struct: start time, addr, mesgs in/out, last msg time, heartbeat
		ws.RemoteAddr()
	}
}

func (m *Manager) generateConnID() string {
	return strconv.FormatInt(m.rng.Int63n(32768)+32768, 16)
}
