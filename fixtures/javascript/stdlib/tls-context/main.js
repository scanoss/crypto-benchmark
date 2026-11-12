import tls from 'node:tls';

function main() {
  const client = tls.createSecureContext({ minVersion: 'TLSv1.2' });
  const strict = tls.createSecureContext({
    minVersion: 'TLSv1.3',
    ciphersuites: 'TLS_AES_256_GCM_SHA384:TLS_CHACHA20_POLY1305_SHA256',
  });
  console.log('ok', client.context !== undefined && strict.context !== undefined);
}

main();
