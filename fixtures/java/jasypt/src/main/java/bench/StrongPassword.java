package bench;

import org.jasypt.util.password.StrongPasswordEncryptor;

public class StrongPassword {
    public static void main(String[] args) throws Exception {
        StrongPasswordEncryptor passwordEncryptor = new StrongPasswordEncryptor();
        String encrypted = passwordEncryptor.encryptPassword("correct horse battery staple");
        System.out.println(encrypted);
        System.out.println(passwordEncryptor.checkPassword("correct horse battery staple", encrypted));
    }
}
