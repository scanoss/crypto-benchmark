import 'reflect-metadata';
import * as x509 from '@peculiar/x509';

async function main() {
  const algorithm = { name: 'ECDSA', namedCurve: 'P-256', hash: 'SHA-256' };
  const keys = await crypto.subtle.generateKey(algorithm, false, ['sign', 'verify']);
  const cert = await x509.X509CertificateGenerator.createSelfSigned({
    serialNumber: '01',
    name: 'CN=Example',
    notBefore: new Date(),
    notAfter: new Date(Date.now() + 365 * 24 * 60 * 60 * 1000),
    keys,
    signingAlgorithm: algorithm,
    extensions: [new x509.BasicConstraintsExtension(true, undefined, true)],
  });
  const pem = cert.toString('pem');
  console.log('ok', pem.startsWith('-----BEGIN CERTIFICATE-----'));
}

await main();
