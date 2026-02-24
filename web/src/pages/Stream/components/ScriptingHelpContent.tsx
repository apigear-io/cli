import { Text } from '@mantine/core';
import {
  HelpSection,
  HelpCode,
  HelpTable,
  HelpAlert,
  HelpList,
} from '@/components/HelpDrawer';

export const scriptingHelpTabs = [
  {
    value: 'overview',
    label: 'Overview',
    content: (
      <>
        <HelpSection title="What is Scripting?">
          <Text>
            Scripts are JavaScript programs that can act as WebSocket clients or backend servers.
            They run continuously until manually stopped or until they call <code>exit()</code>.
          </Text>
        </HelpSection>

        <HelpSection title="Script Types">
          <HelpTable
            headers={['Type', 'Purpose', 'Usage']}
            rows={[
              [
                <strong>Client</strong>,
                'Connect to backend servers',
                'Use connect() to connect to a WebSocket URL and interact with remote services',
              ],
              [
                <strong>Backend</strong>,
                'Provide services to clients',
                'Use createBackend() to start a server that clients can connect to',
              ],
            ]}
          />
        </HelpSection>

        <HelpSection title="Script Lifecycle">
          <HelpList
            items={[
              <>
                <strong>Run</strong>: Click Run button or use <code>POST /api/v1/stream/scripts/{'{name}'}/run</code>
              </>,
              <>
                <strong>Stop</strong>: Click Stop button or use <code>POST /api/v1/stream/scripts/stop/{'{id}'}</code>
              </>,
              <>
                <strong>Auto-stop</strong>: Call <code>exit()</code> in your script
              </>,
            ]}
          />
        </HelpSection>

        <HelpAlert>
          Scripts run forever by design. Use the Stop button or call exit() to terminate.
        </HelpAlert>
      </>
    ),
  },
  {
    value: 'client-api',
    label: 'Client API',
    content: (
      <>
        <HelpSection title="connect(url)">
          <Text>
            Creates a WebSocket connection and returns a client object. Connection happens
            asynchronously with automatic retry on failure.
          </Text>
          <HelpCode
            code={`const client = connect('ws://localhost:8080/ws');

client.onConnect(() => {
  console.log('Connected!');
  client.link('demo.Counter');
});

client.onError((error) => {
  console.error('Error:', error);
});`}
          />
        </HelpSection>

        <HelpSection title="Connection Events">
          <HelpTable
            headers={['Method', 'Description', 'Example']}
            rows={[
              [
                <code>onConnect(callback)</code>,
                'Called when WebSocket connects',
                <code>{'client.onConnect(() => {...})'}</code>,
              ],
              [
                <code>onDisconnect(callback)</code>,
                'Called when connection is lost',
                <code>{'client.onDisconnect(() => {...})'}</code>,
              ],
              [
                <code>onError(callback)</code>,
                'Called on connection errors',
                <code>{'client.onError((err) => {...})'}</code>,
              ],
            ]}
          />
        </HelpSection>

        <HelpSection title="ObjectLink Protocol Events">
          <HelpTable
            headers={['Method', 'Description', 'Example']}
            rows={[
              [
                <code>onInit(callback)</code>,
                'Receives INIT message with service info',
                <code>{'client.onInit((msg) => {...})'}</code>,
              ],
              [
                <code>onPropertyChange(callback)</code>,
                'Receives property change notifications',
                <code>{'client.onPropertyChange((change) => {...})'}</code>,
              ],
              [
                <code>onSignal(callback)</code>,
                'Receives signal notifications',
                <code>{'client.onSignal((signal) => {...})'}</code>,
              ],
            ]}
          />
        </HelpSection>

        <HelpSection title="ObjectLink Operations">
          <HelpTable
            headers={['Method', 'Description', 'Example']}
            rows={[
              [
                <code>link(objectId)</code>,
                'Subscribe to an object/interface',
                <code>{"client.link('demo.Counter')"}</code>,
              ],
              [
                <code>unlink(objectId)</code>,
                'Unsubscribe from an object',
                <code>{"client.unlink('demo.Counter')"}</code>,
              ],
              [
                <code>setProperty(propId, value)</code>,
                'Set a property value',
                <code>{"client.setProperty('count', 5)"}</code>,
              ],
              [
                <code>invoke(methodId, ...args)</code>,
                'Invoke a method (returns Promise)',
                <code>{"client.invoke('increment')"}</code>,
              ],
            ]}
          />
        </HelpSection>

        <HelpSection title="Interface Handles">
          <Text>
            For easier interaction, use <code>interface(objectId)</code> to get a handle:
          </Text>
          <HelpCode
            code={`const counter = client.interface('demo.Counter');

counter.link();

counter.onPropertyChange('count', (value) => {
  console.log('Count:', value);
});

counter.invoke('increment').then(() => {
  console.log('Incremented!');
});`}
          />
        </HelpSection>

        <HelpSection title="Other Methods">
          <HelpList
            items={[
              <>
                <code>client.send(message)</code> - Send raw WebSocket message
              </>,
              <>
                <code>client.onMessage(callback)</code> - Receive raw messages (bypasses ObjectLink
                processing)
              </>,
              <>
                <code>client.close()</code> - Close the connection
              </>,
            ]}
          />
        </HelpSection>
      </>
    ),
  },
  {
    value: 'backend-api',
    label: 'Backend API',
    content: (
      <>
        <HelpSection title="createBackend(url)">
          <Text>
            Creates a WebSocket backend server that clients can connect to. The server starts
            immediately and runs until the script stops.
          </Text>
          <HelpCode
            code={`const backend = createBackend('ws://localhost:8080/ws');

backend.register('demo.Counter', {
  count: 0,

  increment() {
    this.count++;
    backend.notifyPropertyChanged('count', this.count);
  }
});

console.log('Backend running...');`}
          />
        </HelpSection>

        <HelpSection title="Registering Objects">
          <Text>
            Use <code>backend.register(objectId, implementation)</code> to provide an object
            implementation:
          </Text>
          <HelpCode
            code={`backend.register('demo.Calculator', {
  result: 0,

  add(a, b) {
    this.result = a + b;
    backend.notifyPropertyChanged('result', this.result);
    return this.result;
  },

  clear() {
    this.result = 0;
    backend.notifyPropertyChanged('result', this.result);
  }
});`}
          />
        </HelpSection>

        <HelpSection title="Sending Notifications">
          <HelpTable
            headers={['Method', 'Description', 'Example']}
            rows={[
              [
                <code>notifyPropertyChanged(propId, value)</code>,
                'Notify clients of property change',
                <code>{"backend.notifyPropertyChanged('count', 5)"}</code>,
              ],
              [
                <code>notifySignal(signalId, args)</code>,
                'Send a signal to clients',
                <code>{"backend.notifySignal('alarm', [])"}</code>,
              ],
            ]}
          />
        </HelpSection>

        <HelpSection title="Backend Lifecycle">
          <HelpList
            items={[
              'Backend starts automatically when script runs',
              'Accepts connections from multiple clients',
              'Runs until script is stopped or calls exit()',
              'Automatically handles client connect/disconnect',
            ]}
          />
        </HelpSection>

        <HelpAlert>
          The backend server runs on the URL specified in createBackend(). Make sure the port is
          not already in use.
        </HelpAlert>
      </>
    ),
  },
  {
    value: 'utilities',
    label: 'Utilities',
    content: (
      <>
        <HelpSection title="Global Functions">
          <HelpTable
            headers={['Function', 'Description', 'Example']}
            rows={[
              [
                <code>console.log(...args)</code>,
                'Log to console output',
                <code>{"console.log('Hello')"}</code>,
              ],
              [
                <code>console.error(...args)</code>,
                'Log error to console',
                <code>{"console.error('Failed')"}</code>,
              ],
              [
                <code>console.warn(...args)</code>,
                'Log warning to console',
                <code>{"console.warn('Warning')"}</code>,
              ],
              [
                <code>print(...args)</code>,
                'Alias for console.log',
                <code>{"print('Hello')"}</code>,
              ],
              [
                <code>exit()</code>,
                'Stop the script',
                <code>exit()</code>,
              ],
            ]}
          />
        </HelpSection>

        <HelpSection title="Timing Functions">
          <HelpTable
            headers={['Function', 'Description', 'Example']}
            rows={[
              [
                <code>after(ms, callback)</code>,
                'Execute callback after delay (one-time)',
                <code>{'after(1000, () => {...})'}</code>,
              ],
              [
                <code>every(ms, callback)</code>,
                'Execute callback repeatedly',
                <code>{'every(1000, () => {...})'}</code>,
              ],
            ]}
          />
          <HelpCode
            code={`// Run once after 2 seconds
after(2000, () => {
  console.log('Delayed message');
});

// Run every second
every(1000, () => {
  console.log('Tick:', new Date().toISOString());
});`}
          />
        </HelpSection>

        <HelpSection title="Faker - Random Data Generation">
          <Text>
            The <code>faker</code> object provides methods for generating random test data:
          </Text>
          <HelpCode
            code={`// Generate random data
console.log(faker.person.fullName());
console.log(faker.internet.email());
console.log(faker.number.int({ min: 1, max: 100 }));
console.log(faker.lorem.sentence());
console.log(faker.date.recent());`}
          />
          <HelpAlert>
            Faker includes many categories: person, internet, number, lorem, date, address, phone,
            commerce, and more.
          </HelpAlert>
        </HelpSection>

        <HelpSection title="Trace Files">
          <Text>
            Read and replay trace files captured from WebSocket connections:
          </Text>
          <HelpCode
            code={`// Read trace file
const trace = readTrace('my-trace.jsonl');

console.log('Messages:', trace.messages.length);
trace.messages.forEach((msg, i) => {
  console.log(\`Message \${i}:\`, msg);
});`}
          />
        </HelpSection>
      </>
    ),
  },
  {
    value: 'examples',
    label: 'Examples',
    content: (
      <>
        <HelpSection title="Simple Client Script">
          <HelpCode
            code={`// Connect to a backend
const client = connect('ws://localhost:8080/ws');

client.onConnect(() => {
  console.log('✓ Connected');

  // Link to Counter interface
  const counter = client.interface('demo.Counter');
  counter.link();

  // Watch for changes
  counter.onPropertyChange('count', (value) => {
    console.log('Count changed to:', value);
  });

  // Increment every 2 seconds
  every(2000, () => {
    counter.invoke('increment');
  });
});

client.onError((err) => {
  console.error('Connection error:', err);
});`}
          />
        </HelpSection>

        <HelpSection title="Simple Backend Script">
          <HelpCode
            code={`// Create backend server
const backend = createBackend('ws://localhost:8080/ws');

// Register Counter service
backend.register('demo.Counter', {
  count: 0,

  increment() {
    this.count++;
    console.log('Count:', this.count);
    backend.notifyPropertyChanged('count', this.count);
  },

  decrement() {
    this.count--;
    console.log('Count:', this.count);
    backend.notifyPropertyChanged('count', this.count);
  },

  reset() {
    this.count = 0;
    backend.notifyPropertyChanged('count', this.count);
    backend.notifySignal('onReset', []);
  }
});

console.log('Backend running on ws://localhost:8080/ws');
console.log('Registered: demo.Counter');`}
          />
        </HelpSection>

        <HelpSection title="Test Data Generator">
          <HelpCode
            code={`// Generate test data periodically
const backend = createBackend('ws://localhost:8080/ws');

backend.register('demo.User', {
  name: '',
  email: '',

  generateRandom() {
    this.name = faker.person.fullName();
    this.email = faker.internet.email();

    backend.notifyPropertyChanged('name', this.name);
    backend.notifyPropertyChanged('email', this.email);

    console.log('Generated:', this.name, this.email);
  }
});

// Generate new user every 3 seconds
every(3000, () => {
  backend.invoke('demo.User', 'generateRandom', []);
});

console.log('Test data generator running...');`}
          />
        </HelpSection>

        <HelpSection title="Self-Terminating Script">
          <HelpCode
            code={`// Script that runs for a set time then exits
const client = connect('ws://localhost:8080/ws');

let messageCount = 0;

client.onConnect(() => {
  console.log('Connected');

  client.onMessage((msg) => {
    console.log('Message:', msg);
    messageCount++;

    // Stop after 10 messages
    if (messageCount >= 10) {
      console.log('Received 10 messages, stopping...');
      client.close();
      exit();
    }
  });
});

// Or stop after 30 seconds
after(30000, () => {
  console.log('Timeout reached, stopping...');
  client.close();
  exit();
});`}
          />
        </HelpSection>
      </>
    ),
  },
];
