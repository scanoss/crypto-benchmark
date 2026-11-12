package bench;

import java.nio.charset.StandardCharsets;
import java.security.KeyPair;
import java.security.KeyPairGenerator;
import javax.crypto.Cipher;

public class RsaOaep {
    public static void main(String[] args) throws Exception {
        KeyPairGenerator keyPairGen = KeyPairGenerator.getInstance("RSA");
        keyPairGen.initialize(3072);
        KeyPair keyPair = keyPairGen.generateKeyPair();

        Cipher encryptor = Cipher.getInstance("RSA/ECB/OAEPWithSHA-256AndMGF1Padding");
        encryptor.init(Cipher.ENCRYPT_MODE, keyPair.getPublic());
        byte[] ciphertext = encryptor.doFinal("secret message".getBytes(StandardCharsets.UTF_8));

        Cipher decryptor = Cipher.getInstance("RSA/ECB/OAEPWithSHA-256AndMGF1Padding");
        decryptor.init(Cipher.DECRYPT_MODE, keyPair.getPrivate());
        byte[] decrypted = decryptor.doFinal(ciphertext);

        System.out.println(new String(decrypted, StandardCharsets.UTF_8));
    }
}
