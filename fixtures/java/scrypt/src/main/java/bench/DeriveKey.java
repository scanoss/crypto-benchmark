package bench;

import com.lambdaworks.crypto.SCrypt;
import java.nio.charset.StandardCharsets;
import java.security.SecureRandom;
import java.util.HexFormat;

public class DeriveKey {
    public static void main(String[] args) throws Exception {
        byte[] password = "correct horse battery staple".getBytes(StandardCharsets.UTF_8);
        byte[] salt = new byte[16];
        new SecureRandom().nextBytes(salt);
        byte[] key = SCrypt.scrypt(password, salt, 16384, 8, 1, 32);
        System.out.println(HexFormat.of().formatHex(key));
    }
}
