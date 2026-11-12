package bench;

import dev.whyoleg.cryptography.CryptographyProvider;
import dev.whyoleg.cryptography.algorithms.AES;
import dev.whyoleg.cryptography.random.CryptographyRandom;
import java.nio.charset.StandardCharsets;

public class AesCbcEncrypt {
    public static void main(String[] args) throws Exception {
        CryptographyProvider provider = CryptographyProvider.Companion.getDefault();
        byte[] keyBytes = CryptographyRandom.Default.nextBytes(32);
        AES.CBC.Key key = provider.get(AES.CBC.Companion).keyDecoder().decodeFromByteArrayBlocking(AES.Key.Format.RAW, keyBytes);
        byte[] plaintext = "attack at dawn".getBytes(StandardCharsets.UTF_8);
        byte[] ciphertext = key.cipher(true).encryptBlocking(plaintext);
        byte[] decrypted = key.cipher(true).decryptBlocking(ciphertext);
        System.out.println(ciphertext.length);
        System.out.println(new String(decrypted, StandardCharsets.UTF_8));
    }
}
