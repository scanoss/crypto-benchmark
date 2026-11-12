package bench;

import org.springframework.security.crypto.argon2.Argon2PasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;

public class Argon2Encoder {
    public static void main(String[] args) throws Exception {
        PasswordEncoder encoder = Argon2PasswordEncoder.defaultsForSpringSecurity_v5_8();
        String hash = encoder.encode("correct horse battery staple");
        System.out.println(hash);
        System.out.println(encoder.matches("correct horse battery staple", hash));
    }
}
