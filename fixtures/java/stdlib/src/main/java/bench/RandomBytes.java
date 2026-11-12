package bench;

import java.security.SecureRandom;
import java.util.HexFormat;

public class RandomBytes {
    public static void main(String[] args) throws Exception {
        SecureRandom random = new SecureRandom();
        byte[] token = new byte[32];
        random.nextBytes(token);

        SecureRandom strong = SecureRandom.getInstanceStrong();
        byte[] strongToken = new byte[32];
        strong.nextBytes(strongToken);

        byte[] seed = random.generateSeed(16);

        System.out.println(HexFormat.of().formatHex(token));
        System.out.println(HexFormat.of().formatHex(strongToken));
        System.out.println(HexFormat.of().formatHex(seed));
    }
}
