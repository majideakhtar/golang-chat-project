var socket = new WebSocket('ws://localhost:9000/ws');

let connect = (cb) => {
  console.log("connecting");

  socket.onopen = () => {
    console.log("Successfully connected!")
  }

  socket.onmessage = (msg) => {
    console.log("Message from Websocket: ", msg)
  }

  socket.onclose = (event) => {
    console.log("Socket Closed Connection!: ", event)
  }

  socket.onerror = (error) => {
    console.log("Socket Error: ", error);
  }
}; 

let sendMsg = (msg) =>{
  console.log("Sending msg: ", msg);
  socket.send(msg)
}

export {connect, sendMsg} 