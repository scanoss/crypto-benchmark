package bench;

import at.favre.lib.crypto.bcrypt.BCrypt;

public class HashPassword {
    public static void main(String[] args) throws Exception {
        String password = "correct horse battery staple";
        String hash = BCrypt.withDefaults().hashToString(12, password.toCharArray());
        System.out.println(hash);
        BCrypt.Result result = BCrypt.verifyer().verify(password.toCharArray(), hash);
        System.out.println(result.verified);
    }
}
