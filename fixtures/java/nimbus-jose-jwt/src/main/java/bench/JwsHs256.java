package bench;

import com.nimbusds.jose.JWSAlgorithm;
import com.nimbusds.jose.JWSHeader;
import com.nimbusds.jose.JWSObject;
import com.nimbusds.jose.JWSSigner;
import com.nimbusds.jose.JWSVerifier;
import com.nimbusds.jose.Payload;
import com.nimbusds.jose.crypto.MACSigner;
import com.nimbusds.jose.crypto.MACVerifier;
import java.security.SecureRandom;

public class JwsHs256 {
    public static void main(String[] args) throws Exception {
        SecureRandom random = new SecureRandom();
        byte[] sharedSecret = new byte[32];
        random.nextBytes(sharedSecret);

        JWSSigner signer = new MACSigner(sharedSecret);
        JWSObject jwsObject = new JWSObject(new JWSHeader(JWSAlgorithm.HS256), new Payload("Hello, world!"));
        jwsObject.sign(signer);
        String s = jwsObject.serialize();

        JWSObject parsed = JWSObject.parse(s);
        JWSVerifier verifier = new MACVerifier(sharedSecret);
        boolean valid = parsed.verify(verifier);
        System.out.println("hs256 " + valid + " " + parsed.getPayload());
    }
}
