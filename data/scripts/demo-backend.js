// Create an echo server
const backend = createBackend('ws://0.0.0.0:5560/ws');

backend.register('demo.Counter', {
  count: 0,
  increment() {
    this.count++;
    backend.notifyPropertyChanged('count', this.count);
  }
});
