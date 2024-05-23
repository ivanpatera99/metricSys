const setupMetricsChannel = () => {
  let socket = new Phoenix.Socket("ws://localhost:80/socket", {});
  socket.connect();

  let channel = socket.channel("metric:lobby", {});
  
  const [cpu, mem, disk] = [
    document.getElementById('cpu'),
    document.getElementById('mem'),
    document.getElementById('disk')
  ];

  channel.join()
    .receive("ok", resp => { console.log("Joined successfully", resp); })
    .receive("error", resp => { console.log("Unable to join", resp); });

  channel.on("metrics_data", payload => {
    const metric = payload.metrics[0];
    cpu.innerHTML = `CPU: ${metric.cpu_usage}`;
    mem.innerHTML = `Memory: ${metric.memory_usage}`;
    disk.innerHTML = `Disk: ${metric.disk_usage}`;
  });

  socket.onError(() => console.log("There was an error with the connection!"));
  socket.onClose(() => console.log("The connection dropped"));
};

// Export the function for testing purposes
module.exports = setupMetricsChannel;
