package bench;

import java.io.ByteArrayOutputStream;
import java.security.KeyPairGenerator;
import java.security.Security;
import java.util.Date;
import org.bouncycastle.bcpg.ArmoredOutputStream;
import org.bouncycastle.bcpg.HashAlgorithmTags;
import org.bouncycastle.jce.provider.BouncyCastleProvider;
import org.bouncycastle.openpgp.PGPEncryptedData;
import org.bouncycastle.openpgp.PGPKeyPair;
import org.bouncycastle.openpgp.PGPKeyRingGenerator;
import org.bouncycastle.openpgp.PGPPublicKey;
import org.bouncycastle.openpgp.PGPPublicKeyRing;
import org.bouncycastle.openpgp.PGPSecretKeyRing;
import org.bouncycastle.openpgp.PGPSignature;
import org.bouncycastle.openpgp.operator.PGPDigestCalculator;
import org.bouncycastle.openpgp.operator.jcajce.JcaPGPContentSignerBuilder;
import org.bouncycastle.openpgp.operator.jcajce.JcaPGPDigestCalculatorProviderBuilder;
import org.bouncycastle.openpgp.operator.jcajce.JcaPGPKeyPair;
import org.bouncycastle.openpgp.operator.jcajce.JcePBESecretKeyEncryptorBuilder;

public class PgpKeyRing {
    public static void main(String[] args) throws Exception {
        Security.addProvider(new BouncyCastleProvider());
        char[] passphrase = "correct horse battery staple".toCharArray();

        KeyPairGenerator keyPairGen = KeyPairGenerator.getInstance("RSA", "BC");
        keyPairGen.initialize(3072);
        PGPKeyPair signKey = new JcaPGPKeyPair(PGPPublicKey.RSA_SIGN, keyPairGen.generateKeyPair(), new Date());
        PGPKeyPair encryptKey = new JcaPGPKeyPair(PGPPublicKey.RSA_ENCRYPT, keyPairGen.generateKeyPair(), new Date());

        PGPDigestCalculator sha1Calc = new JcaPGPDigestCalculatorProviderBuilder().build().get(HashAlgorithmTags.SHA1);
        PGPKeyRingGenerator keyRingGen = new PGPKeyRingGenerator(
                PGPSignature.POSITIVE_CERTIFICATION, signKey, "alice@example.com", sha1Calc, null, null,
                new JcaPGPContentSignerBuilder(signKey.getPublicKey().getAlgorithm(), HashAlgorithmTags.SHA256),
                new JcePBESecretKeyEncryptorBuilder(PGPEncryptedData.AES_256, sha1Calc).setProvider("BC").build(passphrase));
        keyRingGen.addSubKey(encryptKey);
        PGPPublicKeyRing publicKeyRing = keyRingGen.generatePublicKeyRing();
        PGPSecretKeyRing secretKeyRing = keyRingGen.generateSecretKeyRing();

        ByteArrayOutputStream armored = new ByteArrayOutputStream();
        try (ArmoredOutputStream out = new ArmoredOutputStream(armored)) {
            publicKeyRing.encode(out);
        }
        try (ArmoredOutputStream out = new ArmoredOutputStream(armored)) {
            secretKeyRing.encode(out);
        }

        System.out.println(armored.toString().lines().findFirst().orElse(""));
    }
}
