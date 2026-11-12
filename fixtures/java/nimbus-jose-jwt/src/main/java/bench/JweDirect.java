package bench;

import com.nimbusds.jose.EncryptionMethod;
import com.nimbusds.jose.JWEAlgorithm;
import com.nimbusds.jose.JWEHeader;
import com.nimbusds.jose.JWEObject;
import com.nimbusds.jose.Payload;
import com.nimbusds.jose.crypto.DirectDecrypter;
import com.nimbusds.jose.crypto.DirectEncrypter;
import javax.crypto.KeyGenerator;
import javax.crypto.SecretKey;

public class JweDirect {
    public static void main(String[] args) throws Exception {
        KeyGenerator keyGen = KeyGenerator.getInstance("AES");
        keyGen.init(256);
        SecretKey key = keyGen.generateKey();

        JWEObject jwe = new JWEObject(
                new JWEHeader(JWEAlgorithm.DIR, EncryptionMethod.A256GCM),
                new Payload("Hello, world!"));
        DirectEncrypter encrypter = new DirectEncrypter(key);
        jwe.encrypt(encrypter);
        String s = jwe.serialize();

        JWEObject parsed = JWEObject.parse(s);
        DirectDecrypter decrypter = new DirectDecrypter(key);
        parsed.decrypt(decrypter);
        System.out.println("dir " + parsed.getPayload());
    }
}
