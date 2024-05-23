let socket = new Phoenix.Socket("ws://localhost:80/socket", {})

const [cpu, mem, disk] = [document.getElementById('cpu'), document.getElementById('mem'), document.getElementById('disk')]

socket.connect()

// Join a channel
let channel = socket.channel("metric:lobby", {})

// Log messages from the channel
channel.on("metrics_data", payload => {
  
  metric = payload.metrics[0]
  cpu.innerHTML = `CPU: ${metric.cpu_usage}`
  mem.innerHTML = `Memory: ${metric.memory_usage}`
  disk.innerHTML = `Disk: ${metric.disk_usage}`
})

// Handle successful join
channel.join()
  .receive("ok", resp => { console.log("Joined successfully", resp) })
  .receive("error", resp => { console.log("Unable to join", resp) })

// Listen for connection errors
socket.onError(() => console.log("There was an error with the connection!"))

// Listen for disconnection events
socket.onClose(() => console.log("The connection dropped"))