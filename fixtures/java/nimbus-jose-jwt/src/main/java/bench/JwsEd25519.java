package bench;

import com.nimbusds.jose.JWSAlgorithm;
import com.nimbusds.jose.JWSHeader;
import com.nimbusds.jose.JWSObject;
import com.nimbusds.jose.JWSSigner;
import com.nimbusds.jose.JWSVerifier;
import com.nimbusds.jose.Payload;
import com.nimbusds.jose.crypto.Ed25519Signer;
import com.nimbusds.jose.crypto.Ed25519Verifier;
import com.nimbusds.jose.jwk.Curve;
import com.nimbusds.jose.jwk.OctetKeyPair;
import com.nimbusds.jose.jwk.gen.OctetKeyPairGenerator;

public class JwsEd25519 {
    public static void main(String[] args) throws Exception {
        OctetKeyPair jwk = new OctetKeyPairGenerator(Curve.Ed25519)
                .keyID("123")
                .generate();
        OctetKeyPair publicJWK = jwk.toPublicJWK();

        JWSSigner signer = new Ed25519Signer(jwk);
        JWSObject jwsObject = new JWSObject(
                new JWSHeader.Builder(JWSAlgorithm.EdDSA).keyID(jwk.getKeyID()).build(),
                new Payload("We are having a crypto party!"));
        jwsObject.sign(signer);
        String s = jwsObject.serialize();

        JWSObject parsed = JWSObject.parse(s);
        JWSVerifier verifier = new Ed25519Verifier(publicJWK);
        boolean valid = parsed.verify(verifier);
        System.out.println("ed25519 " + valid + " " + parsed.getPayload());
    }
}
