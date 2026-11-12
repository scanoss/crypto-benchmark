package bench;

import com.nimbusds.jose.jwk.Curve;
import com.nimbusds.jose.jwk.ECKey;
import com.nimbusds.jose.jwk.JWK;
import com.nimbusds.jose.jwk.KeyUse;
import com.nimbusds.jose.jwk.gen.ECKeyGenerator;
import com.nimbusds.jose.util.Base64URL;

public class JwkJson {
    public static void main(String[] args) throws Exception {
        ECKey ecJWK = new ECKeyGenerator(Curve.P_256)
                .keyUse(KeyUse.SIGNATURE)
                .keyID("789")
                .generate();
        String privateJson = ecJWK.toJSONString();
        String publicJson = ecJWK.toPublicJWK().toJSONString();

        JWK parsed = JWK.parse(privateJson);
        ECKey parsedPublic = ECKey.parse(publicJson);
        Base64URL thumbprint = parsedPublic.computeThumbprint();
        System.out.println("jwk " + parsed.isPrivate() + " " + parsedPublic.getCurve() + " " + thumbprint);
    }
}
