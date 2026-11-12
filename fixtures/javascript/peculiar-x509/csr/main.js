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
  const csr = await x509.Pkcs10CertificateRequestGenerator.create({
    name: 'CN=service.example.org',
    keys,
    signingAlgorithm: algorithm,
    extensions: [new x509.KeyUsagesExtension(x509.KeyUsageFlags.digitalSignature | x509.KeyUsageFlags.keyEncipherment)],
  });
  const pem = csr.toString('pem');
  console.log('ok', pem.startsWith('-----BEGIN CERTIFICATE REQUEST-----'));
}

await main();
