import 'reflect-metadata';
import * as x509 from '@peculiar/x509';

async function main() {
  const algorithm = {
    name: 'RSASSA-PKCS1-v1_5',
    hash: 'SHA-256',
    publicExponent: new Uint8Array([1, 0, 1]),
    modulusLength: 2048,
  };
  const keys = await crypto.subtle.generateKey(algorithm, false, ['sign', 'verify']);
  const issued = await x509.X509CertificateGenerator.createSelfSigned({
    serialNumber: '02',
    name: 'CN=Benchmark Root',
    keys,
    signingAlgorithm: algorithm,
  });
  const pem = issued.toString('pem');

  const cert = new x509.X509Certificate(pem);
  const valid = await cert.verify();
  const thumbprint = await cert.getThumbprint('SHA-256');
  console.log('ok', cert.subject, valid, thumbprint.byteLength);
}

await main();
