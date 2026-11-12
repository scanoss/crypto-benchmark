package bench;

import io.jsonwebtoken.Jwts;
import java.nio.charset.StandardCharsets;
import javax.crypto.SecretKey;

public class JweDirect {
    public static void main(String[] args) throws Exception {
        SecretKey key = Jwts.ENC.A256GCM.key().build();
        byte[] content = "Hello World".getBytes(StandardCharsets.UTF_8);
        String jwe = Jwts.builder()
                .content(content, "text/plain")
                .encryptWith(key, Jwts.ENC.A256GCM)
                .compact();
        byte[] decrypted = Jwts.parser()
                .decryptWith(key)
                .build()
                .parseEncryptedContent(jwe)
                .getPayload();
        System.out.println("dir " + new String(decrypted, StandardCharsets.UTF_8));
    }
}
