package bench;

import java.math.BigInteger;
import java.nio.charset.StandardCharsets;
import java.security.SecureRandom;
import org.bouncycastle.asn1.x9.ECNamedCurveTable;
import org.bouncycastle.asn1.x9.X9ECParameters;
import org.bouncycastle.crypto.AsymmetricCipherKeyPair;
import org.bouncycastle.crypto.digests.SHA256Digest;
import org.bouncycastle.crypto.generators.ECKeyPairGenerator;
import org.bouncycastle.crypto.params.ECDomainParameters;
import org.bouncycastle.crypto.params.ECKeyGenerationParameters;
import org.bouncycastle.crypto.params.ParametersWithRandom;
import org.bouncycastle.crypto.signers.ECDSASigner;

public class EcdsaP256 {
    public static void main(String[] args) {
        SecureRandom random = new SecureRandom();
        X9ECParameters curve = ECNamedCurveTable.getByName("P-256");
        ECDomainParameters domain = new ECDomainParameters(curve.getCurve(), curve.getG(), curve.getN(), curve.getH());
        ECKeyPairGenerator generator = new ECKeyPairGenerator();
        generator.init(new ECKeyGenerationParameters(domain, random));
        AsymmetricCipherKeyPair keyPair = generator.generateKeyPair();

        byte[] message = "hello world".getBytes(StandardCharsets.UTF_8);
        SHA256Digest digest = new SHA256Digest();
        digest.update(message, 0, message.length);
        byte[] hash = new byte[digest.getDigestSize()];
        digest.doFinal(hash, 0);

        ECDSASigner signer = new ECDSASigner();
        signer.init(true, new ParametersWithRandom(keyPair.getPrivate(), random));
        BigInteger[] signature = signer.generateSignature(hash);

        ECDSASigner verifier = new ECDSASigner();
        verifier.init(false, keyPair.getPublic());
        boolean valid = verifier.verifySignature(hash, signature[0], signature[1]);

        System.out.println("valid: " + valid);
    }
}
