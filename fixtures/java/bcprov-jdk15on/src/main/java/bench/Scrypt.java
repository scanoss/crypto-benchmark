package bench;

import java.nio.charset.StandardCharsets;
import java.security.SecureRandom;
import org.bouncycastle.crypto.generators.SCrypt;
import org.bouncycastle.util.encoders.Hex;

public class Scrypt {
    public static void main(String[] args) {
        byte[] password = "correct horse battery staple".getBytes(StandardCharsets.UTF_8);
        byte[] salt = new byte[16];
        new SecureRandom().nextBytes(salt);

        byte[] derived = SCrypt.generate(password, salt, 32768, 8, 1, 32);

        System.out.println(Hex.toHexString(derived));
    }
}
