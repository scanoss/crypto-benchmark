package bench;

import java.nio.charset.StandardCharsets;
import java.security.SecureRandom;
import java.security.Security;
import javax.crypto.Cipher;
import javax.crypto.KeyGenerator;
import javax.crypto.SecretKey;
import javax.crypto.spec.GCMParameterSpec;
import org.bouncycastle.jce.provider.BouncyCastleProvider;

public class ProviderAesGcm {
    public static void main(String[] args) throws Exception {
        Security.addProvider(new BouncyCastleProvider());

        KeyGenerator keyGen = KeyGenerator.getInstance("AES", "BC");
        keyGen.init(256);
        SecretKey key = keyGen.generateKey();

        byte[] nonce = new byte[12];
        new SecureRandom().nextBytes(nonce);
        byte[] plaintext = "attack at dawn".getBytes(StandardCharsets.UTF_8);

        Cipher encryptor = Cipher.getInstance("AES/GCM/NoPadding", "BC");
        encryptor.init(Cipher.ENCRYPT_MODE, key, new GCMParameterSpec(128, nonce));
        byte[] ciphertext = encryptor.doFinal(plaintext);

        Cipher decryptor = Cipher.getInstance("AES/GCM/NoPadding", "BC");
        decryptor.init(Cipher.DECRYPT_MODE, key, new GCMParameterSpec(128, nonce));
        byte[] decrypted = decryptor.doFinal(ciphertext);

        System.out.println(new String(decrypted, StandardCharsets.UTF_8));
    }
}
