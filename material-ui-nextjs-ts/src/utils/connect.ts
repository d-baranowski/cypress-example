import {createConnectTransport} from "@connectrpc/connect-web";
import {createRegistry} from "@bufbuild/protobuf";

export const ConnectTransport = createConnectTransport({
    baseUrl: 'http://localhost:8080',
    jsonOptions: {
      typeRegistry: createRegistry(
        //Account,
      )
    },
  })
;
