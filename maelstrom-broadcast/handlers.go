package main

import (
	"encoding/json"
	"sync"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

const rpcBroadcast = "broadcast"
const rpcRead = "read"
const rpcTopology = "topology"

var n = maelstrom.NewNode()

func decodeJson(data []byte) (map[string]any, error) {
	var body map[string]any
	if err := json.Unmarshal(data, &body); err != nil {
		return nil, err
	}

	return body, nil
}

type HandlerRepository struct {
	mu       sync.RWMutex
	messages []int
}

func (h *HandlerRepository) broadcastHandler(msg maelstrom.Message) error {
	var body, err = decodeJson(msg.Body)
	if err != nil {
		return err
	}

	if val, ok := body["message"]; ok {
		h.mu.Lock()
		defer h.mu.Unlock()

		var f float64
		if f, ok = val.(float64); ok {
			h.messages = append(h.messages, int(f))
		}
	}

	return n.Reply(msg, map[string]string{
		"type": "broadcast_ok",
	})
}

func (h *HandlerRepository) readHandler(msg maelstrom.Message) error {
	var _, err = decodeJson(msg.Body)
	if err != nil {
		return err
	}

	h.mu.RLock()
	defer h.mu.RUnlock()
	messages := h.messages

	return n.Reply(msg, map[string]any{
		"type":     "read_ok",
		"messages": messages,
	})
}

func (h *HandlerRepository) topologyHandler(msg maelstrom.Message) error {
	var _, err = decodeJson(msg.Body)
	if err != nil {
		return err
	}

	return n.Reply(msg, map[string]string{
		"type": "topology_ok",
	})
}

// Handlers registry
func (h *HandlerRepository) RegisterHandlers() {
	n.Handle(rpcBroadcast, h.broadcastHandler)
	n.Handle(rpcRead, h.readHandler)
	n.Handle(rpcTopology, h.topologyHandler)
}
