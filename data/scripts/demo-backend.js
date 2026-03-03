// Create an echo server
const backend = createBackend('ws://0.0.0.0:5560/ws');

backend.register('demo.Counter', {
  properties: {
    count: 0
  },
  methods: {
    increment(params, ctx) {
      const count = ctx.get("count")
      count++
      ctx.set("count", count)
      return count
    }
  },
});
