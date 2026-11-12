package bench;

import java.nio.charset.StandardCharsets;
import java.security.SecureRandom;
import org.bouncycastle.crypto.generators.Argon2BytesGenerator;
import org.bouncycastle.crypto.params.Argon2Parameters;
import org.bouncycastle.util.encoders.Hex;

public class Argon2id {
    public static void main(String[] args) {
        byte[] password = "correct horse battery staple".getBytes(StandardCharsets.UTF_8);
        byte[] salt = new byte[16];
        new SecureRandom().nextBytes(salt);

        Argon2Parameters params = new Argon2Parameters.Builder(Argon2Parameters.ARGON2_id)
                .withSalt(salt)
                .withIterations(3)
                .withMemoryAsKB(65536)
                .withParallelism(4)
                .build();
        Argon2BytesGenerator generator = new Argon2BytesGenerator();
        generator.init(params);
        byte[] derived = new byte[32];
        generator.generateBytes(password, derived);

        System.out.println(Hex.toHexString(derived));
    }
}
