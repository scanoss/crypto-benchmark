package bench;

import java.nio.charset.StandardCharsets;
import java.security.SecureRandom;
import javax.crypto.Cipher;
import javax.crypto.KeyGenerator;
import javax.crypto.SecretKey;
import javax.crypto.spec.IvParameterSpec;

public class ChaCha20Poly1305 {
    public static void main(String[] args) throws Exception {
        KeyGenerator keyGen = KeyGenerator.getInstance("ChaCha20");
        keyGen.init(256);
        SecretKey key = keyGen.generateKey();

        byte[] nonce = new byte[12];
        new SecureRandom().nextBytes(nonce);

        Cipher encryptor = Cipher.getInstance("ChaCha20-Poly1305");
        encryptor.init(Cipher.ENCRYPT_MODE, key, new IvParameterSpec(nonce));
        byte[] ciphertext = encryptor.doFinal("secret message".getBytes(StandardCharsets.UTF_8));

        Cipher decryptor = Cipher.getInstance("ChaCha20-Poly1305");
        decryptor.init(Cipher.DECRYPT_MODE, key, new IvParameterSpec(nonce));
        byte[] decrypted = decryptor.doFinal(ciphertext);

        System.out.println(new String(decrypted, StandardCharsets.UTF_8));
    }
}
