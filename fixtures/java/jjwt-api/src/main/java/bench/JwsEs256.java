package bench;

import io.jsonwebtoken.Jwts;
import java.security.KeyPair;

public class JwsEs256 {
    public static void main(String[] args) throws Exception {
        KeyPair pair = Jwts.SIG.ES256.keyPair().build();
        String jws = Jwts.builder()
                .subject("Joe")
                .signWith(pair.getPrivate(), Jwts.SIG.ES256)
                .compact();
        String subject = Jwts.parser()
                .verifyWith(pair.getPublic())
                .build()
                .parseSignedClaims(jws)
                .getPayload()
                .getSubject();
        System.out.println("es256 " + subject);
    }
}
