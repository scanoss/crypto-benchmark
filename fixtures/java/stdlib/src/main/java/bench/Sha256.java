package bench;

import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.util.HexFormat;

public class Sha256 {
    public static void main(String[] args) throws Exception {
        MessageDigest md = MessageDigest.getInstance("SHA-256");
        byte[] hash = md.digest("hello world".getBytes(StandardCharsets.UTF_8));
        System.out.println(HexFormat.of().formatHex(hash));
    }
}
