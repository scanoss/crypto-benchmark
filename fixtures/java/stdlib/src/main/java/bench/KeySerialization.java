package bench;

import java.security.KeyFactory;
import java.security.KeyPair;
import java.security.KeyPairGenerator;
import java.security.PrivateKey;
import java.security.PublicKey;
import java.security.spec.ECGenParameterSpec;
import java.security.spec.PKCS8EncodedKeySpec;
import java.security.spec.X509EncodedKeySpec;

public class KeySerialization {
    public static void main(String[] args) throws Exception {
        KeyPairGenerator keyPairGen = KeyPairGenerator.getInstance("EC");
        keyPairGen.initialize(new ECGenParameterSpec("secp256r1"));
        KeyPair keyPair = keyPairGen.generateKeyPair();

        byte[] publicBytes = keyPair.getPublic().getEncoded();
        byte[] privateBytes = keyPair.getPrivate().getEncoded();

        KeyFactory keyFactory = KeyFactory.getInstance("EC");
        PublicKey publicKey = keyFactory.generatePublic(new X509EncodedKeySpec(publicBytes));
        PrivateKey privateKey = keyFactory.generatePrivate(new PKCS8EncodedKeySpec(privateBytes));

        System.out.println("public: " + publicKey.equals(keyPair.getPublic()));
        System.out.println("private: " + privateKey.equals(keyPair.getPrivate()));
    }
}
