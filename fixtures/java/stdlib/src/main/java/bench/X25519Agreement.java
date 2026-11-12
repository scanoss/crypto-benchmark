package bench;

import java.security.KeyPair;
import java.security.KeyPairGenerator;
import java.security.MessageDigest;
import javax.crypto.KeyAgreement;

public class X25519Agreement {
    public static void main(String[] args) throws Exception {
        KeyPairGenerator keyPairGen = KeyPairGenerator.getInstance("X25519");
        KeyPair alice = keyPairGen.generateKeyPair();
        KeyPair bob = keyPairGen.generateKeyPair();

        KeyAgreement aliceAgreement = KeyAgreement.getInstance("X25519");
        aliceAgreement.init(alice.getPrivate());
        aliceAgreement.doPhase(bob.getPublic(), true);
        byte[] aliceSecret = aliceAgreement.generateSecret();

        KeyAgreement bobAgreement = KeyAgreement.getInstance("X25519");
        bobAgreement.init(bob.getPrivate());
        bobAgreement.doPhase(alice.getPublic(), true);
        byte[] bobSecret = bobAgreement.generateSecret();

        System.out.println("agreed: " + MessageDigest.isEqual(aliceSecret, bobSecret));
    }
}
