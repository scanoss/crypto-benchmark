package bench;

import java.io.StringReader;
import java.io.StringWriter;
import java.math.BigInteger;
import java.security.KeyPair;
import java.security.KeyPairGenerator;
import java.security.Security;
import java.security.cert.X509Certificate;
import java.security.spec.ECGenParameterSpec;
import java.util.Date;
import org.bouncycastle.asn1.x500.X500Name;
import org.bouncycastle.cert.X509CertificateHolder;
import org.bouncycastle.cert.jcajce.JcaX509CertificateConverter;
import org.bouncycastle.cert.jcajce.JcaX509v3CertificateBuilder;
import org.bouncycastle.jce.provider.BouncyCastleProvider;
import org.bouncycastle.openssl.PEMKeyPair;
import org.bouncycastle.openssl.PEMParser;
import org.bouncycastle.openssl.jcajce.JcaPEMKeyConverter;
import org.bouncycastle.openssl.jcajce.JcaPEMWriter;
import org.bouncycastle.operator.ContentSigner;
import org.bouncycastle.operator.jcajce.JcaContentSignerBuilder;

public class PemRoundTrip {
    public static void main(String[] args) throws Exception {
        Security.addProvider(new BouncyCastleProvider());

        KeyPairGenerator keyPairGen = KeyPairGenerator.getInstance("EC", "BC");
        keyPairGen.initialize(new ECGenParameterSpec("secp256r1"));
        KeyPair keyPair = keyPairGen.generateKeyPair();

        X500Name subject = new X500Name("CN=bench.example");
        Date notBefore = new Date();
        Date notAfter = new Date(notBefore.getTime() + 365L * 24 * 60 * 60 * 1000);
        ContentSigner signer = new JcaContentSignerBuilder("SHA256withECDSA").setProvider("BC").build(keyPair.getPrivate());
        X509CertificateHolder holder = new JcaX509v3CertificateBuilder(
                subject, BigInteger.ONE, notBefore, notAfter, subject, keyPair.getPublic())
                .build(signer);

        StringWriter pem = new StringWriter();
        try (JcaPEMWriter writer = new JcaPEMWriter(pem)) {
            writer.writeObject(keyPair.getPrivate());
            writer.writeObject(holder);
        }

        try (PEMParser parser = new PEMParser(new StringReader(pem.toString()))) {
            PEMKeyPair pemKeyPair = (PEMKeyPair) parser.readObject();
            KeyPair parsedKeyPair = new JcaPEMKeyConverter().setProvider("BC").getKeyPair(pemKeyPair);
            X509CertificateHolder parsedHolder = (X509CertificateHolder) parser.readObject();
            X509Certificate cert = new JcaX509CertificateConverter().setProvider("BC").getCertificate(parsedHolder);

            cert.verify(parsedKeyPair.getPublic());
            System.out.println(cert.getSubjectX500Principal().getName());
        }
    }
}
