package bench;

import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.util.HexFormat;

public class Sha3 {
    public static void main(String[] args) throws Exception {
        MessageDigest md = MessageDigest.getInstance("SHA3-256");
        md.update("hello ".getBytes(StandardCharsets.UTF_8));
        md.update("world".getBytes(StandardCharsets.UTF_8));
        byte[] hash = md.digest();
        System.out.println(HexFormat.of().formatHex(hash));
    }
}
