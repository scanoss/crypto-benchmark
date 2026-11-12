package bench;

import java.security.Key;
import java.security.SecureRandom;
import org.jose4j.jwa.AlgorithmConstraints;
import org.jose4j.jwa.AlgorithmConstraints.ConstraintType;
import org.jose4j.jws.AlgorithmIdentifiers;
import org.jose4j.jws.JsonWebSignature;
import org.jose4j.keys.HmacKey;

public class JwsHs256 {
    public static void main(String[] args) throws Exception {
        byte[] secret = new byte[32];
        new SecureRandom().nextBytes(secret);
        Key key = new HmacKey(secret);

        JsonWebSignature jws = new JsonWebSignature();
        jws.setPayload("{\"sub\":\"bench\"}");
        jws.setAlgorithmHeaderValue(AlgorithmIdentifiers.HMAC_SHA256);
        jws.setKey(key);
        String compactSerialization = jws.getCompactSerialization();

        JsonWebSignature received = new JsonWebSignature();
        received.setAlgorithmConstraints(new AlgorithmConstraints(ConstraintType.PERMIT, AlgorithmIdentifiers.HMAC_SHA256));
        received.setCompactSerialization(compactSerialization);
        received.setKey(key);
        boolean signatureVerified = received.verifySignature();
        System.out.println("hs256 " + signatureVerified + " " + received.getPayload());
    }
}
