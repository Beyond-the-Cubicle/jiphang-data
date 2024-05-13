import { makeGGDataset, makeSeoulDataset } from "./utils";
import dotenv from "dotenv";
import { PrismaClient } from "@prisma/client";
import { ISpeed } from "./interface";

dotenv.config();

const prisma = new PrismaClient();

async function getAverageSpeed() {
  const host = "https://api.odcloud.kr";
  const path = "/api/15066795/v1/uddi:69bf5e71-10ce-4ecc-8df8-e39609bc0144";
  const page = 1;
  const perPage = 20;
  const serviceKey = process.env.DATA_GO_KR_KEY;

  const resp = await fetch(
    `${host}${path}?page=${page}&perPage=${perPage}&serviceKey=${serviceKey}`,
  );
  const result: {
    data: { 구분: string; 평일: string; 토요일: string; 일요일: string }[];
  } = await resp.json();

  const speeds = result.data.reduce((acc, cur) => {
    acc.push({
      year: 2022,
      location: cur.구분,
      weekdayKmh: parseFloat(cur.평일),
      weekdayMs: parseFloat(cur.평일) / 3.6,
      saturdayKmh: parseFloat(cur.토요일),
      saturdayMs: parseFloat(cur.토요일) / 3.6,
      sundayKmh: parseFloat(cur.일요일),
      sundayMs: parseFloat(cur.일요일) / 3.6,
    });
    return acc;
  }, [] as ISpeed[]);

  await prisma.link_speed_average.createMany({
    data: speeds.map((v) => ({
      ...v,
      weekdayMs: Number(v.weekdayMs.toFixed(2)),
      saturdayMs: Number(v.saturdayMs.toFixed(2)),
      sundayMs: Number(v.sundayMs.toFixed(2)),
    })),
    skipDuplicates: true,
  });
}

export async function makeStandardLinkData() {
  await getAverageSpeed();

  const averageSpeed = await prisma.link_speed_average.findMany();
  const speeds = averageSpeed.filter((s) =>
    ["서울", "경기"].includes(s.location),
  );

  const seoulSpeed = speeds.find((s) => s.location === "서울")!;
  const ggSpeed = speeds.find((s) => s.location === "경기")!;

  const ggRouteIds = await prisma.gyeonggiLink.groupBy({
    by: ["route_id"],
  });

  const seoulRouteIds = await prisma.seoulLink.groupBy({
    by: ["route_id"],
  });

  for (let i = 0; i < ggRouteIds.length; i++) {
    console.log(`경기 ${i + 1}번째`);
    const routeId = ggRouteIds[i].route_id;
    const ggLink = await prisma.gyeonggiLink.findMany({
      where: {
        route_id: routeId,
      },
    });
    const dataset = makeGGDataset(ggLink, ggSpeed);
    await prisma.link.createMany({
      data: dataset.map((v) => ({
        routeId: Number(v.routeId),
        startStationId: Number(v.startStationId),
        endStationId: Number(v.endStationId),
        tripTime: v.tripTime,
        tripDistance: v.tripDistance,
        stationOrder: v.stationOrder,
      })),
      skipDuplicates: true,
    });
  }

  for (let i = 0; i < seoulRouteIds.length; i++) {
    console.log(`서울 ${i + 1}번째`);
    const routeId = seoulRouteIds[i].route_id;
    const seoulLink = await prisma.seoulLink.findMany({
      where: {
        route_id: routeId,
      },
    });
    const dataset = makeSeoulDataset(seoulLink, seoulSpeed);
    await prisma.link.createMany({
      data: dataset.map((v) => ({
        routeId: Number(v.routeId),
        startStationId: Number(v.startStationId),
        endStationId: Number(v.endStationId),
        tripTime: v.tripTime,
        tripDistance: v.tripDistance,
        stationOrder: v.stationOrder,
      })),
      skipDuplicates: true,
    });
  }
}
