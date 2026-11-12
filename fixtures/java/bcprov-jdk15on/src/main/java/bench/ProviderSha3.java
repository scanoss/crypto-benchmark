package bench;

import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.security.Security;
import org.bouncycastle.jce.provider.BouncyCastleProvider;
import org.bouncycastle.util.encoders.Hex;

public class ProviderSha3 {
    public static void main(String[] args) throws Exception {
        Security.addProvider(new BouncyCastleProvider());

        MessageDigest digest = MessageDigest.getInstance("SHA3-256", "BC");
        byte[] hash = digest.digest("hello world".getBytes(StandardCharsets.UTF_8));

        System.out.println(Hex.toHexString(hash));
    }
}
