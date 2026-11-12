import { X509Certificate } from 'node:crypto';
import { rootCertificates } from 'node:tls';

function main() {
  const cert = new X509Certificate(rootCertificates[0]);
  const fingerprint = cert.fingerprint256;
  console.log('ok', cert.subject.split('\n')[0], fingerprint.length);
}

main();
