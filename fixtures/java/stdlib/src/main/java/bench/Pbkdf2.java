package bench;

import java.security.SecureRandom;
import java.util.HexFormat;
import javax.crypto.SecretKey;
import javax.crypto.SecretKeyFactory;
import javax.crypto.spec.PBEKeySpec;
import javax.crypto.spec.SecretKeySpec;

public class Pbkdf2 {
    public static void main(String[] args) throws Exception {
        char[] password = "correct horse battery staple".toCharArray();
        byte[] salt = new byte[16];
        new SecureRandom().nextBytes(salt);

        PBEKeySpec spec = new PBEKeySpec(password, salt, 600000, 256);
        SecretKeyFactory factory = SecretKeyFactory.getInstance("PBKDF2WithHmacSHA256");
        byte[] derived = factory.generateSecret(spec).getEncoded();
        spec.clearPassword();

        SecretKey aesKey = new SecretKeySpec(derived, "AES");
        System.out.println(aesKey.getAlgorithm() + " " + HexFormat.of().formatHex(derived));
    }
}
