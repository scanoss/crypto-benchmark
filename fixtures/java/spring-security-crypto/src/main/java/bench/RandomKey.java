package bench;

import java.util.HexFormat;
import org.springframework.security.crypto.keygen.BytesKeyGenerator;
import org.springframework.security.crypto.keygen.KeyGenerators;

public class RandomKey {
    public static void main(String[] args) throws Exception {
        BytesKeyGenerator generator = KeyGenerators.secureRandom(16);
        byte[] key = generator.generateKey();
        System.out.println(HexFormat.of().formatHex(key));
    }
}
