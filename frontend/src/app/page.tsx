import { createPromiseClient } from "@connectrpc/connect";
// import { createConnectTransport } from "@connectrpc/connect-web";
import styles from "./page.module.css";
import { GreetService } from "../../gen/proto/v1/greet_connect";
import { createConnectTransport } from "@connectrpc/connect-node";
import { GreetResponse } from "../../gen/proto/v1/greet_pb";

async function getGreet(): Promise<GreetResponse> {
  const response = await fetch(
    `http://localhost:8080/${GreetService.typeName}/${GreetService.methods.greet.name}`,
    {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name: "Test",
      }),
    }
  );
  return response.json();
}

export default async function Home() {
  // connect-web
  // const transport = createConnectTransport({
  //   // Note: you cannot use a relative path like `/api` here because SSR requires absolute URLs.
  //   baseUrl: "http://localhost:8080",
  // });

  // connect-node
  // const transport = createConnectTransport({
  //   httpVersion: "2",
  //   baseUrl: "http://localhost:8080",
  // });
  // const client = createPromiseClient(GreetService, transport);
  // const response = await client.greet({ name: "World" });

  // fetch
  const response = await getGreet();
  // console.log(response);

  return <main className={styles.main}>{response.greeting}</main>;
}
