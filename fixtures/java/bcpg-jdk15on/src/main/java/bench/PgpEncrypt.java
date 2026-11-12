package bench;

import java.io.ByteArrayOutputStream;
import java.io.InputStream;
import java.io.OutputStream;
import java.nio.charset.StandardCharsets;
import java.security.KeyPairGenerator;
import java.security.SecureRandom;
import java.security.Security;
import java.util.Date;
import org.bouncycastle.bcpg.SymmetricKeyAlgorithmTags;
import org.bouncycastle.jce.provider.BouncyCastleProvider;
import org.bouncycastle.openpgp.PGPEncryptedDataGenerator;
import org.bouncycastle.openpgp.PGPEncryptedDataList;
import org.bouncycastle.openpgp.PGPKeyPair;
import org.bouncycastle.openpgp.PGPLiteralData;
import org.bouncycastle.openpgp.PGPLiteralDataGenerator;
import org.bouncycastle.openpgp.PGPPublicKey;
import org.bouncycastle.openpgp.PGPPublicKeyEncryptedData;
import org.bouncycastle.openpgp.jcajce.JcaPGPObjectFactory;
import org.bouncycastle.openpgp.operator.jcajce.JcaPGPKeyPair;
import org.bouncycastle.openpgp.operator.jcajce.JcePGPDataEncryptorBuilder;
import org.bouncycastle.openpgp.operator.jcajce.JcePublicKeyDataDecryptorFactoryBuilder;
import org.bouncycastle.openpgp.operator.jcajce.JcePublicKeyKeyEncryptionMethodGenerator;

public class PgpEncrypt {
    public static void main(String[] args) throws Exception {
        Security.addProvider(new BouncyCastleProvider());

        KeyPairGenerator keyPairGen = KeyPairGenerator.getInstance("RSA", "BC");
        keyPairGen.initialize(3072);
        PGPKeyPair keyPair = new JcaPGPKeyPair(PGPPublicKey.RSA_GENERAL, keyPairGen.generateKeyPair(), new Date());

        byte[] message = "attack at dawn".getBytes(StandardCharsets.UTF_8);
        ByteArrayOutputStream literal = new ByteArrayOutputStream();
        PGPLiteralDataGenerator literalGen = new PGPLiteralDataGenerator();
        try (OutputStream out = literalGen.open(literal, PGPLiteralData.BINARY, "message.txt", message.length, new Date())) {
            out.write(message);
        }

        PGPEncryptedDataGenerator encryptedGen = new PGPEncryptedDataGenerator(
                new JcePGPDataEncryptorBuilder(SymmetricKeyAlgorithmTags.AES_256)
                        .setWithIntegrityPacket(true)
                        .setSecureRandom(new SecureRandom())
                        .setProvider("BC"));
        encryptedGen.addMethod(new JcePublicKeyKeyEncryptionMethodGenerator(keyPair.getPublicKey()).setProvider("BC"));
        ByteArrayOutputStream encrypted = new ByteArrayOutputStream();
        try (OutputStream out = encryptedGen.open(encrypted, new byte[4096])) {
            out.write(literal.toByteArray());
        }

        JcaPGPObjectFactory factory = new JcaPGPObjectFactory(encrypted.toByteArray());
        PGPEncryptedDataList encryptedList = (PGPEncryptedDataList) factory.nextObject();
        PGPPublicKeyEncryptedData encryptedData = (PGPPublicKeyEncryptedData) encryptedList.get(0);
        InputStream clear = encryptedData.getDataStream(
                new JcePublicKeyDataDecryptorFactoryBuilder().setProvider("BC").build(keyPair.getPrivateKey()));
        PGPLiteralData literalData = (PGPLiteralData) new JcaPGPObjectFactory(clear).nextObject();
        byte[] decrypted = literalData.getInputStream().readAllBytes();

        System.out.println(new String(decrypted, StandardCharsets.UTF_8) + " integrity: " + encryptedData.verify());
    }
}
