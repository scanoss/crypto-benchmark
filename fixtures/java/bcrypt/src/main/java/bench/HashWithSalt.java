package bench;

import at.favre.lib.crypto.bcrypt.BCrypt;
import java.nio.charset.StandardCharsets;
import java.security.SecureRandom;

public class HashWithSalt {
    public static void main(String[] args) throws Exception {
        byte[] password = "correct horse battery staple".getBytes(StandardCharsets.UTF_8);
        byte[] salt = new byte[16];
        new SecureRandom().nextBytes(salt);
        byte[] hash = BCrypt.withDefaults().hash(12, salt, password);
        System.out.println(new String(hash, StandardCharsets.UTF_8));
        BCrypt.Result result = BCrypt.verifyer().verify(password, hash);
        System.out.println(result.verified);
    }
}
