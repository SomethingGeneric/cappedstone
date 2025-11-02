# API Spec

When a client opens a socket to the server, it should issue `tasks id`, which should return a list of task IDs and associated commands to execute:

```json
{
    "tasks": [
        {
            "id": "xyyhgiufngf",
            "command": "apt update && apt list --upgradable"
        },
        {
            "id": "uhifughufg",
            "command": "df -h"
        }
    ]
}
```

The client should then return something like the following after executing the tasks:

```json
{
    "task_results": [
        {
            "id": "xyyhgiufngf",
            "exit_code": 0,
            "stdout": "found 123 packages to update"
        },
        {
            "id": "uhifughufg",
            "exit_code": 0,
            "stdout": "<stdout here>"
        }
    ]
}
```