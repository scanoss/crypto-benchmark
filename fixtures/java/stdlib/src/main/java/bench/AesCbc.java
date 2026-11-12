package bench;

import java.nio.charset.StandardCharsets;
import java.security.SecureRandom;
import javax.crypto.Cipher;
import javax.crypto.KeyGenerator;
import javax.crypto.SecretKey;
import javax.crypto.spec.IvParameterSpec;

public class AesCbc {
    public static void main(String[] args) throws Exception {
        KeyGenerator keyGen = KeyGenerator.getInstance("AES");
        keyGen.init(256);
        SecretKey key = keyGen.generateKey();

        byte[] iv = new byte[16];
        new SecureRandom().nextBytes(iv);
        IvParameterSpec ivSpec = new IvParameterSpec(iv);

        Cipher encryptor = Cipher.getInstance("AES/CBC/PKCS5Padding");
        encryptor.init(Cipher.ENCRYPT_MODE, key, ivSpec);
        byte[] ciphertext = encryptor.doFinal("secret message".getBytes(StandardCharsets.UTF_8));

        Cipher decryptor = Cipher.getInstance("AES/CBC/PKCS5Padding");
        decryptor.init(Cipher.DECRYPT_MODE, key, ivSpec);
        byte[] decrypted = decryptor.doFinal(ciphertext);

        System.out.println(new String(decrypted, StandardCharsets.UTF_8));
    }
}
