const SEOUL_BASE_URL = "http://openapi.seoul.go.kr:8088";

// 노선-정류장
export const linkUrl = (start: number, end: number) => {
  return `${SEOUL_BASE_URL}/${process.env.SEOUL_KEY}/json/masterRouteNode/${start}/${end}`;
};
