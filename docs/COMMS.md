# Communication Spec
TCP sockets with TLS encryption

Agent will seek server response on port 5900 (default, can be changed.)

## General Flow
Agent exec should support either "daemon" mode, where you just run the exec, and every <X> interval it reaches out to the server to see if there are any tasks for it, or in a "cron" mode, where it immediately checks with the server for any tasks, does them, and then exits.

The interval for daemon mode can be configured in the conf file which will also have the ip/port for server.

## Cert Management
The endpoint exec should trust the cert advertised      by the server on first start, and fail later if it changed (can be a flag in exec arguements as well).