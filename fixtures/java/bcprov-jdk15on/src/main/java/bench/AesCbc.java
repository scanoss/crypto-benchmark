package bench;

import java.nio.charset.StandardCharsets;
import java.security.SecureRandom;
import java.util.Arrays;
import org.bouncycastle.crypto.engines.AESEngine;
import org.bouncycastle.crypto.modes.CBCBlockCipher;
import org.bouncycastle.crypto.paddings.PaddedBufferedBlockCipher;
import org.bouncycastle.crypto.params.KeyParameter;
import org.bouncycastle.crypto.params.ParametersWithIV;

public class AesCbc {
    public static void main(String[] args) throws Exception {
        SecureRandom random = new SecureRandom();
        byte[] keyBytes = new byte[32];
        random.nextBytes(keyBytes);
        KeyParameter key = new KeyParameter(keyBytes);
        byte[] iv = new byte[16];
        random.nextBytes(iv);
        byte[] plaintext = "attack at dawn".getBytes(StandardCharsets.UTF_8);

        PaddedBufferedBlockCipher encryptor = new PaddedBufferedBlockCipher(new CBCBlockCipher(new AESEngine()));
        encryptor.init(true, new ParametersWithIV(key, iv));
        byte[] ciphertext = new byte[encryptor.getOutputSize(plaintext.length)];
        int len = encryptor.processBytes(plaintext, 0, plaintext.length, ciphertext, 0);
        len += encryptor.doFinal(ciphertext, len);
        ciphertext = Arrays.copyOf(ciphertext, len);

        PaddedBufferedBlockCipher decryptor = new PaddedBufferedBlockCipher(new CBCBlockCipher(new AESEngine()));
        decryptor.init(false, new ParametersWithIV(key, iv));
        byte[] decrypted = new byte[decryptor.getOutputSize(ciphertext.length)];
        len = decryptor.processBytes(ciphertext, 0, ciphertext.length, decrypted, 0);
        len += decryptor.doFinal(decrypted, len);

        System.out.println(new String(decrypted, 0, len, StandardCharsets.UTF_8));
    }
}
