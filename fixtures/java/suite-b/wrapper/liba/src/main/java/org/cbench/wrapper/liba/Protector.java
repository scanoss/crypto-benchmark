package org.cbench.wrapper.liba;

import org.cbench.wrapper.libb.Encoder;

public final class Protector {
    private Protector() {
    }

    public static byte[] protect(byte[] data) throws Exception {
        return protectInner(data);
    }

    private static byte[] protectInner(byte[] data) throws Exception {
        return Encoder.encode(data);
    }
}
