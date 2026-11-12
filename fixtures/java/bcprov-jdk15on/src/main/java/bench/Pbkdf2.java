package bench;

import java.nio.charset.StandardCharsets;
import java.security.SecureRandom;
import org.bouncycastle.crypto.PBEParametersGenerator;
import org.bouncycastle.crypto.digests.SHA256Digest;
import org.bouncycastle.crypto.generators.PKCS5S2ParametersGenerator;
import org.bouncycastle.crypto.params.KeyParameter;
import org.bouncycastle.util.encoders.Hex;

public class Pbkdf2 {
    public static void main(String[] args) {
        char[] password = "correct horse battery staple".toCharArray();
        byte[] salt = new byte[16];
        new SecureRandom().nextBytes(salt);

        PKCS5S2ParametersGenerator generator = new PKCS5S2ParametersGenerator(new SHA256Digest());
        generator.init(PBEParametersGenerator.PKCS5PasswordToUTF8Bytes(password), salt, 600000);
        KeyParameter derived = (KeyParameter) generator.generateDerivedParameters(256);

        System.out.println(Hex.toHexString(derived.getKey()));
    }
}
