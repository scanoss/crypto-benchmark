import cryptoRandomString from 'crypto-random-string';

function main() {
  const sessionId = cryptoRandomString({ length: 32 });
  const resetToken = cryptoRandomString({ length: 24, type: 'url-safe' });
  const inviteCode = cryptoRandomString({ length: 10, type: 'distinguishable' });
  console.log('ok', sessionId.length, resetToken.length, inviteCode.length);
}

main();
