// test_require.js
const helper = require('./helper');

function main() {
  const message = helper.greet("World");
  console.log(message);
  return message;
}