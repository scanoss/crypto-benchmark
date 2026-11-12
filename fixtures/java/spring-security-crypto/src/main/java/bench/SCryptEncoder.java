package bench;

import org.springframework.security.crypto.scrypt.SCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;

public class SCryptEncoder {
    public static void main(String[] args) throws Exception {
        PasswordEncoder encoder = SCryptPasswordEncoder.defaultsForSpringSecurity_v5_8();
        String hash = encoder.encode("correct horse battery staple");
        System.out.println(hash);
        System.out.println(encoder.matches("correct horse battery staple", hash));
    }
}
