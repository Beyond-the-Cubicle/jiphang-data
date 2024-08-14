import { collection } from "./collection";
import { processing } from "./processing";

async function main() {
  await collection();
  await processing();
}

main();
