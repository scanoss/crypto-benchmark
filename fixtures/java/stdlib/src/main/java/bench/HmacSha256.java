package bench;

import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.util.HexFormat;
import javax.crypto.KeyGenerator;
import javax.crypto.Mac;
import javax.crypto.SecretKey;

public class HmacSha256 {
    public static void main(String[] args) throws Exception {
        byte[] message = "authenticated message".getBytes(StandardCharsets.UTF_8);

        KeyGenerator keyGen = KeyGenerator.getInstance("HmacSHA256");
        SecretKey key = keyGen.generateKey();

        Mac mac = Mac.getInstance("HmacSHA256");
        mac.init(key);
        byte[] tag = mac.doFinal(message);

        Mac verifier = Mac.getInstance("HmacSHA256");
        verifier.init(key);
        byte[] expected = verifier.doFinal(message);

        System.out.println(HexFormat.of().formatHex(tag));
        System.out.println("valid: " + MessageDigest.isEqual(tag, expected));
    }
}
