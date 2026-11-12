package org.cbench.deadfn.libb;

import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;

public final class Codec {
    private static final byte[] KEY = "cbench-suite-b-key".getBytes(StandardCharsets.UTF_8);

    private Codec() {
    }

    public static byte[] sign(byte[] data) throws Exception {
        Mac mac = Mac.getInstance("HmacSHA256");
        mac.init(new SecretKeySpec(KEY, "HmacSHA256"));
        return mac.doFinal(data);
    }

    public static int checksum(byte[] data) {
        int c = 0;
        for (byte b : data) {
            c ^= b;
        }
        return c;
    }

    public static byte[] legacyDigest(byte[] data) throws Exception {
        MessageDigest md = MessageDigest.getInstance("SHA-1");
        return md.digest(data);
    }
}
