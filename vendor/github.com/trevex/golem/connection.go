/*

   Copyright 2013 Niklas Voss

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

*/

package golem

import (
	"github.com/gorilla/websocket"
	"reflect"
	"time"
)

const (
	// Time allowed to write a message to the client.
	writeWait = 10 * time.Second
	// Time allowed to read the next message from the client.
	readWait = 60 * time.Second
	// Send pings to client with this period. Must be less than readWait.
	pingPeriod = (readWait * 9) / 10
	// Maximum message size allowed from client.
	maxMessageSize = 512
	// Outgoing default channel size.
	sendChannelSize = 512
)

var (
	defaultConnectionExtension = reflect.ValueOf(nil)
)

// SetDefaultConnectionExtension sets the initial extension used by all freshly instanced routers.
// For more information see the Router SetConnectionExtension() - method.
func SetDefaultConnectionExtension(constructor interface{}) {
	defaultConnectionExtension = reflect.ValueOf(constructor)
}

// Connection holds information about the underlying WebSocket-Connection,
// the associated router and the outgoing data channel.
type Connection struct {
	// The websocket connection.
	socket *websocket.Conn
	// Associated router.
	router *Router
	// Buffered channel of outbound messages.
	send chan *message
	//
	extension interface{}
}

// Create a new connection using the specified socket and router.
func newConnection(s *websocket.Conn, r *Router) *Connection {
	return &Connection{
		socket:    s,
		router:    r,
		send:      make(chan *message, sendChannelSize),
		extension: nil,
	}
}

// Register connection and start writing and reading loops.
func (conn *Connection) run() {
	hub.register <- conn
    readMode := websocket.TextMessage
    writeMode := websocket.TextMessage
    if conn.router.protocol.GetReadMode() != TextMode {
        readMode = websocket.BinaryMessage
    }
    if conn.router.protocol.GetWriteMode() != TextMode {
        writeMode = websocket.BinaryMessage
    }
    if conn.router.useHeartbeats {
        go conn.writePumpHeartbeat(writeMode)
        conn.readPumpHeartbeat(readMode)
    } else {
        go conn.writePump(writeMode)
        conn.readPump(readMode)
    }
}

func (conn *Connection) extend(e interface{}) {
	conn.extension = e
}

// Emit event with provided data. The data will be automatically marshalled and packed according
// to the active protocol of the router the connection belongs to.
func (conn *Connection) Emit(event string, data interface{}) {
	conn.send <- &message{
		event: event,
		data:  data,
	}
}

// Close closes and cleans up the connection.
func (conn *Connection) Close() {
	hub.unregister <- conn
}

// Helper for writing to socket with deadline.
func (conn *Connection) write(mode int, payload []byte) error {
	conn.socket.SetWriteDeadline(time.Now().Add(writeWait))
	return conn.socket.WriteMessage(mode, payload)
}

/*
 * Pumps with Heartbeat.
 */

func (conn *Connection) readPumpHeartbeat(mode int) {
	defer func() {
		hub.unregister <- conn
		conn.socket.Close()
		conn.router.closeFunc(conn)
	}()
	conn.socket.SetReadLimit(maxMessageSize)
	conn.socket.SetReadDeadline(time.Now().Add(readWait))
    conn.socket.SetPongHandler(func(string) error {
        conn.socket.SetReadDeadline(time.Now().Add(readWait))
        return nil
    })
	for {
		mm, message, err := conn.socket.ReadMessage()
		if err != nil {
			break
		}
        if mm == mode {
            conn.router.processMessage(conn, message)
        }
	}
}

func (conn *Connection) writePumpHeartbeat(mode int) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		conn.socket.Close() // Necessary to force reading to stop
	}()
	for {
		select {
		case message, ok := <-conn.send:
			if ok {
				if data, err := conn.router.protocol.MarshalAndPack(message.event, message.data); err == nil {
					if err := conn.write(mode, data); err != nil {
						return 
					}
				} else {
                    // TODO: logging
                }
			} else {
				conn.write(websocket.CloseMessage, []byte{})
				return
			}
		case <-ticker.C:
			if err := conn.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

/*
 * Pumps without Heartbeat
 */

func (conn *Connection) readPump(mode int) {
	defer func() {
		hub.unregister <- conn
		conn.socket.Close()
		conn.router.closeFunc(conn)
	}()
	conn.socket.SetReadLimit(maxMessageSize)
	for {
		mm, message, err := conn.socket.ReadMessage()
		if err != nil {
			break
		}
        if mm == mode {
            conn.router.processMessage(conn, message)
        }
	}
}

func (conn *Connection) writePump(mode int) {
	defer func() {
		conn.socket.Close() // Necessary to force reading to stop
	}()
	for {
		select {
		case message, ok := <-conn.send:
			if ok {
				if data, err := conn.router.protocol.MarshalAndPack(message.event, message.data); err == nil {
					if err := conn.write(mode, data); err != nil {
						return 
					}
				} else {
                    // TODO: logging
                }
			} else {
				conn.write(websocket.CloseMessage, []byte{})
				return
			}
		}
	}
}


