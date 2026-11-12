package org.cbench.depth3.libc;

import java.nio.charset.StandardCharsets;
import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;

public final class Codec {
    private static final byte[] KEY = "cbench-suite-b-key".getBytes(StandardCharsets.UTF_8);

    private Codec() {
    }

    public static byte[] mac(byte[] data) throws Exception {
        Mac mac = Mac.getInstance("HmacSHA256");
        mac.init(new SecretKeySpec(KEY, "HmacSHA256"));
        return mac.doFinal(data);
    }
}
