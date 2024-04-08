import {Prisma} from "@prisma/client";
import {IBasicLink} from "../interface.ts";

//#region 서울 구간 데이터셋 정제
export const makeSeoulDataset = (links: Prisma.SeoulLinkGetPayload<{}>[], speed: Prisma.link_speed_averageGetPayload<{}>) => {
  const dataset: IBasicLink[] = [];

  const linkMap = makeSeoulLinkMap(links);

  for(const key of linkMap.keys()) {
    const map = linkMap.get(key);
    if (!map) continue;
    const linkData:IBasicLink[] = [];

    for (let i = 0 ; i < map.length; i++) {
      const start = map[i];
      let end;
      let totalDistance = 0;

      for (let j = i + 1; j < map.length; j++) {
        end = map[j];
        totalDistance += end.sttn_dstnc_mtr;

        const data: IBasicLink = {
          routeId: start.route_id,
          startStationId: start.sttn_id,
          endStationId: end.sttn_id,
          tripTime: totalDistance / speed.weekdayMs,
          tripDistance: totalDistance,
          stationOrder: start.sttn_ordr,
        };

        linkData.push(data);
      }

      totalDistance = 0;
    }

    dataset.push(...linkData);
  }

  return dataset;
}

const makeSeoulLinkMap = (links: Prisma.SeoulLinkGetPayload<{}>[]) => {
  const linkMap = new Map<string, Prisma.SeoulLinkGetPayload<{}>[]>();

  links.forEach(link => {
    const routeId = link.route_id;
    if (!linkMap.has(routeId)) {
      linkMap.set(routeId, []);
    }

    linkMap.get(routeId)!.push(link);
  });

  return linkMap;
}
//#endregion

export const makeGGDataset = (links: Prisma.GyeonggiLinkGetPayload<{}>[], speed: Prisma.link_speed_averageGetPayload<{}>) => {
  const dataset: IBasicLink[] = [];

  const linkMap = makeGGLinkMap(links);

  for(const key of linkMap.keys()) {
    const map = linkMap.get(key);
    if (!map) continue;
    const linkData:IBasicLink[] = [];

    for (let i = 0 ; i < map.length; i++) {
      const start = map[i];
      let end;
      let totalDistance = 0;

      for (let j = i + 1; j < map.length; j++) {
        end = map[j];
        if (end.dcsn_dstn === null) continue;
        totalDistance += end.dcsn_dstn;

        const data: IBasicLink = {
          routeId: start.route_id,
          startStationId: start.sttn_id,
          endStationId: end.sttn_id,
          tripTime: totalDistance / speed.weekdayMs,
          tripDistance: totalDistance,
          stationOrder: start.sttn_ordr,
        };

        linkData.push(data);
      }

      totalDistance = 0;
    }

    dataset.push(...linkData);
  }

  return dataset;
}

const makeGGLinkMap = (links: Prisma.GyeonggiLinkGetPayload<{}>[]) => {
  const linkMap = new Map<string, Prisma.GyeonggiLinkGetPayload<{}>[]>();

  links.forEach(link => {
    const routeId = link.route_id;
    if (!linkMap.has(routeId)) {
      linkMap.set(routeId, []);
    }

    linkMap.get(routeId)!.push(link);
  });

  return linkMap;
}


