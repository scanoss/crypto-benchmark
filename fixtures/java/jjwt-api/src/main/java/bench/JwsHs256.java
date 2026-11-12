package bench;

import io.jsonwebtoken.Jwts;
import javax.crypto.SecretKey;

public class JwsHs256 {
    public static void main(String[] args) throws Exception {
        SecretKey key = Jwts.SIG.HS256.key().build();
        String jws = Jwts.builder()
                .subject("Joe")
                .signWith(key)
                .compact();
        String subject = Jwts.parser()
                .verifyWith(key)
                .build()
                .parseSignedClaims(jws)
                .getPayload()
                .getSubject();
        System.out.println("hs256 " + subject);
    }
}
