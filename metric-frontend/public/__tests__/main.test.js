const { JSDOM } =  require("jsdom")
const dom = new JSDOM()
global.document = dom.window.document
global.window = dom.window
const setupMetricsChannel = require('../main.js')

// Mock the global Phoenix object
global.Phoenix = {
  Socket: jest.fn().mockImplementation(() => {
    return {
      connect: jest.fn(),
      onError: jest.fn(),
      onClose: jest.fn(),
      channel: jest.fn().mockImplementation(() => {
        return {
          join: jest.fn().mockReturnValue({
            receive: jest.fn((status, callback) => {
              if (status === 'ok') callback();
              return { receive: jest.fn() }; // Ensure the chained receive method
            })
          }),
          on: jest.fn(),
          trigger: jest.fn()
        };
      })
    };
  })
};

beforeEach(() => {
  // Set up the DOM elements
  document.body.innerHTML = `
    <body>
      <h1>Live metrics</h1>
      <p>This project implements sockets to perform realtime data update on this very simple dashboard</p>
      <div id="metrics">
          <p id="cpu">CPU: []</p>
          <p id="mem">Memory: []</p>
          <p id="disk">diskio: []</p>
      </div>
    </body>
  `;
});

afterEach(() => {
  // Clean up the DOM elements
  document.body.innerHTML = '';
});

test('updates innerHTML on new metrics_data message', async () => {
  setupMetricsChannel(); // Run the setup function from main.js

  const cpu = document.getElementById('cpu');
  const mem = document.getElementById('mem');
  const disk = document.getElementById('disk');


  expect(global.Phoenix.Socket).toHaveBeenCalled(); // Ensure the Socket constructor was called
  const mockSocketInstance = global.Phoenix.Socket.mock.instances[0];
  expect(mockSocketInstance).toBeDefined(); // Ensure the instance is defined
  if (mockSocketInstance && mockSocketInstance.channel ) {

    expect(mockSocketInstance.channel).toHaveBeenCalled();
    const mockChannel = mockSocketInstance.channel.mock.results[0].value;

    await mockChannel.join().receive('ok', () => {});

    const metricsDataCallback = mockChannel.on.mock.calls.find(call => call[0] === 'metrics_data')[1];
    metricsDataCallback({
      metrics: [{
        cpu_usage: '20%',
        memory_usage: '30%',
        disk_usage: '40%'
      }]
    });

    // Check if the innerHTML is updated correctly
    expect(cpu.innerHTML).toBe('CPU: 20%');
    expect(mem.innerHTML).toBe('Memory: 30%');
    expect(disk.innerHTML).toBe('Disk: 40%');
  }
});
