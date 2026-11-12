package bench;

import org.jasypt.util.password.BasicPasswordEncryptor;

public class BasicPassword {
    public static void main(String[] args) throws Exception {
        BasicPasswordEncryptor passwordEncryptor = new BasicPasswordEncryptor();
        String encrypted = passwordEncryptor.encryptPassword("correct horse battery staple");
        System.out.println(encrypted);
        System.out.println(passwordEncryptor.checkPassword("correct horse battery staple", encrypted));
    }
}
