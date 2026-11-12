package bench;

import java.nio.charset.StandardCharsets;
import java.security.KeyPairGenerator;
import java.security.Security;
import java.util.Date;
import org.bouncycastle.bcpg.HashAlgorithmTags;
import org.bouncycastle.jce.provider.BouncyCastleProvider;
import org.bouncycastle.openpgp.PGPKeyPair;
import org.bouncycastle.openpgp.PGPPublicKey;
import org.bouncycastle.openpgp.PGPSignature;
import org.bouncycastle.openpgp.PGPSignatureGenerator;
import org.bouncycastle.openpgp.PGPSignatureList;
import org.bouncycastle.openpgp.jcajce.JcaPGPObjectFactory;
import org.bouncycastle.openpgp.operator.jcajce.JcaPGPContentSignerBuilder;
import org.bouncycastle.openpgp.operator.jcajce.JcaPGPContentVerifierBuilderProvider;
import org.bouncycastle.openpgp.operator.jcajce.JcaPGPKeyPair;

public class PgpSign {
    public static void main(String[] args) throws Exception {
        Security.addProvider(new BouncyCastleProvider());

        KeyPairGenerator keyPairGen = KeyPairGenerator.getInstance("RSA", "BC");
        keyPairGen.initialize(3072);
        PGPKeyPair keyPair = new JcaPGPKeyPair(PGPPublicKey.RSA_GENERAL, keyPairGen.generateKeyPair(), new Date());
        byte[] message = "hello world".getBytes(StandardCharsets.UTF_8);

        PGPSignatureGenerator signatureGen = new PGPSignatureGenerator(
                new JcaPGPContentSignerBuilder(keyPair.getPublicKey().getAlgorithm(), HashAlgorithmTags.SHA256).setProvider("BC"));
        signatureGen.init(PGPSignature.BINARY_DOCUMENT, keyPair.getPrivateKey());
        signatureGen.update(message);
        byte[] encoded = signatureGen.generate().getEncoded();

        PGPSignatureList signatures = (PGPSignatureList) new JcaPGPObjectFactory(encoded).nextObject();
        PGPSignature signature = signatures.get(0);
        signature.init(new JcaPGPContentVerifierBuilderProvider().setProvider("BC"), keyPair.getPublicKey());
        signature.update(message);

        System.out.println("valid: " + signature.verify());
    }
}
