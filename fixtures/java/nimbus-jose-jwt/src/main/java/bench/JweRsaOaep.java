package bench;

import com.nimbusds.jose.EncryptionMethod;
import com.nimbusds.jose.JWEAlgorithm;
import com.nimbusds.jose.JWEHeader;
import com.nimbusds.jose.JWEObject;
import com.nimbusds.jose.Payload;
import com.nimbusds.jose.crypto.RSADecrypter;
import com.nimbusds.jose.crypto.RSAEncrypter;
import com.nimbusds.jose.jwk.RSAKey;
import com.nimbusds.jose.jwk.gen.RSAKeyGenerator;

public class JweRsaOaep {
    public static void main(String[] args) throws Exception {
        RSAKey rsaJWK = new RSAKeyGenerator(3072)
                .keyID("456")
                .generate();

        JWEHeader header = new JWEHeader(JWEAlgorithm.RSA_OAEP_256, EncryptionMethod.A256GCM);
        JWEObject jwe = new JWEObject(header, new Payload("Hello, world!"));
        RSAEncrypter encrypter = new RSAEncrypter(rsaJWK.toRSAPublicKey());
        jwe.encrypt(encrypter);
        String s = jwe.serialize();

        JWEObject parsed = JWEObject.parse(s);
        RSADecrypter decrypter = new RSADecrypter(rsaJWK);
        parsed.decrypt(decrypter);
        System.out.println("rsa-oaep-256 " + parsed.getPayload());
    }
}
