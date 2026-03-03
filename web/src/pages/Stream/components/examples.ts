export interface ScriptExample {
  name: string;
  description: string;
  code: string;
}

export const EXAMPLES: ScriptExample[] = [
  {
    name: 'Simple Client',
    description: 'Connect to an ObjectLink backend and log when connected',
    code: `// Connect to ObjectLink backend
const client = connect('ws://localhost:5560/ws');

client.onConnect(() => {
  console.log('Connected to backend!');
});

client.onDisconnect(() => {
  console.log('Disconnected from backend');
});

client.onError((error) => {
  console.error('Connection error:', error);
});`,
  },
  {
    name: 'Echo Backend',
    description: 'Create a simple echo server that responds to method calls',
    code: `// Create an echo server
const backend = createBackend('ws://0.0.0.0:5560/ws');

backend.register('demo.Echo', {
  methods: {
    echo: (params, ctx) => {
      console.log('Received echo request:', params.message);
      return { result: params.message };
    },
    uppercase: (params, ctx) => {
      const result = params.text.toUpperCase();
      console.log(\`Uppercase: "\${params.text}" -> "\${result}"\`);
      return { result };
    },
  },
});

console.log('Echo backend listening on ws://0.0.0.0:5560/ws');
console.log('Available methods: echo, uppercase');`,
  },
  {
    name: 'Counter Backend',
    description: 'Backend with properties, methods, and lifecycle callbacks',
    code: `// Create a counter backend with state
const backend = createBackend('ws://0.0.0.0:5560/ws');

backend.register('demo.Counter', {
  properties: {
    count: 0
  },
  methods: {
    increment(params, ctx) {
      let count = ctx.get('count');
      count++;
      ctx.set('count', count);  // Broadcasts PROPERTY_CHANGE automatically
      console.log('Count incremented to', count);

      // Emit signal when count reaches 10
      if (count === 10) {
        ctx.emit('milestone', { value: 10 });
      }

      return count;
    },

    reset(params, ctx) {
      ctx.set('count', 0);
      console.log('Count reset');
      return 0;
    },

    getStatus(params, ctx) {
      const count = ctx.get('count');
      const clients = ctx.clientCount();
      return {
        count: count,
        clients: clients,
        status: count > 0 ? 'active' : 'idle'
      };
    }
  },
  onLink(ctx) {
    console.log('Client linked! Total clients:', ctx.clientCount());
  },
  onUnlink(ctx) {
    console.log('Client unlinked. Remaining clients:', ctx.clientCount());
  }
});

console.log('Counter backend running on ws://0.0.0.0:5560/ws');`,
  },
  {
    name: 'Backend with Auto-Update',
    description: 'Backend that automatically updates properties using object handle',
    code: `// Create backend with auto-updating temperature sensor
const backend = createBackend('ws://0.0.0.0:5560/ws');

const sensor = backend.register('demo.TempSensor', {
  properties: {
    temperature: 20.0,
    unit: 'celsius',
    lastUpdate: new Date().toISOString()
  },
  methods: {
    setUnit(params, ctx) {
      ctx.set('unit', params.unit);
      console.log('Unit changed to', params.unit);
    },

    getReading(params, ctx) {
      return {
        temperature: ctx.get('temperature'),
        unit: ctx.get('unit'),
        lastUpdate: ctx.get('lastUpdate')
      };
    }
  }
});

// Auto-update temperature every 3 seconds
every(3000, () => {
  // Generate random temperature between 18-25°C
  const temp = 18 + Math.random() * 7;
  const rounded = Math.round(temp * 10) / 10;

  // Update properties using object handle
  sensor.set('temperature', rounded);
  sensor.set('lastUpdate', new Date().toISOString());

  console.log(\`Temperature updated: \${rounded}°C\`);

  // Emit warning if temperature is too high
  if (rounded > 24) {
    sensor.emit('warning', { message: 'Temperature too high!' });
  }
});

console.log('Temperature sensor running on ws://0.0.0.0:5560/ws');`,
  },
  {
    name: 'Faker Data Generator',
    description: 'Generate fake data every 2 seconds using the faker library',
    code: `// Generate fake data every 2 seconds
console.log('Starting fake data generator...');

every(2000, () => {
  const person = {
    name: faker.name(),
    email: faker.email(),
    company: faker.company(),
    address: faker.address(),
    phone: faker.phone(),
  };

  console.log('Generated person:', JSON.stringify(person, null, 2));
});

console.log('Generator started. Will output data every 2 seconds.');`,
  },
  {
    name: 'Timer Demo',
    description: 'Demonstrate setTimeout and setInterval',
    code: `// Timer demonstration
console.log('Starting timer demo...');

// One-time delayed execution
setTimeout(() => {
  console.log('This message appears after 3 seconds');
}, 3000);

// Repeating execution
let count = 0;
const intervalId = setInterval(() => {
  count++;
  console.log(\`Interval tick #\${count}\`);

  if (count >= 5) {
    clearInterval(intervalId);
    console.log('Interval stopped after 5 ticks');
  }
}, 1000);

console.log('Timers started');`,
  },
  {
    name: 'Console API Demo',
    description: 'Demonstrate different console logging levels',
    code: `// Console API demonstration
console.log('This is a regular log message');
console.info('This is an info message');
console.warn('This is a warning message');
console.error('This is an error message');
console.debug('This is a debug message');

// Logging objects
const user = {
  id: 123,
  name: 'John Doe',
  role: 'developer',
};

console.log('User object:', user);
console.log('JSON formatted:', JSON.stringify(user, null, 2));`,
  },
  {
    name: 'ObjectLink Client with Calls',
    description: 'Connect to backend and make method calls',
    code: `// Connect and make method calls
const client = connect('ws://localhost:5560/ws');

client.onConnect(() => {
  console.log('Connected to backend!');

  // Link to the object first
  client.link('demo.Echo');

  // Call a remote method
  client.invoke('demo.Echo', 'echo', { message: 'Hello from script!' })
    .then(result => {
      console.log('Echo response:', result);

      // Call another method
      return client.invoke('demo.Echo', 'uppercase', { text: 'hello world' });
    })
    .then(result => {
      console.log('Uppercase response:', result);
    })
    .catch(error => {
      console.error('Call failed:', error);
    });
});`,
  },
];
