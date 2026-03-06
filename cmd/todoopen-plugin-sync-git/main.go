package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

const protocolVersion = "todoopen.plugin.v1"

type health struct {
	State   string `json:"state"`
	Message string `json:"message,omitempty"`
}

type handshakeResponse struct {
	ProtocolVersion string   `json:"protocol_version"`
	Name            string   `json:"name"`
	Kind            string   `json:"kind"`
	Capabilities    []string `json:"capabilities,omitempty"`
	Health          health   `json:"health"`
}

type requestEnvelope struct {
	ID      string         `json:"id"`
	Method  string         `json:"method"`
	Payload map[string]any `json:"payload,omitempty"`
}

type pluginError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
}

type responseEnvelope struct {
	ID      string         `json:"id"`
	Payload map[string]any `json:"payload,omitempty"`
	Error   *pluginError   `json:"error,omitempty"`
}

func main() {
	enc := json.NewEncoder(os.Stdout)
	dec := json.NewDecoder(bufio.NewReader(os.Stdin))

	hs := handshakeResponse{
		ProtocolVersion: protocolVersion,
		Name:            "git",
		Kind:            "sync",
		Capabilities:    []string{"pull", "push", "status"},
		Health:          health{State: "ready", Message: "example git sync plugin"},
	}
	if err := enc.Encode(hs); err != nil {
		fmt.Fprintf(os.Stderr, "write handshake: %v\n", err)
		os.Exit(1)
	}

	for {
		var req requestEnvelope
		if err := dec.Decode(&req); err != nil {
			return
		}

		resp := handle(req)
		if err := enc.Encode(resp); err != nil {
			fmt.Fprintf(os.Stderr, "write response: %v\n", err)
			return
		}
	}
}

func handle(req requestEnvelope) responseEnvelope {
	switch req.Method {
	case "status":
		return responseEnvelope{
			ID: req.ID,
			Payload: map[string]any{
				"state":   "ready",
				"message": "git sync example ready",
			},
		}
	case "push":
		return responseEnvelope{
			ID: req.ID,
			Payload: map[string]any{
				"ok": true,
			},
		}
	case "pull":
		return responseEnvelope{
			ID: req.ID,
			Payload: map[string]any{
				"tasks": []any{},
			},
		}
	default:
		return responseEnvelope{
			ID: req.ID,
			Error: &pluginError{
				Code:    "not_supported",
				Message: "unsupported method",
				Detail:  req.Method,
			},
		}
	}
}
