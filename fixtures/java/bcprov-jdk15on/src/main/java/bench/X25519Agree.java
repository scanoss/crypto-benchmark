package bench;

import java.security.SecureRandom;
import java.util.Arrays;
import org.bouncycastle.crypto.AsymmetricCipherKeyPair;
import org.bouncycastle.crypto.agreement.X25519Agreement;
import org.bouncycastle.crypto.generators.X25519KeyPairGenerator;
import org.bouncycastle.crypto.params.X25519KeyGenerationParameters;

public class X25519Agree {
    public static void main(String[] args) {
        SecureRandom random = new SecureRandom();
        X25519KeyPairGenerator generator = new X25519KeyPairGenerator();
        generator.init(new X25519KeyGenerationParameters(random));
        AsymmetricCipherKeyPair alice = generator.generateKeyPair();
        AsymmetricCipherKeyPair bob = generator.generateKeyPair();

        X25519Agreement aliceAgreement = new X25519Agreement();
        aliceAgreement.init(alice.getPrivate());
        byte[] aliceSecret = new byte[aliceAgreement.getAgreementSize()];
        aliceAgreement.calculateAgreement(bob.getPublic(), aliceSecret, 0);

        X25519Agreement bobAgreement = new X25519Agreement();
        bobAgreement.init(bob.getPrivate());
        byte[] bobSecret = new byte[bobAgreement.getAgreementSize()];
        bobAgreement.calculateAgreement(alice.getPublic(), bobSecret, 0);

        System.out.println("shared secrets match: " + Arrays.equals(aliceSecret, bobSecret));
    }
}
