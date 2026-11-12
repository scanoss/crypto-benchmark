package bench;

import java.util.Arrays;
import javax.net.ssl.SSLContext;
import javax.net.ssl.SSLEngine;

public class TlsContext {
    public static void main(String[] args) throws Exception {
        SSLContext context = SSLContext.getInstance("TLSv1.3");
        context.init(null, null, null);

        SSLEngine engine = context.createSSLEngine("example.com", 443);
        engine.setUseClientMode(true);
        engine.setEnabledProtocols(new String[] {"TLSv1.3"});

        System.out.println(context.getProtocol() + " " + Arrays.toString(engine.getEnabledProtocols()));
    }
}
