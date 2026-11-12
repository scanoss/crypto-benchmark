const forge = require('node-forge');

function main() {
  let clientHello = '';
  const client = forge.tls.createConnection({
    server: false,
    verify: (connection, verified) => verified,
    connected: () => {},
    tlsDataReady: (connection) => {
      clientHello += connection.tlsData.getBytes();
    },
    dataReady: () => {},
    closed: () => {},
    error: (connection, error) => {
      throw new Error(error.message);
    },
  });
  client.handshake();
  console.log('ok', clientHello.length > 0);
}

main();
