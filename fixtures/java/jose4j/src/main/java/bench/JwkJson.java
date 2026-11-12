package bench;

import org.jose4j.jwk.EcJwkGenerator;
import org.jose4j.jwk.EllipticCurveJsonWebKey;
import org.jose4j.jwk.JsonWebKey;
import org.jose4j.jwk.JsonWebKey.OutputControlLevel;
import org.jose4j.jwk.PublicJsonWebKey;
import org.jose4j.keys.EllipticCurves;
import org.jose4j.lang.HashUtil;

public class JwkJson {
    public static void main(String[] args) throws Exception {
        EllipticCurveJsonWebKey ecJsonWebKey = EcJwkGenerator.generateJwk(EllipticCurves.P256);
        ecJsonWebKey.setKeyId("k3");
        String privateJson = ecJsonWebKey.toJson(OutputControlLevel.INCLUDE_PRIVATE);
        String publicJson = ecJsonWebKey.toJson(OutputControlLevel.PUBLIC_ONLY);

        PublicJsonWebKey parsed = PublicJsonWebKey.Factory.newPublicJwk(privateJson);
        JsonWebKey parsedPublic = JsonWebKey.Factory.newJwk(publicJson);
        String thumbprint = parsedPublic.calculateBase64urlEncodedThumbprint(HashUtil.SHA_256);
        System.out.println("jwk " + (parsed.getPrivateKey() != null) + " " + parsedPublic.getKeyType() + " " + thumbprint);
    }
}
