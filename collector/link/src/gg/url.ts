import { GG_KEY } from "./constant.ts";

// 구간
export const linkUrl = (start: number, end: number) => {
  return `https://openapi.gg.go.kr/TBBMSROUTESTATIONM?key=${GG_KEY}&type=json&pIndex=${start}&pSize=${end}`;
};

// 정류장
export const stationUrl = (start: number, end: number) => {
  return `https://openapi.gg.go.kr/TBBMSSTATIONM?key=${GG_KEY}&type=json&pIndex=${start}&pSize=${end}`;
};
