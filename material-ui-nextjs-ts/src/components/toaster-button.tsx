'use client'

import React from 'react';
import {useTransport} from "@connectrpc/connect-query";
import {createPromiseClient} from "@connectrpc/connect";
import {GreetService} from "../../gen/example/v1/greet_connect";
import toast from "react-hot-toast";
import {Button} from "@mui/material";


interface Props {

}

const ToasterButton: React.FunctionComponent<Props> = function ToasterButton() {
  const t = useTransport();
  const client = createPromiseClient(GreetService, t)

  return (
    <Button
      data-testid={"toaster-btn"}
      variant="contained"
      onClick={async () => {
        try {
          const response = await client.greet({name: "World"})
          toast(response.greeting, {
            ariaProps: {
              role: 'status',
              "aria-live": "polite",
            }
          })
        } catch (e: any) {
          toast.error(e.message)
        }
      }}
    >Hello World
    </Button>
  );
};

export default ToasterButton;
