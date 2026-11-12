package org.cbench.deadfn.liba;

import org.cbench.deadfn.libb.Codec;

public final class Api {
    private Api() {
    }

    public static byte[] sign(byte[] data) throws Exception {
        return Codec.sign(data);
    }

    public static int checksum(byte[] data) {
        return Codec.checksum(data);
    }
}
