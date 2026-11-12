package bench;

import org.jose4j.jwa.AlgorithmConstraints;
import org.jose4j.jwa.AlgorithmConstraints.ConstraintType;
import org.jose4j.jwe.ContentEncryptionAlgorithmIdentifiers;
import org.jose4j.jwe.JsonWebEncryption;
import org.jose4j.jwe.KeyManagementAlgorithmIdentifiers;
import org.jose4j.jwk.RsaJsonWebKey;
import org.jose4j.jwk.RsaJwkGenerator;

public class JweRsaOaep {
    public static void main(String[] args) throws Exception {
        RsaJsonWebKey rsaJsonWebKey = RsaJwkGenerator.generateJwk(3072);

        JsonWebEncryption jwe = new JsonWebEncryption();
        jwe.setPlaintext("Well, as of this moment, they're on DOUBLE SECRET PROBATION!");
        jwe.setAlgorithmHeaderValue(KeyManagementAlgorithmIdentifiers.RSA_OAEP_256);
        jwe.setEncryptionMethodHeaderParameter(ContentEncryptionAlgorithmIdentifiers.AES_256_GCM);
        jwe.setKey(rsaJsonWebKey.getPublicKey());
        String serializedJwe = jwe.getCompactSerialization();

        JsonWebEncryption received = new JsonWebEncryption();
        received.setAlgorithmConstraints(new AlgorithmConstraints(ConstraintType.PERMIT, KeyManagementAlgorithmIdentifiers.RSA_OAEP_256));
        received.setContentEncryptionAlgorithmConstraints(new AlgorithmConstraints(ConstraintType.PERMIT, ContentEncryptionAlgorithmIdentifiers.AES_256_GCM));
        received.setCompactSerialization(serializedJwe);
        received.setKey(rsaJsonWebKey.getPrivateKey());
        String plaintext = received.getPlaintextString();
        System.out.println("rsa-oaep-256 " + plaintext);
    }
}
