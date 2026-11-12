package bench;

import io.jsonwebtoken.Jwts;
import java.security.KeyPair;

public class JwsRs256 {
    public static void main(String[] args) throws Exception {
        KeyPair pair = Jwts.SIG.RS256.keyPair().build();
        String jws = Jwts.builder()
                .subject("Joe")
                .issuer("crypto-benchmark")
                .signWith(pair.getPrivate())
                .compact();
        String subject = Jwts.parser()
                .verifyWith(pair.getPublic())
                .requireIssuer("crypto-benchmark")
                .build()
                .parseSignedClaims(jws)
                .getPayload()
                .getSubject();
        System.out.println("rs256 " + subject);
    }
}
