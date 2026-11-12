package bench;

import java.nio.charset.StandardCharsets;
import org.springframework.security.crypto.encrypt.BytesEncryptor;
import org.springframework.security.crypto.encrypt.Encryptors;
import org.springframework.security.crypto.keygen.KeyGenerators;

public class StrongerEncryption {
    public static void main(String[] args) throws Exception {
        String salt = KeyGenerators.string().generateKey();
        BytesEncryptor encryptor = Encryptors.stronger("encryption-password", salt);
        byte[] encrypted = encryptor.encrypt("api token".getBytes(StandardCharsets.UTF_8));
        byte[] decrypted = encryptor.decrypt(encrypted);
        System.out.println(encrypted.length);
        System.out.println(new String(decrypted, StandardCharsets.UTF_8));
    }
}
