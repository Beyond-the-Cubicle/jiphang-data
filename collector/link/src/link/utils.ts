import { Prisma } from "@prisma/client";
import { IBasicLink } from "./interface";

//#region 서울 구간 데이터셋 정제
export const makeSeoulDataset = (
  links: Prisma.SeoulLinkGetPayload<{}>[],
  speed: Prisma.link_speed_averageGetPayload<{}>,
) => {
  const linkData: IBasicLink[] = [];

  for (let i = 0; i < links.length; i++) {
    const start = links[i];
    let end;
    let totalDistance = 0;

    for (let j = i + 1; j < links.length; j++) {
      end = links[j];
      totalDistance += end.sttn_dstnc_mtr;

      const data: IBasicLink = {
        routeId: Number(start.route_id),
        startStationId: Number(start.sttn_id),
        endStationId: Number(end.sttn_id),
        tripTime: Math.round(totalDistance / speed.weekdayMs),
        tripDistance: totalDistance,
        stationOrder: start.sttn_ordr,
      };

      linkData.push(data);
    }

    totalDistance = 0;
  }

  return linkData;
};
//#endregion

export const makeGGDataset = (
  links: Prisma.GyeonggiLinkGetPayload<{}>[],
  speed: Prisma.link_speed_averageGetPayload<{}>,
) => {
  const linkData: IBasicLink[] = [];

  for (let i = 0; i < links.length; i++) {
    const start = links[i];
    let end;
    let totalDistance = 0;

    for (let j = i + 1; j < links.length; j++) {
      end = links[j];
      if (end.dcsn_dstn === null) continue;
      totalDistance += end.dcsn_dstn;

      const data: IBasicLink = {
        routeId: Number(start.route_id),
        startStationId: Number(start.sttn_id),
        endStationId: Number(end.sttn_id),
        tripTime: Math.round(totalDistance / speed.weekdayMs),
        tripDistance: totalDistance,
        stationOrder: start.sttn_ordr,
      };

      linkData.push(data);
    }

    totalDistance = 0;
  }

  return linkData;
};
