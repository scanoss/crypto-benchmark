package bench;

import org.jasypt.digest.StandardStringDigester;

public class StringDigest {
    public static void main(String[] args) throws Exception {
        StandardStringDigester digester = new StandardStringDigester();
        digester.setAlgorithm("SHA-256");
        digester.setIterations(50000);
        String digest = digester.digest("message to digest");
        System.out.println(digest);
        System.out.println(digester.matches("message to digest", digest));
    }
}
