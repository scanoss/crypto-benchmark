import { calculateJwkThumbprint, exportJWK, exportPKCS8, generateKeyPair, importJWK, importPKCS8 } from 'jose';

async function main() {
  const { publicKey, privateKey } = await generateKeyPair('PS256', { modulusLength: 2048, extractable: true });
  const pkcs8 = await exportPKCS8(privateKey);
  const imported = await importPKCS8(pkcs8, 'PS256');

  const publicJwk = await exportJWK(publicKey);
  const restored = await importJWK(publicJwk, 'PS256');
  const thumbprint = await calculateJwkThumbprint(publicJwk, 'sha256');
  console.log('ok', imported.type, restored.type, thumbprint.length);
}

await main();
