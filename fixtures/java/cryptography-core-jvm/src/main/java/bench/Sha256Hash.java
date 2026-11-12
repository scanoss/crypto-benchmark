package bench;

import dev.whyoleg.cryptography.CryptographyProvider;
import dev.whyoleg.cryptography.algorithms.SHA256;
import java.nio.charset.StandardCharsets;
import java.util.HexFormat;

public class Sha256Hash {
    public static void main(String[] args) throws Exception {
        CryptographyProvider provider = CryptographyProvider.Companion.getDefault();
        byte[] message = "message to digest".getBytes(StandardCharsets.UTF_8);
        byte[] digest = provider.get(SHA256.INSTANCE).hasher().hashBlocking(message);
        System.out.println(HexFormat.of().formatHex(digest));
    }
}
