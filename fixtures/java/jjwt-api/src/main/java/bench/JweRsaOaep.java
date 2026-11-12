package bench;

import io.jsonwebtoken.Jwts;
import java.security.KeyPair;

public class JweRsaOaep {
    public static void main(String[] args) throws Exception {
        KeyPair pair = Jwts.SIG.RS512.keyPair().build();
        String jwe = Jwts.builder()
                .audience().add("Alice").and()
                .subject("Joe")
                .encryptWith(pair.getPublic(), Jwts.KEY.RSA_OAEP_256, Jwts.ENC.A256GCM)
                .compact();
        String subject = Jwts.parser()
                .decryptWith(pair.getPrivate())
                .build()
                .parseEncryptedClaims(jwe)
                .getPayload()
                .getSubject();
        System.out.println("rsa-oaep-256 " + subject);
    }
}
