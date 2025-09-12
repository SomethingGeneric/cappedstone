#!/usr/bin/env python3

import socket
import threading

# Server configuration
HOST = '0.0.0.0'  # Standard loopback interface address (localhost)
PORT = 65432        # Port to listen on (non-privileged ports are > 1023)

def handle_client(conn, addr):
    """Handles communication with a connected client."""
    print(f"Connected by {addr}")
    with conn:
        while True:
            data = conn.recv(1024)  # Receive up to 1024 bytes
            if not data:
                break  # Client disconnected
            print(f"Received from {addr}: {data.decode()}")
            conn.sendall(data) # Echo back the received data
    print(f"Client {addr} disconnected")

def start_server():
    """Starts the TCP server."""
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.bind((HOST, PORT))
        s.listen()
        print(f"Server listening on {HOST}:{PORT}")
        while True:
            conn, addr = s.accept()  # Accept a new connection
            # Create a new thread to handle the client
            client_handler = threading.Thread(target=handle_client, args=(conn, addr))
            client_handler.start()

if __name__ == "__main__":
    start_server()
