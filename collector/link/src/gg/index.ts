import { PrismaClient } from "@prisma/client";
import { GGLinkResponse } from "./interface";
import { request } from "../request";
import { linkUrl } from "./url";

const prisma = new PrismaClient();

export async function getGGLink() {
  const start = 1;
  let end = 1000;
  const step = 1000;

  console.log("===== 경기도 구간정보 수집 시작 =====");

  // 구간 정보 0417기준 483페이지
  for (let i = start, page = 1; i < end; i += step, page++) {
    console.log(`page: ${page}`);
    const result = await request<GGLinkResponse>(linkUrl, page, step);

    // next step
    const totalCount = result.TBBMSROUTESTATIONM[0].head[0].list_total_count;
    end = end + step > totalCount ? totalCount : end + step;

    const routes = result.TBBMSROUTESTATIONM[1].row;

    await prisma.gyeonggiLink.createMany({
      data: routes.map((route) => ({
        route_id: route.ROUTE_ID,
        sttn_ordr: route.STTN_ORDR,
        sttn_id: route.STTN_ID,
        gis_dstn: route.GIS_DSTN,
        accmlt_dstn: route.ACCMLT_DSTN,
        real_dstn: route.REAL_DSTN,
        dcsn_dstn: route.DCSN_DSTN,
        progrs_div_cd: route.PROGRS_DIV_CD,
        use_div: route.USE_DIV,
        unwel_hno_statn_route_extno: route.UNWEL_HNO_STATN_ROUTE_EXTNO,
        progrs_div_cd_nm: route.PROGRS_DIV_CD_NM,
        use_div_nm: route.USE_DIV_NM,
      })),
      skipDuplicates: true,
    });
  }
  console.log("===== 경기도 구간정보 수집 종료 =====");
}
