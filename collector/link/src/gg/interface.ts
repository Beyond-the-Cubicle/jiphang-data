export interface GGLinkResponse {
  TBBMSROUTESTATIONM: [
    {
      head: [
        {
          list_total_count: number;
        },
        {
          RESULT: {
            CODE: string;
            MESSAGE: string;
          };
        },
        {
          api_version: string;
        },
      ];
    },
    {
      row: GGLinkData[];
    },
  ];
}

export interface GGLinkData {
  // 노선 아이디
  ROUTE_ID: string;
  // 정류장 순서
  STTN_ORDR: number;
  // 정류장 아이디
  STTN_ID: string;
  // GIS 거리
  GIS_DSTN: number;
  // 누적 거리
  ACCMLT_DSTN: number;
  // 실제 거리
  REAL_DSTN: number;
  // 확정 거리
  DCSN_DSTN: number;
  // 진행구분 코드
  PROGRS_DIV_CD: string;
  // 등록아이디
  REGIST_ID: string;
  // 등록일자
  REGIST_DE: string;
  // 사용구분
  USE_DIV: string | null;
  // 벽지노선 유무
  UNWEL_HNO_STATN_ROUTE_EXTNO: string;
  // 진행구분 코드명
  PROGRS_DIV_CD_NM: string;
  // 벽지노선 유무명
  USE_DIV_NM: string;
}
