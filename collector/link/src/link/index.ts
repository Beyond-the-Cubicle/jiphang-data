import * as fs from "fs";
import { makeGGDataset, makeSeoulDataset } from "./utils";
import { format, writeToPath } from "fast-csv";
import dotenv from "dotenv";
import { PrismaClient } from "@prisma/client";
import { IBasicLink, ISpeed } from "./interface";
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

  const generateSeoulBasicLink = async () => {
    const seoulSpeed = speeds.find((s) => s.location === "서울")!;
    const seoulLink = await prisma.seoulLink.findMany({
      orderBy: [
        {
          route_id: "asc",
        },
        {
          sttn_ordr: "asc",
        },
      ],
    });
    const dataset = makeSeoulDataset(seoulLink, seoulSpeed);

    const data = dataset.map((v) => ({
      routeId: v.routeId,
      startStationId: v.startStationId,
      endStationId: v.endStationId,
      tripTime: v.tripTime,
      tripDistance: v.tripDistance,
      stationOrder: v.stationOrder,
    }));

    return data;
  };

  const generateGGBasicLinkData = async () => {
    const ggSpeed = speeds.find((s) => s.location === "경기")!;
    const ggLink = await prisma.gyeonggiLink.findMany();
    const dataset = makeGGDataset(ggLink, ggSpeed);

    const data = dataset.map((v) => ({
      routeId: v.routeId,
      startStationId: v.startStationId,
      endStationId: v.endStationId,
      tripTime: v.tripTime,
      tripDistance: v.tripDistance,
      stationOrder: v.stationOrder,
    }));

    return data;
  };

  const seoulData = await generateSeoulBasicLink();
  const ggData = await generateGGBasicLinkData();

  const dataset = seoulData.concat(ggData);
  console.log("중복삭제 전:", dataset.length);
  // 중복 삭제
  const map = new Map<string, IBasicLink>();
  dataset.forEach((v) => {
    const key = `${v.routeId}-${v.startStationId}-${v.endStationId}`;
    map.set(key, v);
  });
  const result = Array.from(map.values());

  console.log("중복삭제 후:", result.length);
  console.log("중복 갯수: ", dataset.length - result.length);

  // TODO 메모리 터지는거 고쳐야함 ...
  writeToPath("./total-link.csv", result, {
    quote: true,
    headers: [
      "routeId",
      "startStationId",
      "endStationId",
      "tripTime",
      "tripDistance",
      "stationOrder",
    ],
  });
}
