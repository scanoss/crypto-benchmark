const express = require('express');
const liba = require('cbench-handler-express-liba');

function health(req, res) {
  res.send('ok');
}

function fingerprint(req, res) {
  res.send(liba.fingerprint(Buffer.from(String(req.query.data ?? ''))).toString('hex'));
}

function main() {
  const app = express();
  app.get('/health', health);
  app.get('/fingerprint', fingerprint);
  if (process.argv[2] === 'serve') {
    app.listen(8080, '127.0.0.1');
    return;
  }
  // Without "serve", answer one request on an ephemeral port and exit.
  const server = app.listen(0, '127.0.0.1', async () => {
    const response = await fetch(`http://127.0.0.1:${server.address().port}/fingerprint?data=hello`);
    console.log(response.status, await response.text());
    server.close();
  });
}

main();
