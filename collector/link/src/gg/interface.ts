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

export interface GGStationResponse {
  TBBMSSTATIONM: [
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
      row: GGStationData[];
    },
  ];
}

export interface GGStationData {
  // 정류장 아이디
  STTN_ID: string;
  // 정류장 명칭
  STTN_NM: string;
  // x 좌표
  X_CRDNT: number;
  // Y 좌표
  Y_CRDNT: number;
  // GPS X 좌표
  GPS_X_CRDNT: number | null;
  // GPS Y 좌표
  GPS_Y_CRDNT: number | null;
  // 링크 ID
  RINK_ID: string;
  // 정류장 유형
  STTN_TYPE: string;
  // 환승정류장 유무
  TRANSIT_STTN_EXTNO: string | null;
  // 중앙차로 여부
  CNTR_CARTRK_YN: string;
  // 정류장 영문명
  STTN_ENG_NM: string;
  // ARS ID
  ARS_ID: string;
  // 기관코드
  INST_CD: string;
  // 데이터 표출 유무
  DATA_EXPRS_EXTNO: string;
  // 등록 아이디
  REGIST_ID: string;
  // 등록 일자
  REGIST_DE: string;
  // 비고
  RM: string | null;
  // 표지판 유형
  SIGNPOST_TYPE: string | null;
  // 행정동코드
  ADMINIST_DONG_CD: string;
  // 권역코드
  VOLM_STATN_CD: string;
  // 사용구분
  USE_DIV: string;
  // 중국어명
  STTN_CHN_NM: string;
  // 일본어명
  STTN_JPNLANG_NM: string;
  // 베트남명
  STTN_VIETNAM_NM: string;
  // DRT 유무
  DRT_EXTNO: string;
  // 정류장 유형명
  STATION_TP_NM: string;
  // 환승역 타입명
  CHNG_STATION_YN_NM: string | null;
  // 표지판 유형명
  MARK_TYPE_NM: string | null;
}
