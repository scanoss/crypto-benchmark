package bench;

import org.jasypt.encryption.pbe.StandardPBEStringEncryptor;
import org.jasypt.iv.RandomIvGenerator;

public class PbeAes256 {
    public static void main(String[] args) throws Exception {
        StandardPBEStringEncryptor encryptor = new StandardPBEStringEncryptor();
        encryptor.setPassword("jasypt-master-password");
        encryptor.setAlgorithm("PBEWITHHMACSHA512ANDAES_256");
        encryptor.setIvGenerator(new RandomIvGenerator());
        String encrypted = encryptor.encrypt("database password");
        String decrypted = encryptor.decrypt(encrypted);
        System.out.println(encrypted);
        System.out.println(decrypted);
    }
}
