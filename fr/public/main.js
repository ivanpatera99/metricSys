let socket = new Phoenix.Socket("ws://localhost:80/socket", {})

socket.connect()

// Join a channel
let channel = socket.channel("metric:lobby", {})

// Log messages from the channel
channel.on("metrics_data", payload => {
  console.log("Received message:", payload)
})

// Handle successful join
channel.join()
  .receive("ok", resp => { console.log("Joined successfully", resp) })
  .receive("error", resp => { console.log("Unable to join", resp) })

// Listen for connection errors
socket.onError(() => console.log("There was an error with the connection!"))

// Listen for disconnection events
socket.onClose(() => console.log("The connection dropped"))