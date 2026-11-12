package bench;

import org.jasypt.encryption.pbe.StandardPBEStringEncryptor;

public class PbeDefault {
    public static void main(String[] args) throws Exception {
        StandardPBEStringEncryptor encryptor = new StandardPBEStringEncryptor();
        encryptor.setPassword("jasypt-master-password");
        String encrypted = encryptor.encrypt("database password");
        String decrypted = encryptor.decrypt(encrypted);
        System.out.println(encrypted);
        System.out.println(decrypted);
    }
}
