// Connect to ObjectLink backend
const client = connect('ws://localhost:5550/ws');

client.onConnect(() => {
  console.log('connected to proxy!');
  console.log('available services:', client.services());
});

client.onError((error) => {
  console.error('connection error:', error);
});