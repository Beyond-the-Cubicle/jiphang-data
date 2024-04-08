import {PrismaClient} from "@prisma/client";
import * as fs from "fs";
import {makeGGDataset, makeSeoulDataset} from "./utils";
import {format} from "fast-csv";

const prisma = new PrismaClient();

const averageSpeed = await prisma.link_speed_average.findMany();
const speeds = averageSpeed.filter(s => ['서울', '경기'].includes(s.location))

const generateSeoulBasicLink = async () => {
  const seoulSpeed = speeds.find(s => s.location === '서울')!;
  const seoulLink = await prisma.seoulLink.findMany({
    orderBy: [{
      route_id: 'asc',
    }, {
      sttn_ordr: 'asc',
    }]
  });
  const dataset = makeSeoulDataset(seoulLink, seoulSpeed);
  console.log('서울:', dataset.length);

  const data = dataset.map(v => ({
    routeId: BigInt(v.routeId),
    startStationId: BigInt(v.startStationId),
    endStationId: BigInt(v.endStationId),
    tripTime: v.tripTime,
    tripDistance: v.tripDistance,
    stationOrder: v.stationOrder,
  }));

  return data;
}

const generateGGBasicLinkData = async () => {
  const ggSpeed = speeds.find(s => s.location === '경기')!;
  const ggLink = await prisma.gyeonggiLink.findMany();
  const dataset = makeGGDataset(ggLink, ggSpeed);
  console.log('경기:', dataset.length);

  const data = dataset.map(v => ({
    routeId: BigInt(v.routeId),
    startStationId: BigInt(v.startStationId),
    endStationId: BigInt(v.endStationId),
    tripTime: v.tripTime,
    tripDistance: v.tripDistance,
    stationOrder: v.stationOrder,
  }));

  return data;
}

const seoulData = await generateSeoulBasicLink();
console.log('서울:', seoulData.length);
const ggData = await generateGGBasicLinkData();
console.log('경기:', ggData.length);

const dataset = seoulData.concat(ggData);
console.log('총:', dataset.length);
// 중복 삭제
const set = new Set<string>();
const result = dataset.filter(v => {
  const key = `${v.routeId}-${v.startStationId}-${v.endStationId}`;
  if (set.has(key)) {
    return false;
  }

  set.add(key);
  return true;
});

console.log('중복삭제 후:', result.length)

const stream = format({
  quote: true,
  headers: ['routeId', 'startStationId', 'endStationId', 'tripTime', 'tripDistance', 'stationOrder'],
})

stream.pipe(fs.createWriteStream('./total-link.csv')).on('end', () => {
  console.log('done');
});

for (let i = 0; i < result.length; i++) {
  stream.write(result[i]);
}

