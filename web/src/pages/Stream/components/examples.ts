export interface ScriptExample {
  name: string;
  description: string;
  code: string;
}

export const EXAMPLES: ScriptExample[] = [
  {
    name: 'Simple Client',
    description: 'Connect to an ObjectLink backend and log when ready',
    code: `// Connect to ObjectLink backend
const client = connect('ws://localhost:5560/ws');

client.onReady(() => {
  console.log('Connected to backend!');
  console.log('Available services:', client.services());
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

backend.defineObject('demo.Echo', {
  methods: {
    echo: (params) => {
      console.log('Received echo request:', params.message);
      return { result: params.message };
    },
    uppercase: (params) => {
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

client.onReady(async () => {
  console.log('Connected! Available services:', client.services());

  try {
    // Call a remote method
    const result = await client.call('demo.Echo', 'echo', {
      message: 'Hello from script!'
    });

    console.log('Echo response:', result);

    // Call another method
    const upper = await client.call('demo.Echo', 'uppercase', {
      text: 'hello world'
    });

    console.log('Uppercase response:', upper);
  } catch (error) {
    console.error('Call failed:', error);
  }
});`,
  },
];
