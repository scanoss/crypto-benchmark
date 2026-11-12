package bench;

import java.nio.charset.StandardCharsets;
import java.security.SecureRandom;
import org.bouncycastle.crypto.engines.AESEngine;
import org.bouncycastle.crypto.modes.GCMBlockCipher;
import org.bouncycastle.crypto.params.AEADParameters;
import org.bouncycastle.crypto.params.KeyParameter;

public class AesGcm {
    public static void main(String[] args) throws Exception {
        SecureRandom random = new SecureRandom();
        byte[] keyBytes = new byte[32];
        random.nextBytes(keyBytes);
        KeyParameter key = new KeyParameter(keyBytes);
        byte[] nonce = new byte[12];
        random.nextBytes(nonce);
        byte[] plaintext = "attack at dawn".getBytes(StandardCharsets.UTF_8);

        GCMBlockCipher encryptor = new GCMBlockCipher(new AESEngine());
        encryptor.init(true, new AEADParameters(key, 128, nonce));
        byte[] ciphertext = new byte[encryptor.getOutputSize(plaintext.length)];
        int len = encryptor.processBytes(plaintext, 0, plaintext.length, ciphertext, 0);
        encryptor.doFinal(ciphertext, len);

        GCMBlockCipher decryptor = new GCMBlockCipher(new AESEngine());
        decryptor.init(false, new AEADParameters(key, 128, nonce));
        byte[] decrypted = new byte[decryptor.getOutputSize(ciphertext.length)];
        len = decryptor.processBytes(ciphertext, 0, ciphertext.length, decrypted, 0);
        decryptor.doFinal(decrypted, len);

        System.out.println(new String(decrypted, StandardCharsets.UTF_8));
    }
}
