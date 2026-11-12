package bench;

import java.io.ByteArrayInputStream;
import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.security.cert.CertificateFactory;
import java.security.cert.X509Certificate;
import java.util.HexFormat;

public class CertificateParse {
    private static final String CERT_PEM = """
            -----BEGIN CERTIFICATE-----
            MIIBhDCCASugAwIBAgIUREADo1EBr86gkWmn9DsGiws1bf4wCgYIKoZIzj0EAwIw
            GDEWMBQGA1UEAwwNYmVuY2guZXhhbXBsZTAeFw0yNjExMDQxMzQ1MThaFw0zNjEx
            MDExMzQ1MThaMBgxFjAUBgNVBAMMDWJlbmNoLmV4YW1wbGUwWTATBgcqhkjOPQIB
            BggqhkjOPQMBBwNCAAQYncJIwjY6rVligU8lLJEqrsEN2RBjxcBe4GLJSm/e1s1R
            wTCzhYUqq4Abz6QIYTarj5oOuPxQcQUWtNF8mIP+o1MwUTAdBgNVHQ4EFgQUZoJd
            24Mxkmi+k9NdvP2ZmwqCx/swHwYDVR0jBBgwFoAUZoJd24Mxkmi+k9NdvP2ZmwqC
            x/swDwYDVR0TAQH/BAUwAwEB/zAKBggqhkjOPQQDAgNHADBEAiAQEZ4fjZrABf/C
            v66M2HzCi7R0KmVc5oJ8cAkmgUEduAIgVsomFlCCt2iqazG2/Vq4rVt+M+t09Tpr
            XSCkRPMsnyM=
            -----END CERTIFICATE-----
            """;

    public static void main(String[] args) throws Exception {
        CertificateFactory factory = CertificateFactory.getInstance("X.509");
        X509Certificate cert = (X509Certificate) factory.generateCertificate(
                new ByteArrayInputStream(CERT_PEM.getBytes(StandardCharsets.US_ASCII)));

        MessageDigest md = MessageDigest.getInstance("SHA-256");
        byte[] fingerprint = md.digest(cert.getEncoded());

        System.out.println(cert.getSubjectX500Principal().getName());
        System.out.println(HexFormat.of().formatHex(fingerprint));
    }
}
