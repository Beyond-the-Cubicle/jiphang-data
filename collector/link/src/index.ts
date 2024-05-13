import dotenv from "dotenv";
import { getGGLink } from "./gg";
import { getSeoulLink } from "./seoul";
import { makeStandardLinkData } from "./link";
dotenv.config();

async function main() {
  // 수집
  await getGGLink();
  await getSeoulLink();

  // 가공
  await makeStandardLinkData();
}

main();
