package bench;

import com.jcraft.jsch.JSch;
import com.jcraft.jsch.KeyPair;
import java.io.ByteArrayOutputStream;

public class GenerateKeyPair {
    public static void main(String[] args) throws Exception {
        JSch jsch = new JSch();
        KeyPair kpair = KeyPair.genKeyPair(jsch, KeyPair.RSA, 3072);
        ByteArrayOutputStream privateKey = new ByteArrayOutputStream();
        kpair.writePrivateKey(privateKey);
        ByteArrayOutputStream publicKey = new ByteArrayOutputStream();
        kpair.writePublicKey(publicKey, "bench@example.com");
        System.out.println("Finger print: " + kpair.getFingerPrint());
        System.out.println(publicKey.toString());
        kpair.dispose();
    }
}
