package org.cbench.handlerspring.app;

import java.nio.charset.StandardCharsets;
import java.util.HexFormat;
import org.cbench.handlerspring.liba.Fingerprint;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class FingerprintController {
    @GetMapping("/fingerprint")
    public String fingerprint(@RequestParam(defaultValue = "") String data) throws Exception {
        return HexFormat.of().formatHex(Fingerprint.of(data.getBytes(StandardCharsets.UTF_8)));
    }
}
