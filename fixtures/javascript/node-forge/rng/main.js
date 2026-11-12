const forge = require('node-forge');

function main() {
  const token = forge.random.getBytesSync(32);
  forge.random.getBytes(16, (err, bytes) => {
    if (err) throw err;
    console.log('ok', forge.util.bytesToHex(token).length, bytes.length);
  });
}

main();
