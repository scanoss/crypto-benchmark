package bench;

import org.springframework.security.crypto.encrypt.Encryptors;
import org.springframework.security.crypto.encrypt.TextEncryptor;
import org.springframework.security.crypto.keygen.KeyGenerators;

public class TextEncryption {
    public static void main(String[] args) throws Exception {
        String salt = KeyGenerators.string().generateKey();
        TextEncryptor encryptor = Encryptors.text("encryption-password", salt);
        String encrypted = encryptor.encrypt("api token");
        String decrypted = encryptor.decrypt(encrypted);
        System.out.println(encrypted);
        System.out.println(decrypted);
    }
}
