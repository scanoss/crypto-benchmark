package bench;

import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;

public class BCryptEncoder {
    public static void main(String[] args) throws Exception {
        PasswordEncoder encoder = new BCryptPasswordEncoder(12);
        String hash = encoder.encode("correct horse battery staple");
        System.out.println(hash);
        System.out.println(encoder.matches("correct horse battery staple", hash));
    }
}
