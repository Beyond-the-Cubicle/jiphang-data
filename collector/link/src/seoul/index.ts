import { linkUrl, request, routeUrl, stationUrl } from "./request.ts";
import { PrismaClient } from "@prisma/client";
import {
  SeoulLinkResponse,
  SeoulRouteResponse,
  SeoulStationResponse,
} from "./interface.ts";

const prisma = new PrismaClient();

const start = 1;
let end = 1000;
const step = 1000;

console.log("===== route start =====");
// 노선 정보
for (let i = start, page = 1; i < end; i += step, page++) {
  console.log(`page: ${page}`);
  const result = await request<SeoulRouteResponse>(routeUrl, i, end);

  // next step
  const totalCount = result.tbisMasterRoute.list_total_count;
  end = end + step > totalCount ? totalCount : end + step;

  const routes = result.tbisMasterRoute.row;

  await prisma.seoulRoute.createMany({
    data: routes.map((route) => ({
      route_id: route.ROUTE_ID,
      route_nm: route.ROUTE_NM,
      route_type: route.ROUTE_TYPE,
      dstnc: route.DSTNC,
    })),
    skipDuplicates: true,
  });
}
console.log("===== route end =====");

end = 1000;

console.log("===== station start =====");
// 정류장 정보
for (let i = start, page = 1; i < end; i += step, page++) {
  console.log(`page: ${page}`);
  const result = await request<SeoulStationResponse>(stationUrl, i, end);

  // next step
  const totalCount = result.tbisMasterStation.list_total_count;
  end = end + step > totalCount ? totalCount : end + step;

  const routes = result.tbisMasterStation.row;

  await prisma.seoulStation.createMany({
    data: routes.map((route) => ({
      sttn_id: route.STTN_ID,
      sttn_nm: route.STTN_NM,
      sttn_no: route.STTN_NO,
      sttn_type: route.STTN_TYPE,
      crdnt_x: route.CRDNT_X,
      crdnt_y: route.CRDNT_Y,
      businfo_fclt_instl_yn: route.BUSINFO_FCLT_INSTL_YN,
    })),
    skipDuplicates: true,
  });
}
console.log("===== station end =====");

end = 1000;

console.log("===== link start =====");
// 구간 정보
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
console.log("===== link end =====");
