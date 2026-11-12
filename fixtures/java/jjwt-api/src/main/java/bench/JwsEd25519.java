package bench;

import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.security.Jwks;
import java.security.KeyPair;

public class JwsEd25519 {
    public static void main(String[] args) throws Exception {
        KeyPair pair = Jwks.CRV.Ed25519.keyPair().build();
        String jws = Jwts.builder()
                .subject("Joe")
                .signWith(pair.getPrivate(), Jwts.SIG.EdDSA)
                .compact();
        String subject = Jwts.parser()
                .verifyWith(pair.getPublic())
                .build()
                .parseSignedClaims(jws)
                .getPayload()
                .getSubject();
        System.out.println("eddsa " + pair.getPrivate().getAlgorithm() + " " + subject);
    }
}
