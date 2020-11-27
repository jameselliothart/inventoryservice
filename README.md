# Go Web Services

## Setup

`docker run --rm -v $HOME/dockervolumes/GoWebServices:/var/lib/mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password123 mysql`

## Notes

`_ "github.com/go-sql-driver/mysql"` imports just for side effects (working with MySQL) - no functions are actually used.

Test websocket via browser console:

```js
let ws = new WebSocket("ws://localhost:5000/websocket")
ws.send(JSON.stringify({data: "test message", type: "test"}))
```
