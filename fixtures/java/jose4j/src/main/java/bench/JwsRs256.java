package bench;

import org.jose4j.jwa.AlgorithmConstraints;
import org.jose4j.jwa.AlgorithmConstraints.ConstraintType;
import org.jose4j.jwk.RsaJsonWebKey;
import org.jose4j.jwk.RsaJwkGenerator;
import org.jose4j.jws.AlgorithmIdentifiers;
import org.jose4j.jws.JsonWebSignature;

public class JwsRs256 {
    public static void main(String[] args) throws Exception {
        RsaJsonWebKey rsaJsonWebKey = RsaJwkGenerator.generateJwk(3072);
        rsaJsonWebKey.setKeyId("k1");

        JsonWebSignature jws = new JsonWebSignature();
        jws.setPayload("This is some text that is to be signed.");
        jws.setAlgorithmHeaderValue(AlgorithmIdentifiers.RSA_USING_SHA256);
        jws.setKey(rsaJsonWebKey.getPrivateKey());
        jws.setKeyIdHeaderValue(rsaJsonWebKey.getKeyId());
        String compactSerialization = jws.getCompactSerialization();

        JsonWebSignature received = new JsonWebSignature();
        received.setAlgorithmConstraints(new AlgorithmConstraints(ConstraintType.PERMIT, AlgorithmIdentifiers.RSA_USING_SHA256));
        received.setCompactSerialization(compactSerialization);
        received.setKey(rsaJsonWebKey.getPublicKey());
        boolean signatureVerified = received.verifySignature();
        System.out.println("rs256 " + signatureVerified + " " + received.getPayload());
    }
}
