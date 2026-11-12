package org.cbench.depth3.libb;

import org.cbench.depth3.libc.Codec;

public final class Relay {
    private Relay() {
    }

    public static byte[] seal(byte[] data) throws Exception {
        return Codec.mac(data);
    }
}
