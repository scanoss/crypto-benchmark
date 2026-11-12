package bench;

import dev.whyoleg.cryptography.CryptographyProvider;
import dev.whyoleg.cryptography.algorithms.HMAC;
import dev.whyoleg.cryptography.algorithms.SHA256;
import java.nio.charset.StandardCharsets;
import java.util.HexFormat;

public class HmacSha256 {
    public static void main(String[] args) throws Exception {
        CryptographyProvider provider = CryptographyProvider.Companion.getDefault();
        HMAC.Key key = provider.get(HMAC.Companion).keyGenerator(SHA256.INSTANCE).generateKeyBlocking();
        byte[] message = "message to authenticate".getBytes(StandardCharsets.UTF_8);
        byte[] signature = key.signatureGenerator().generateSignatureBlocking(message);
        System.out.println(HexFormat.of().formatHex(signature));
        System.out.println(key.signatureVerifier().tryVerifySignatureBlocking(message, signature));
    }
}
