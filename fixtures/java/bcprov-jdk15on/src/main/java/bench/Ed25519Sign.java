package bench;

import java.nio.charset.StandardCharsets;
import java.security.SecureRandom;
import org.bouncycastle.crypto.AsymmetricCipherKeyPair;
import org.bouncycastle.crypto.generators.Ed25519KeyPairGenerator;
import org.bouncycastle.crypto.params.Ed25519KeyGenerationParameters;
import org.bouncycastle.crypto.signers.Ed25519Signer;

public class Ed25519Sign {
    public static void main(String[] args) {
        Ed25519KeyPairGenerator generator = new Ed25519KeyPairGenerator();
        generator.init(new Ed25519KeyGenerationParameters(new SecureRandom()));
        AsymmetricCipherKeyPair keyPair = generator.generateKeyPair();
        byte[] message = "hello world".getBytes(StandardCharsets.UTF_8);

        Ed25519Signer signer = new Ed25519Signer();
        signer.init(true, keyPair.getPrivate());
        signer.update(message, 0, message.length);
        byte[] signature = signer.generateSignature();

        Ed25519Signer verifier = new Ed25519Signer();
        verifier.init(false, keyPair.getPublic());
        verifier.update(message, 0, message.length);
        boolean valid = verifier.verifySignature(signature);

        System.out.println("valid: " + valid);
    }
}
