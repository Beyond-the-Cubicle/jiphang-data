export const linkUrl = (start: number, end: number) => {
  return `https://openapi.gg.go.kr/TBBMSROUTESTATIONM?key=${process.env.GG_KEY}&type=json&pIndex=${start}&pSize=${end}`;
};
