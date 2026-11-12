package bench;

import com.jcraft.jsch.JSch;
import com.jcraft.jsch.Session;

public class SessionConfig {
    public static void main(String[] args) throws Exception {
        JSch jsch = new JSch();
        Session session = jsch.getSession("bench", "ssh.example.com", 22);
        session.setConfig("StrictHostKeyChecking", "yes");
        session.setConfig("cipher.c2s", "aes256-ctr,aes128-ctr");
        session.setConfig("cipher.s2c", "aes256-ctr,aes128-ctr");
        session.setConfig("kex", "ecdh-sha2-nistp256,diffie-hellman-group14-sha1");
        session.setConfig("server_host_key", "ecdsa-sha2-nistp256,ssh-rsa");
        System.out.println(session.getHost() + " " + session.getConfig("cipher.c2s"));
    }
}
