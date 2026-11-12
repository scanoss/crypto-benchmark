package bench;

import java.math.BigInteger;
import java.nio.charset.StandardCharsets;
import java.security.SecureRandom;
import org.bouncycastle.crypto.AsymmetricBlockCipher;
import org.bouncycastle.crypto.AsymmetricCipherKeyPair;
import org.bouncycastle.crypto.encodings.OAEPEncoding;
import org.bouncycastle.crypto.engines.RSAEngine;
import org.bouncycastle.crypto.generators.RSAKeyPairGenerator;
import org.bouncycastle.crypto.params.RSAKeyGenerationParameters;

public class RsaOaep {
    public static void main(String[] args) throws Exception {
        RSAKeyPairGenerator generator = new RSAKeyPairGenerator();
        generator.init(new RSAKeyGenerationParameters(BigInteger.valueOf(65537), new SecureRandom(), 3072, 100));
        AsymmetricCipherKeyPair keyPair = generator.generateKeyPair();
        byte[] plaintext = "attack at dawn".getBytes(StandardCharsets.UTF_8);

        AsymmetricBlockCipher encryptor = new OAEPEncoding(new RSAEngine());
        encryptor.init(true, keyPair.getPublic());
        byte[] ciphertext = encryptor.processBlock(plaintext, 0, plaintext.length);

        AsymmetricBlockCipher decryptor = new OAEPEncoding(new RSAEngine());
        decryptor.init(false, keyPair.getPrivate());
        byte[] decrypted = decryptor.processBlock(ciphertext, 0, ciphertext.length);

        System.out.println(new String(decrypted, StandardCharsets.UTF_8));
    }
}
