package bench;

import java.io.ByteArrayOutputStream;
import java.io.InputStream;
import java.io.OutputStream;
import java.nio.charset.StandardCharsets;
import java.security.SecureRandom;
import java.security.Security;
import java.util.Date;
import org.bouncycastle.bcpg.SymmetricKeyAlgorithmTags;
import org.bouncycastle.jce.provider.BouncyCastleProvider;
import org.bouncycastle.openpgp.PGPEncryptedDataGenerator;
import org.bouncycastle.openpgp.PGPEncryptedDataList;
import org.bouncycastle.openpgp.PGPLiteralData;
import org.bouncycastle.openpgp.PGPLiteralDataGenerator;
import org.bouncycastle.openpgp.PGPPBEEncryptedData;
import org.bouncycastle.openpgp.jcajce.JcaPGPObjectFactory;
import org.bouncycastle.openpgp.operator.jcajce.JcaPGPDigestCalculatorProviderBuilder;
import org.bouncycastle.openpgp.operator.jcajce.JcePBEDataDecryptorFactoryBuilder;
import org.bouncycastle.openpgp.operator.jcajce.JcePBEKeyEncryptionMethodGenerator;
import org.bouncycastle.openpgp.operator.jcajce.JcePGPDataEncryptorBuilder;

public class PgpPassword {
    public static void main(String[] args) throws Exception {
        Security.addProvider(new BouncyCastleProvider());
        char[] passphrase = "correct horse battery staple".toCharArray();

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
        encryptedGen.addMethod(new JcePBEKeyEncryptionMethodGenerator(passphrase).setProvider("BC"));
        ByteArrayOutputStream encrypted = new ByteArrayOutputStream();
        try (OutputStream out = encryptedGen.open(encrypted, new byte[4096])) {
            out.write(literal.toByteArray());
        }

        JcaPGPObjectFactory factory = new JcaPGPObjectFactory(encrypted.toByteArray());
        PGPEncryptedDataList encryptedList = (PGPEncryptedDataList) factory.nextObject();
        PGPPBEEncryptedData encryptedData = (PGPPBEEncryptedData) encryptedList.get(0);
        InputStream clear = encryptedData.getDataStream(
                new JcePBEDataDecryptorFactoryBuilder(new JcaPGPDigestCalculatorProviderBuilder().setProvider("BC").build())
                        .setProvider("BC").build(passphrase));
        PGPLiteralData literalData = (PGPLiteralData) new JcaPGPObjectFactory(clear).nextObject();
        byte[] decrypted = literalData.getInputStream().readAllBytes();

        System.out.println(new String(decrypted, StandardCharsets.UTF_8) + " integrity: " + encryptedData.verify());
    }
}
