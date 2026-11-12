package bench;

import org.springframework.security.crypto.factory.PasswordEncoderFactories;
import org.springframework.security.crypto.password.PasswordEncoder;

public class DelegatingEncoder {
    public static void main(String[] args) throws Exception {
        PasswordEncoder encoder = PasswordEncoderFactories.createDelegatingPasswordEncoder();
        String hash = encoder.encode("correct horse battery staple");
        System.out.println(hash);
        System.out.println(encoder.matches("correct horse battery staple", hash));
    }
}
