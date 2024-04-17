import { PrismaClient } from "@prisma/client";
import { SeoulLinkResponse } from "./interface";
import { request } from "../request";
import { linkUrl } from "./url";

const prisma = new PrismaClient();

export async function getSeoulLink() {
  const start = 1;
  let end = 1000;
  const step = 1000;

  console.log("===== 서울 구간정보 수집 시작 =====");

  // 구간 정보 0417기준 149페이지
  for (let i = start, page = 1; i < end; i += step, page++) {
    console.log(`page: ${page}`);
    const result = await request<SeoulLinkResponse>(linkUrl, i, end);

    // next step
    const totalCount = result.masterRouteNode.list_total_count;
    end = end + step > totalCount ? totalCount : end + step;

    const routes = result.masterRouteNode.row;

    await prisma.seoulLink.createMany({
      data: routes.map((route) => ({
        sttn_id: route.STTN_ID,
        route_id: route.ROUTE_ID,
        sttn_ordr: route.STTN_ORD,
        sttn_dstnc_mtr: route.STTN_DSTNC_MTR,
      })),
      skipDuplicates: true,
    });
  }
  console.log("===== 서울 구간정보 수집 종료 =====");
}
