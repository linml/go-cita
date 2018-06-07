package connection

import (
	"sync"
	"time"

	"github.com/golang/glog"
	networkConfig "github.com/yejiayu/go-cita/config/network"
	"github.com/yejiayu/go-cita/network/protocol"
)

type Manager interface {
	Run(quit chan<- error)
	UpdateConfig(config networkConfig.Config)

	Available() []networkConfig.Peer

	Send(key string, msg *protocol.Message)
}

func NewManager(config networkConfig.Config) Manager {
	peers := make(map[networkConfig.Peer]bool)
	for _, peer := range config.Peers {
		peers[peer] = true
	}

	return &manager{
		id:        config.ID,
		peers:     peers,
		conns:     make(map[networkConfig.Peer]*connection),
		unReadyCh: make(chan networkConfig.Peer),
	}
}

type manager struct {
	mu sync.RWMutex

	id    uint32
	peers map[networkConfig.Peer]bool

	conns     map[networkConfig.Peer]*connection
	unReadyCh chan networkConfig.Peer
}

func (m *manager) Run(quit chan<- error) {
	for peer := range m.peers {
		conn, err := newConnection(m.id, peer)
		if err != nil {
			glog.Error(err)
			m.unReadyCh <- peer
			continue
		}

		m.conns[peer] = conn
	}

	go m.checkReadied()
}

func (m *manager) UpdateConfig(config networkConfig.Config) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.id = config.ID

	for _, newPeer := range config.Peers {
		_, ok := m.conns[newPeer]
		if !ok {
			m.unReadyCh <- newPeer
		}
	}

	// Close failed peer
	for peer, conn := range m.conns {
		_, ok := m.peers[peer]
		if ok {
			continue
		}
		if err := conn.close(); err != nil {
			glog.Error(err)
		}
	}
}

func (m *manager) Available() []networkConfig.Peer {
	peers := []networkConfig.Peer{}

	m.mu.RLock()
	defer m.mu.RUnlock()

	for peer := range m.peers {
		peers = append(peers, peer)
	}

	return peers
}

func (m *manager) Send(key string, msg *protocol.Message) {
	origin := msg.Origin()
	opType := msg.Optype()
	sendMsg := protocol.NewMessage(opType, m.id, msg.Payload())

	m.mu.RLock()
	defer m.mu.Unlock()

	switch opType {
	case protocol.OpTypeSingle:
		for peer, conn := range m.conns {
			if peer.ID == origin {
				err := conn.send(key, sendMsg.Raw())
				if err != nil {
					glog.Errorf("network: the message send to %-v failed, %s\n", peer, err.Error())
				}
			}
		}
	case protocol.OpTypeSubtract:
		for peer, conn := range m.conns {
			if peer.ID != origin {
				err := conn.send(key, sendMsg.Raw())
				if err != nil {
					glog.Errorf("network: the message send to %-v failed, %s\n", peer, err.Error())
				}
			}
		}
	case protocol.OpTypeBroadcast:
		for peer, conn := range m.conns {
			err := conn.send(key, sendMsg.Raw())
			if err != nil {
				glog.Errorf("network: the message send to %-v failed, %s\n", peer, err.Error())
			}
		}
	}
}

func (m *manager) checkReadied() {
	for unReadiedPeer := range m.unReadyCh {
		time.Sleep(1 * time.Second)

		m.mu.Lock()
		_, ok := m.peers[unReadiedPeer]
		if !ok {
			m.mu.Unlock()
			continue
		}

		conn, err := newConnection(m.id, unReadiedPeer)
		if err != nil {
			m.mu.Unlock()
			m.unReadyCh <- unReadiedPeer
			continue
		}

		m.conns[unReadiedPeer] = conn
		m.mu.Unlock()
	}
}
