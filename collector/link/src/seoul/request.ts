import { SEOUL_BASE_URL, SEOUL_KEY } from "./constant.ts";

export const request = async <T>(
  url: (start: number, end: number) => string,
  start: number,
  end: number,
) => {
  const resp = await fetch(url(start, end), {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });

  return (await resp.json()) as T;
};

// 노선
export const routeUrl = (start: number, end: number) => {
  return `${SEOUL_BASE_URL}/${SEOUL_KEY}/json/tbisMasterRoute/${start}/${end}`;
};

// 정류장
export const stationUrl = (start: number, end: number) => {
  return `${SEOUL_BASE_URL}/${SEOUL_KEY}/json/tbisMasterStation/${start}/${end}`;
};

// 노선-정류장
export const linkUrl = (start: number, end: number) => {
  return `${SEOUL_BASE_URL}/${SEOUL_KEY}/json/masterRouteNode/${start}/${end}`;
};
