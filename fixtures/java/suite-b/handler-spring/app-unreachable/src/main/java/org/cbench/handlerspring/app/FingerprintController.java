package org.cbench.handlerspring.app;

import java.nio.charset.StandardCharsets;
import java.util.HexFormat;
import org.cbench.handlerspring.liba.Fingerprint;

public class FingerprintController {
    public String fingerprint(String data) throws Exception {
        return HexFormat.of().formatHex(Fingerprint.of(data.getBytes(StandardCharsets.UTF_8)));
    }
}
