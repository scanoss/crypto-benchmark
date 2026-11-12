package bench;

import org.jasypt.util.text.BasicTextEncryptor;

public class TextEncryption {
    public static void main(String[] args) throws Exception {
        BasicTextEncryptor textEncryptor = new BasicTextEncryptor();
        textEncryptor.setPassword("jasypt-master-password");
        String encrypted = textEncryptor.encrypt("api token");
        String decrypted = textEncryptor.decrypt(encrypted);
        System.out.println(encrypted);
        System.out.println(decrypted);
    }
}
