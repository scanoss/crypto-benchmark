package bench;

import java.security.KeyPair;
import java.security.spec.PKCS8EncodedKeySpec;
import java.security.spec.X509EncodedKeySpec;
import java.util.HexFormat;
import net.i2p.crypto.eddsa.EdDSAPrivateKey;
import net.i2p.crypto.eddsa.EdDSAPublicKey;
import net.i2p.crypto.eddsa.KeyPairGenerator;

public class KeyEncoding {
    public static void main(String[] args) throws Exception {
        KeyPairGenerator keyPairGen = new KeyPairGenerator();
        KeyPair keyPair = keyPairGen.generateKeyPair();

        byte[] privateBytes = keyPair.getPrivate().getEncoded();
        byte[] publicBytes = keyPair.getPublic().getEncoded();

        EdDSAPrivateKey privateKey = new EdDSAPrivateKey(new PKCS8EncodedKeySpec(privateBytes));
        EdDSAPublicKey publicKey = new EdDSAPublicKey(new X509EncodedKeySpec(publicBytes));

        System.out.println(HexFormat.of().formatHex(publicBytes));
        System.out.println("private: " + privateKey.equals(keyPair.getPrivate()));
        System.out.println("public: " + publicKey.equals(keyPair.getPublic()));
    }
}
