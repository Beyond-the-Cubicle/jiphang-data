export const ErrorCode = {
  "INFO-000": "정상 처리되었습니다",
  "ERROR-300": "필수 값이 누락되어 있습니다.\n요청인자를 참고 하십시오.",
  "INFO-100":
    "인증키가 유효하지 않습니다.\n 인증키가 없는 경우, 열린 데이터 광장 홈페이지에서 인증키를 신청하십시오.",
  "ERROR-301":
    "파일타입 값이 누락 혹은 유효하지 않습니다.\n요청인자 중 TYPE을 확인하십시오.",
  "ERROR-310":
    "해당하는 서비스를 찾을 수 없습니다.\n요청인자 중 SERVICE를 확인하십시오.",
  "ERROR-331":
    "요청시작위치 값을 확인하십시오.\n요청인자 중 START_INDEX를 확인하십시오.",
  "ERROR-332":
    "요청종료위치 값을 확인하십시오.\n요청인자 중 END_INDEX를 확인하십시오.",
  "ERROR-333":
    "요청위치 값의 타입이 유효하지 않습니다.\n요청위치 값은 정수를 입력하세요.",
  "ERROR-334":
    "요청종료위치 보다 요청시작위치가 더 큽니다.\n요청시작조회건수는 정수를 입력하세요.",
  "ERROR-335":
    "샘플데이터(샘플키:sample) 는 한번에 최대 5건을 넘을 수 없습니다.\n요청시작위치와 요청종료위치 값은 1 ~ 5 사이만 가능합니다.",
  "ERROR-336":
    "데이터요청은 한번에 최대 1000건을 넘을 수 없습니다.\n요청종료위치에서 요청시작위치를 뺀 값이 1000을 넘지 않도록 수정하세요.",
  "ERROR-500":
    "서버 오류입니다.\n지속적으로 발생시 열린 데이터 광장으로 문의(Q&A) 바랍니다.",
  "ERROR-600":
    "데이터베이스 연결 오류입니다.\n지속적으로 발생시 열린 데이터 광장으로 문의(Q&A) 바랍니다.",
  "ERROR-601":
    "SQL 문장 오류 입니다.\n지속적으로 발생시 열린 데이터 광장으로 문의(Q&A) 바랍니다.",
  "INFO-200": "해당하는 데이터가 없습니다.",
} as const;
export type ErrorCode = keyof typeof ErrorCode;
export type ErrorMessage = (typeof ErrorCode)[keyof typeof ErrorCode];

//# region 노선-정류장 마스터 정보
export interface SeoulLinkResponse {
  masterRouteNode: {
    list_total_count: number;
    RESULT: {
      CODE: ErrorCode;
      MESSAGE: ErrorMessage;
    };
    row: SeoulLinkData[];
  };
}

export interface SeoulLinkData {
  // 노선 ID
  ROUTE_ID: string;
  // 정류장 ID
  STTN_ID: string;
  // 링크 구간거리 (m)
  STTN_DSTNC_MTR: number;
  // 정류장 순서
  STTN_ORD: number;
}
//#endregion
