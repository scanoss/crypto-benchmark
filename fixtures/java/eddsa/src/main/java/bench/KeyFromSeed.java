package bench;

import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.security.SecureRandom;
import java.security.Signature;
import net.i2p.crypto.eddsa.EdDSAEngine;
import net.i2p.crypto.eddsa.EdDSAPrivateKey;
import net.i2p.crypto.eddsa.EdDSAPublicKey;
import net.i2p.crypto.eddsa.spec.EdDSANamedCurveTable;
import net.i2p.crypto.eddsa.spec.EdDSAParameterSpec;
import net.i2p.crypto.eddsa.spec.EdDSAPrivateKeySpec;
import net.i2p.crypto.eddsa.spec.EdDSAPublicKeySpec;

public class KeyFromSeed {
    public static void main(String[] args) throws Exception {
        EdDSAParameterSpec curve = EdDSANamedCurveTable.getByName(EdDSANamedCurveTable.ED_25519);

        byte[] seed = new byte[32];
        new SecureRandom().nextBytes(seed);

        EdDSAPrivateKeySpec privateSpec = new EdDSAPrivateKeySpec(seed, curve);
        EdDSAPrivateKey privateKey = new EdDSAPrivateKey(privateSpec);
        EdDSAPublicKey publicKey = new EdDSAPublicKey(new EdDSAPublicKeySpec(privateSpec.getA(), curve));

        byte[] message = "message to sign".getBytes(StandardCharsets.UTF_8);

        Signature signer = new EdDSAEngine(MessageDigest.getInstance(curve.getHashAlgorithm()));
        signer.initSign(privateKey);
        signer.update(message);
        byte[] signature = signer.sign();

        Signature verifier = new EdDSAEngine(MessageDigest.getInstance(curve.getHashAlgorithm()));
        verifier.initVerify(publicKey);
        verifier.update(message);
        System.out.println("valid: " + verifier.verify(signature));
    }
}
