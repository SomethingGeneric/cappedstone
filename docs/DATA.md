# Data storage
We'll use flat-file at least for the app demo thru semester 1

Maybe look into some SQL server or a more production-grade NoSQL solution in the future?

## Endpoint
The "endpoints" file (later an SQL table?) is:
```md
Endpoint ID | IP | OS FAMILY | OS | Tasks <List> | Last Seen | Next Expected At

1 | 192.168.1.13 | Linux | Debian | Task_list_ref | 10:00:00 11/2/25 | 10:05:00 11/2/25
```

## Task List
For each host, there's a list of tasks, as such:
```md
Task ID | Task_Name | Args | Assigned_At |

1 | checkupdates | NULL | 13:00 11/02/25
```