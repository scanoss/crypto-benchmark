package org.cbench.direct.liba;

import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;

public final class Fingerprint {
    private Fingerprint() {
    }

    public static byte[] digest(byte[] data) throws NoSuchAlgorithmException {
        MessageDigest md = MessageDigest.getInstance("SHA-256");
        return md.digest(data);
    }
}
