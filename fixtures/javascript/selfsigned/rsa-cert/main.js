const { X509Certificate } = require('node:crypto');
const selfsigned = require('selfsigned');

async function main() {
  const attrs = [{ name: 'commonName', value: 'localhost' }];
  const pems = await selfsigned.generate(attrs, {
    keySize: 2048,
    algorithm: 'sha256',
    extensions: [{ name: 'subjectAltName', altNames: [{ type: 2, value: 'localhost' }] }],
  });
  const cert = new X509Certificate(pems.cert);
  console.log('ok', cert.subject, pems.private.startsWith('-----BEGIN PRIVATE KEY-----'));
}

main();
