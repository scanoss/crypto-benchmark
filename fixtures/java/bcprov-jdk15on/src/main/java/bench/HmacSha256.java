package bench;

import java.nio.charset.StandardCharsets;
import java.security.SecureRandom;
import org.bouncycastle.crypto.digests.SHA256Digest;
import org.bouncycastle.crypto.macs.HMac;
import org.bouncycastle.crypto.params.KeyParameter;
import org.bouncycastle.util.encoders.Hex;

public class HmacSha256 {
    public static void main(String[] args) {
        byte[] keyBytes = new byte[32];
        new SecureRandom().nextBytes(keyBytes);
        byte[] message = "hello world".getBytes(StandardCharsets.UTF_8);

        HMac hmac = new HMac(new SHA256Digest());
        hmac.init(new KeyParameter(keyBytes));
        hmac.update(message, 0, message.length);
        byte[] mac = new byte[hmac.getMacSize()];
        hmac.doFinal(mac, 0);

        System.out.println(Hex.toHexString(mac));
    }
}
