package bench;

import com.nimbusds.jose.JWSAlgorithm;
import com.nimbusds.jose.JWSHeader;
import com.nimbusds.jose.JWSObject;
import com.nimbusds.jose.JWSSigner;
import com.nimbusds.jose.JWSVerifier;
import com.nimbusds.jose.Payload;
import com.nimbusds.jose.crypto.ECDSASigner;
import com.nimbusds.jose.crypto.ECDSAVerifier;
import com.nimbusds.jose.jwk.Curve;
import com.nimbusds.jose.jwk.ECKey;
import com.nimbusds.jose.jwk.gen.ECKeyGenerator;

public class JwsEs256 {
    public static void main(String[] args) throws Exception {
        ECKey ecJWK = new ECKeyGenerator(Curve.P_256)
                .keyID("123")
                .generate();
        ECKey ecPublicJWK = ecJWK.toPublicJWK();

        JWSSigner signer = new ECDSASigner(ecJWK);
        JWSObject jwsObject = new JWSObject(
                new JWSHeader.Builder(JWSAlgorithm.ES256).keyID(ecJWK.getKeyID()).build(),
                new Payload("Elliptic cure"));
        jwsObject.sign(signer);
        String s = jwsObject.serialize();

        JWSObject parsed = JWSObject.parse(s);
        JWSVerifier verifier = new ECDSAVerifier(ecPublicJWK);
        boolean valid = parsed.verify(verifier);
        System.out.println("es256 " + valid + " " + parsed.getPayload());
    }
}
