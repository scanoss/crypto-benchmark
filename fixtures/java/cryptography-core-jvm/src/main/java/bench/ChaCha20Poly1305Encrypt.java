package bench;

import dev.whyoleg.cryptography.CryptographyProvider;
import dev.whyoleg.cryptography.algorithms.ChaCha20Poly1305;
import java.nio.charset.StandardCharsets;

public class ChaCha20Poly1305Encrypt {
    public static void main(String[] args) throws Exception {
        CryptographyProvider provider = CryptographyProvider.Companion.getDefault();
        ChaCha20Poly1305.Key key = provider.get(ChaCha20Poly1305.Companion).keyGenerator().generateKeyBlocking();
        byte[] plaintext = "attack at dawn".getBytes(StandardCharsets.UTF_8);
        byte[] ciphertext = key.cipher().encryptBlocking(plaintext);
        byte[] decrypted = key.cipher().decryptBlocking(ciphertext);
        System.out.println(ciphertext.length);
        System.out.println(new String(decrypted, StandardCharsets.UTF_8));
    }
}
