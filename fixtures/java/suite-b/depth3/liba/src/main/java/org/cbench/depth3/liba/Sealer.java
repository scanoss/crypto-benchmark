package org.cbench.depth3.liba;

import org.cbench.depth3.libb.Relay;

public final class Sealer {
    private Sealer() {
    }

    public static byte[] seal(byte[] data) throws Exception {
        return Relay.seal(data);
    }
}
