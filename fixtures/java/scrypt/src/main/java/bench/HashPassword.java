package bench;

import com.lambdaworks.crypto.SCryptUtil;

public class HashPassword {
    public static void main(String[] args) throws Exception {
        String password = "correct horse battery staple";
        String hashed = SCryptUtil.scrypt(password, 16384, 8, 1);
        System.out.println(hashed);
        System.out.println(SCryptUtil.check(password, hashed));
        System.out.println(SCryptUtil.check("wrong password", hashed));
    }
}
