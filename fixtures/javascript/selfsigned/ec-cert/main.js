const selfsigned = require('selfsigned');

async function main() {
  const pems = await selfsigned.generate([{ name: 'commonName', value: 'localhost' }], {
    keyType: 'ec',
    curve: 'P-256',
    algorithm: 'sha256',
  });
  console.log('ok', pems.cert.startsWith('-----BEGIN CERTIFICATE-----'));
}

main();
