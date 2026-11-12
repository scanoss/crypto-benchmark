package bench;

import java.nio.charset.StandardCharsets;
import java.security.SecureRandom;
import javax.crypto.Cipher;
import javax.crypto.KeyGenerator;
import javax.crypto.SecretKey;
import javax.crypto.spec.GCMParameterSpec;

public class AesGcm {
    public static void main(String[] args) throws Exception {
        KeyGenerator keyGen = KeyGenerator.getInstance("AES");
        keyGen.init(256);
        SecretKey key = keyGen.generateKey();

        SecureRandom random = new SecureRandom();
        byte[] iv = new byte[12];
        random.nextBytes(iv);

        byte[] aad = "header".getBytes(StandardCharsets.UTF_8);
        byte[] plaintext = "secret message".getBytes(StandardCharsets.UTF_8);

        Cipher encryptor = Cipher.getInstance("AES/GCM/NoPadding");
        encryptor.init(Cipher.ENCRYPT_MODE, key, new GCMParameterSpec(128, iv));
        encryptor.updateAAD(aad);
        byte[] ciphertext = encryptor.doFinal(plaintext);

        Cipher decryptor = Cipher.getInstance("AES/GCM/NoPadding");
        decryptor.init(Cipher.DECRYPT_MODE, key, new GCMParameterSpec(128, iv));
        decryptor.updateAAD(aad);
        byte[] decrypted = decryptor.doFinal(ciphertext);

        System.out.println(new String(decrypted, StandardCharsets.UTF_8));
    }
}
