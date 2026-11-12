package bench;

import java.nio.charset.StandardCharsets;
import java.security.KeyPair;
import java.security.MessageDigest;
import java.security.Signature;
import net.i2p.crypto.eddsa.EdDSAEngine;
import net.i2p.crypto.eddsa.KeyPairGenerator;

public class SignVerify {
    public static void main(String[] args) throws Exception {
        KeyPairGenerator keyPairGen = new KeyPairGenerator();
        KeyPair keyPair = keyPairGen.generateKeyPair();

        byte[] message = "message to sign".getBytes(StandardCharsets.UTF_8);

        Signature signer = new EdDSAEngine(MessageDigest.getInstance("SHA-512"));
        signer.initSign(keyPair.getPrivate());
        signer.update(message);
        byte[] signature = signer.sign();

        Signature verifier = new EdDSAEngine(MessageDigest.getInstance("SHA-512"));
        verifier.initVerify(keyPair.getPublic());
        verifier.update(message);
        System.out.println("valid: " + verifier.verify(signature));
    }
}
