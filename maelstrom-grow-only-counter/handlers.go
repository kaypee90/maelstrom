package main

import (
	"context"
	"encoding/json"
	"sync"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

const rpcAdd = "add"
const rpcRead = "read"

var n = maelstrom.NewNode()
var kv = maelstrom.NewSeqKV(n)
var ctx = context.Background()

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

func (h *HandlerRepository) addHandler(msg maelstrom.Message) error {
	var body, err = decodeJson(msg.Body)
	if err != nil {
		return err
	}

	if val, ok := body["delta"].(float64); ok {
		delta := int(val)
		key := n.ID()
		counter, _ := kv.ReadInt(ctx, key)
		counter += delta
		kv.Write(ctx, key, counter)

		body["type"] = "add_ok"
	}

	delete(body, "delta")

	return n.Reply(msg, body)
}

func (h *HandlerRepository) readHandler(msg maelstrom.Message) error {
	var body, err = decodeJson(msg.Body)
	if err != nil {
		return err
	}

	value, _ := kv.ReadInt(ctx, n.ID())

	body["type"] = "read_ok"
	body["value"] = value

	return n.Reply(msg, body)
}

// Handlers registry
func (h *HandlerRepository) RegisterHandlers() {
	n.Handle(rpcAdd, h.addHandler)
	n.Handle(rpcRead, h.readHandler)
}
