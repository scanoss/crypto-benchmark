package bench;

import java.nio.charset.StandardCharsets;
import org.bouncycastle.crypto.digests.SHA256Digest;
import org.bouncycastle.util.encoders.Hex;

public class Sha256 {
    public static void main(String[] args) {
        byte[] message = "hello world".getBytes(StandardCharsets.UTF_8);

        SHA256Digest digest = new SHA256Digest();
        digest.update(message, 0, message.length);
        byte[] hash = new byte[digest.getDigestSize()];
        digest.doFinal(hash, 0);

        System.out.println(Hex.toHexString(hash));
    }
}
