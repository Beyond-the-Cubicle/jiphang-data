import { PrismaClient } from "@prisma/client";
import { GGLinkResponse, GGStationResponse } from "./interface.ts";
import { request } from "../request.ts";
import { linkUrl, stationUrl } from "./url.ts";

const prisma = new PrismaClient();

const start = 1;
let end = 1000;
const step = 1000;

console.log("===== link start =====");
// 구간 정보
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
      // regist_id: route.REGIST_ID,
      // regist_de: route.REGIST_DE,
      use_div: route.USE_DIV,
      unwel_hno_statn_route_extno: route.UNWEL_HNO_STATN_ROUTE_EXTNO,
      progrs_div_cd_nm: route.PROGRS_DIV_CD_NM,
      use_div_nm: route.USE_DIV_NM,
    })),
    skipDuplicates: true,
  });
}
console.log("===== link end =====");

end = 1000;

console.log("===== station start =====");
// 정류장 정보
for (let i = start, page = 1; i < end; i += step, page++) {
  console.log(`page: ${page}`);
  const result = await request<GGStationResponse>(stationUrl, page, step);

  // next step
  const totalCount = result.TBBMSSTATIONM[0].head[0].list_total_count;
  end = end + step > totalCount ? totalCount : end + step;

  const station = result.TBBMSSTATIONM[1].row;

  await prisma.gyeonggiStation.createMany({
    data: station.map((route) => ({
      sttn_id: route.STTN_ID,
      sttn_nm: route.STTN_NM,
      x_crdnt: route.X_CRDNT,
      y_crdnt: route.Y_CRDNT,
      gps_x_crdnt: route.GPS_X_CRDNT,
      gps_y_crdnt: route.GPS_Y_CRDNT,
      rink_id: route.RINK_ID,
      sttn_type: route.STTN_TYPE,
      transit_sttn_extno: route.TRANSIT_STTN_EXTNO,
      cntr_cartrk_yn: route.CNTR_CARTRK_YN,
      sttn_eng_nm: route.STTN_ENG_NM,
      ars_id: route.ARS_ID,
      inst_cd: route.INST_CD,
      data_exprs_extno: route.DATA_EXPRS_EXTNO,
      // regist_id: route.REGIST_ID,
      // regist_de: route.REGIST_DE,
      // rm: route.RM,
      signpost_type: route.SIGNPOST_TYPE,
      administ_dong_cd: route.ADMINIST_DONG_CD,
      volm_statn_cd: route.VOLM_STATN_CD,
      use_div: route.USE_DIV,
      // sttn_chn_nm: route.STTN_CHN_NM,
      // sttn_jpnlang_nm: route.STTN_JPNLANG_NM,
      // sttn_vietnam_nm: route.STTN_VIETNAM_NM,
      drt_extno: route.DRT_EXTNO,
      station_tp_nm: route.STATION_TP_NM,
      chng_station_yn_nm: route.CHNG_STATION_YN_NM,
      mark_type_nm: route.MARK_TYPE_NM,
    })),
    skipDuplicates: true,
  });
}

console.log("===== station end =====");
