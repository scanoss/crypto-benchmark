package bench;

import org.mindrot.jbcrypt.BCrypt;

public class HashPassword {
    public static void main(String[] args) throws Exception {
        String password = "correct horse battery staple";
        String hashed = BCrypt.hashpw(password, BCrypt.gensalt(12));
        System.out.println(hashed);
        System.out.println(BCrypt.checkpw(password, hashed));
        System.out.println(BCrypt.checkpw("wrong password", hashed));
    }
}
