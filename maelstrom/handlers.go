package main

import (
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

const rpcEcho = "echo"
const rpcGenerate = "generate"

var n = maelstrom.NewNode()

// Echo
func EchoHandler(msg maelstrom.Message) error {
	var body, err = decodeJson(msg.Body)
	if err != nil {
		return err
	}

	body["type"] = "echo_ok"

	return n.Reply(msg, body)
}

// Unique ID Generation
func GenerateHandler(msg maelstrom.Message) error {
	var body, err = decodeJson(msg.Body)
	if err != nil {
		return err
	}

	var uuid = ""
	uuid, err = generateUUIDUsingCommand()
	if err != nil {
		return err
	}

	body["id"] = uuid
	body["type"] = "generate_ok"

	return n.Reply(msg, body)
}

// Handlers registry
func RegisterHandlers() {
	n.Handle(rpcEcho, EchoHandler)
	n.Handle(rpcGenerate, GenerateHandler)
}
