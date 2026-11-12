package bench;

import at.favre.lib.crypto.bcrypt.BCrypt;

public class HashVersion2y {
    public static void main(String[] args) throws Exception {
        String password = "correct horse battery staple";
        String hash = BCrypt.with(BCrypt.Version.VERSION_2Y).hashToString(12, password.toCharArray());
        System.out.println(hash);
        BCrypt.Result result = BCrypt.verifyer(BCrypt.Version.VERSION_2Y).verify(password.toCharArray(), hash);
        System.out.println(result.verified);
    }
}
